package coordinator

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"math/big"
	"net/url"
	"time"

	"github.com/el-tumero/banana-vrf-client/proposals"
	"github.com/el-tumero/banana-vrf-client/user"
	"github.com/gorilla/websocket"
)

const THRESHOLD int64 = 5
const DECISION int64 = 2
const PROPOSAL int64 = 1

const MAX_RECONN = 10

const ROUND_EMPTY = 0
const ROUND_PROPOSAL = 1
const ROUND_FINAL = 2

var currentRoundId uint32 = 0

func GetDecisionHeight(roundHeight *big.Int) *big.Int {
	n := big.NewInt(0)
	return n.Add(roundHeight, big.NewInt(2))
}

func GetCloseHeight(roundHeight *big.Int) *big.Int {
	n := big.NewInt(0)
	return n.Add(roundHeight, big.NewInt(5))
}

func GetLateHeight(roundHeight *big.Int) *big.Int {
	n := big.NewInt(0)
	return n.Add(roundHeight, big.NewInt(15))
}

func Start(ctx context.Context, u *user.User, relay *string) {
	disCh := make(chan struct{})
	clsCh := make(chan struct{})

	isActive := u.IsOperatorActive(u.GetAddress2())
	if !isActive {
		fmt.Println("your operator is not active yet - you need to wait")
		return
	}

	addr := url.URL{Scheme: "ws", Host: *relay, Path: "/ws"}
	var conn *websocket.Conn

	c, _, err := websocket.DefaultDialer.Dial(addr.String(), nil)
	if err != nil {
		log.Fatal("dial error", err)
	}
	conn = c

	go DecisionProc(ctx, u, conn, clsCh)
	go Read(ctx, conn, u, disCh)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Coordinator closed!")
			conn.Close()
			return
		case <-disCh:
			fmt.Println("Disconnected from socket!")
			close(clsCh)
			conn.Close()
			conn = Reconnect(addr.String())
			if conn == nil {
				fmt.Println("Can't reconnect!")
				return
			}
			fmt.Println("Reconnected!")
			disCh = make(chan struct{})
			clsCh = make(chan struct{})
			go DecisionProc(ctx, u, conn, clsCh)
			go Read(ctx, conn, u, disCh)
		}
	}

}

func DecisionProc(ctx context.Context, u *user.User, conn *websocket.Conn, clsChan chan struct{}) {
	blockSub, headers, err := u.CreateNewBlockSub(ctx)
	if err != nil {
		fmt.Println("[DecisionProc] error:", err)
		return
	}

	for {
		select {
		case err := <-blockSub.Err():
			fmt.Println("[DecisionProc] error:", err)
			return
		case <-ctx.Done():
			fmt.Println("[DecisionProc] closed")
			return
		case <-clsChan:
			fmt.Println("[DecisionProc closed] - client disconnected")
			blockSub.Unsubscribe()
			currentRoundId = 0
			proposals.FlushStorage()
			return
		case header := <-headers:
			block, err := u.GetBlockchainClient().BlockByHash(ctx, header.Hash())
			if err != nil {
				fmt.Println("[DecisionProc] block by hash err ", err)
				break
			}
			fmt.Println(block.Number().String())

			// Fetching id of the current round when initializing the client or after the round is finalized
			if currentRoundId == 0 {
				currentRoundId, err = u.GetCurrRoundId()
				if err != nil {
					fmt.Println("[DecisionProc] can't get current round id ", err)
					break
				}
				fmt.Println("Round:", currentRoundId)
			}

			// Fetching data of the current round
			r, err := u.GetRoundData(currentRoundId)
			if err != nil {
				fmt.Println("[DecisionProc] can't get round data ", err)
				break
			}

			// Restarting mechanism when the round is finalized
			if r.State == ROUND_FINAL {
				fmt.Println("Round closed!")
				currentRoundId = 0
				proposals.FlushStorage()
				break
			}

			// Proposal stage
			if r.State == ROUND_EMPTY {
				pr, err := u.GetRoundData(currentRoundId - 1)
				if err != nil {
					fmt.Println("[DecisionProc] can't get round data ", err)
					break
				}
				vrf, err := u.GenerateVrf(pr.RandomNumber)
				if err != nil {
					fmt.Println(err)
					break
				}
				proposals.Propose(conn, currentRoundId, vrf)
			}

			// Contract write stage
			if r.State == ROUND_EMPTY && block.Number().Cmp(GetDecisionHeight(r.BlockHeight)) == 1 {
				smallest := proposals.DiscoverSmallest()
				if smallest == nil {
					fmt.Println("[DecisionProc] no values in storage...")
					break
				}
				if bytes.Equal(smallest.Vrf, proposals.LastLocalProposal.Vrf) {
					if err := u.SetRandomNumber(ctx, proposals.LastLocalProposal.Vrf); err != nil {
						fmt.Println("[DecisionProc] can't set random number", err)
						break
					}
					fmt.Println("Random number set!")
				}
				break
			}

			// Close stage
			if r.State == ROUND_PROPOSAL &&
				block.Number().Cmp(GetCloseHeight(r.BlockHeight)) == 1 &&
				u.GetAddress2().Cmp(r.Proposer) == 0 {
				if err := u.FinalizeRound(ctx); err != nil {
					fmt.Println("can't finalize round ", err)
					break
				}
				fmt.Println("Closing round...")
				break
			}

			// Close late stage
			if r.State == ROUND_PROPOSAL && block.Number().Cmp(GetLateHeight(r.BlockHeight)) == 1 {
				if err := u.FinalizeRoundLate(ctx); err != nil {
					fmt.Println("can't finalize round late ", err)
					break
				}
				fmt.Println("Closing round (late)!")
			}

		}
	}

}

func Read(ctx context.Context, conn *websocket.Conn, u *user.User, disCh chan struct{}) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Read closed")
			return
		default:
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("read error ", err)
				fmt.Println("Read closed")
				close(disCh)
				return
			}
			p, err := proposals.CastBytes(message)
			if err != nil {
				log.Println("cast error ", err)
				break
			}
			log.Println("Received propsal!")
			proposals.AddProposalToStorage(p, u)
		}
	}
}

func Reconnect(addr string) *websocket.Conn {
	for i := 0; i < MAX_RECONN; i++ {
		time.Sleep(2 * time.Second)
		c, _, err := websocket.DefaultDialer.Dial(addr, nil)
		fmt.Println("Trying to reconnect!")
		if err != nil {
			continue
		}
		return c
	}
	return nil
}
