package coordinator

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/el-tumero/banana-vrf-client/proposals"
	"github.com/el-tumero/banana-vrf-client/user"
	"github.com/gorilla/websocket"
)

var thrBlockHeight *big.Int
var decBlockHeight *big.Int
var proBlockHeight *big.Int

const THRESHOLD int64 = 5
const DECISION int64 = 2
const PROPOSAL int64 = 1

func UpdateHeights(roundId uint32, u *user.User) error {
	round, err := u.GetRoundData(roundId)
	if err != nil {
		return err
	}

	thr := big.NewInt(THRESHOLD)
	dec := big.NewInt(DECISION)
	pro := big.NewInt(PROPOSAL)

	thrBlockHeight = big.NewInt(0).Add(round.BlockHeight, thr)
	decBlockHeight = big.NewInt(0).Add(round.BlockHeight, dec)
	proBlockHeight = big.NewInt(0).Add(round.BlockHeight, pro)
	return nil
}

func Init(u *user.User, conn *websocket.Conn) error {
	isActive := u.IsOperatorActive(u.GetAddress2())
	if !isActive {
		return fmt.Errorf("your operator is not active yet - you need to wait")
	}

	roundId, err := u.GetCurrRoundId()
	if err != nil {
		return err
	}

	err = UpdateHeights(roundId, u)
	if err != nil {
		return err
	}

	prevRound, err := u.GetRoundData(roundId - 1)
	if err != nil {
		return err
	}

	vrf, err := u.GenerateVrf(prevRound.RandomNumber)
	if err != nil {
		return err
	}
	if err := proposals.Propose(conn, roundId, vrf); err != nil {
		return err
	}

	return nil
}

func DecisionProc(ctx context.Context, u *user.User, conn *websocket.Conn) {
	blockSub, headers, err := u.CreateNewBlockSub(ctx)
	if err != nil {
		fmt.Println("decision:", err)
		return
	}

	evtSub, info, err := u.CreateEventSub()
	if err != nil {
		fmt.Println("decision:", err)
		return
	}

	waitForNextRound := false

	for {
		select {
		case err := <-blockSub.Err():
			fmt.Println("decision:", err)
			return
		case err := <-evtSub.Err():
			fmt.Println("decision:", err)
			return
		case <-ctx.Done():
			fmt.Println("DecProc closed")
			return
		case evt := <-info:
			roundId := evt.Topics[1].Big()
			fmt.Println("Next round! ", roundId.String())
			proposals.FlushStorage()
			roundId32 := uint32(roundId.Uint64())

			err := UpdateHeights(roundId32, u)
			if err != nil {
				fmt.Println("Can't update heights!")
				break
			}

			prevRound, err := u.GetRoundData(roundId32 - 1)
			if err != nil {
				fmt.Println(err)
				break
			}
			vrf, err := u.GenerateVrf(prevRound.RandomNumber)
			if err != nil {
				fmt.Println("Can't generate VRF")
				break
			}
			err = proposals.Propose(conn, roundId32, vrf)
			if err != nil {
				fmt.Println("Can't propose", err)
			}

		case header := <-headers:
			block, err := u.GetBlockchainClient().BlockByHash(ctx, header.Hash())
			if err != nil {
				fmt.Println("block by hash err ", err)
			}
			fmt.Println(block.Number().String())

			if block.Number().Cmp(proBlockHeight) == 0 {
				roundId, err := u.GetCurrRoundId()
				if err != nil {
					fmt.Println(err)
					break
				}
				round, err := u.GetRoundData(roundId - 1)
				if err != nil {
					fmt.Println(err)
					break
				}
				vrf, err := u.GenerateVrf(round.RandomNumber)
				if err != nil {
					fmt.Println(err)
					break
				}
				err = proposals.Propose(conn, roundId, vrf)
				if err != nil {
					fmt.Println(err)
					break
				}
				waitForNextRound = false
			}

			if block.Number().Cmp(thrBlockHeight) == 1 && proposals.Candidate && !waitForNextRound {
				// finalize round
				if err := u.FinalizeRound(ctx); err != nil {
					fmt.Println("can't finalize round", err)
					break
				}
				fmt.Println("Round closed!")
				proposals.Candidate = false
				waitForNextRound = true
				break
			}

			// comparing block number
			if block.Number().Cmp(decBlockHeight) == 1 && !proposals.Candidate && !waitForNextRound {
				if bytes.Equal(proposals.DiscoverSmallest().Vrf, proposals.LastLocalProposal.Vrf) {
					// add proposal
					if err := u.SetRandomNumber(ctx, proposals.LastLocalProposal.Vrf); err != nil {
						fmt.Println("can't set random number", err)
						break
					}
					fmt.Println("Random number set!")
					proposals.Candidate = true
				}
			}

		}

	}

}

func Read(ctx context.Context, conn *websocket.Conn, u *user.User) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Read closed")
			return
		default:
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("read error ", err)
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

// func NextRoundActions(ctx context.Context, u *user.User) {
// 	sub, info, err := u.CreateEventSub()
// 	if err != nil {
// 		fmt.Println("decision:", err)
// 		return
// 	}

// 	for {
// 		select {
// 		case err := <-sub.Err():
// 			fmt.Println("nextRoundActions:", err)
// 			return
// 		case <-ctx.Done():
// 			fmt.Println("NextRoundActions closed")
// 			return
// 		case evt := <-info:
// 			data, err := user.LocalAbi.Unpack("NewRound", evt.Data)
// 			if err != nil {
// 				fmt.Println("Can't read event!")
// 			}
// 			fmt.Println(data)
// 			// fmt.Println(evt.)
// 		}

// 	}
// }

// case header := <-headers:
