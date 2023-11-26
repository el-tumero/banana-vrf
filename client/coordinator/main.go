package coordinator

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/el-tumero/banana-vrf-client/proposals"
	"github.com/el-tumero/banana-vrf-client/user"
	"github.com/gorilla/websocket"
)

// var roundBlockHeight *big.Int
var decBlockHeight *big.Int

const THRESHOLD int64 = 5
const DEC int64 = 2

func Init(u *user.User, conn *websocket.Conn) error {
	isActive := u.IsOperatorActive(u.GetAddress2())
	if !isActive {
		return fmt.Errorf("your operator is not active yet - you need to wait")
	}

	roundId, err := u.GetCurrRoundId()
	if err != nil {
		return err
	}
	round, err := u.GetRoundData(roundId - 1)
	if err != nil {
		return err
	}

	dec := big.NewInt(DEC)
	decBlockHeight = big.NewInt(0).Add(round.BlockHeight, dec)

	vrf, err := u.GenerateVrf(round.RandomNumber)
	if err != nil {
		return err
	}
	if err := proposals.Propose(conn, vrf); err != nil {
		return err
	}

	return nil
}

func DecisionProc(ctx context.Context, u *user.User) {
	sub, headers, err := u.CreateNewBlockSub(ctx)
	if err != nil {
		fmt.Println("decision:", err)
		return
	}

	for {
		select {
		case err := <-sub.Err():
			fmt.Println("decision:", err)
			return
		case <-ctx.Done():
			fmt.Println("DecProc closed")
			return
		case header := <-headers:
			block, err := u.GetBlockchainClient().BlockByHash(ctx, header.Hash())
			if err != nil {
				fmt.Println("block by hash err ", err)
			}

			// comparing block number
			fmt.Println(block.Number().String())
			if block.Number().Cmp(decBlockHeight) == 1 {
				proposals.PrintStorage()
				// fmt.Println(proposals.GetStorage())
				// fmt.Println("TIME FOR DEC")

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

// wait for bnt/2
