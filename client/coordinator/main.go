package coordinator

import (
	"context"
	"fmt"
	"math/big"

	"github.com/el-tumero/banana-vrf-client/user"
	"github.com/gorilla/websocket"
	"github.com/holiman/uint256"
)

var roundBlockHeight *big.Int

func Init(u *user.User, conn *websocket.Conn) error {
	isActive := u.IsOperatorActive(u.GetAddress2())
	if !isActive {
		return fmt.Errorf("your operator is not active yet - you need to wait")
	}

	roundId, err := u.GetCurrRoundId()
	if err != nil {
		return err
	}
	round, err := u.GetRoundData(roundId)
	if err != nil {
		return err
	}
	roundBlockHeight = round.BlockHeight

	vrf, err := u.GenerateVrf(uint256.MustFromBig(round.RandomNumber))
	if err := u.Propose(conn, vrf); err != nil {
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
			fmt.Println(headers, header.Hash().Hex())
		}

	}

}

// wait for bnt/2
