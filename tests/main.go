package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/el-tumero/banana-vrf-client/user"
	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
)

var u *user.User
var ctx context.Context

func deployReqHandler(w http.ResponseWriter, r *http.Request) {
	_, _, err := DeployContract(ctx, u)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("Deployed!"))
}

func nextBlockHandler(w http.ResponseWriter, r *http.Request) {
	if err := NextBlock(ctx); err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("Next block!"))
}

func addStake0(w http.ResponseWriter, r *http.Request) {
	if err := u.AddStake(ctx, uint256.NewInt(200)); err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("Stake added!"))
}

func addStake1(w http.ResponseWriter, r *http.Request) {
	u1, err := GetTestUser(ctx, 1)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	u1.AddContract(common.HexToAddress(user.TEST_CONTRACT_ADDR))

	if err := u1.AddStake(ctx, uint256.NewInt(200)); err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("Stake added!"))
}

func checkStakeHandler(w http.ResponseWriter, r *http.Request) {
	stake, err := u.CheckStake(u.GetAddress2())
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	// fmt.Println("Jadziem")
	fmt.Println(u.GetAddress2().Hex())
	fmt.Println(stake.String())
	w.Write([]byte(stake.String()))
}

func checkRoundData(w http.ResponseWriter, r *http.Request) {
	round, err := u.GetRoundData(1)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Println(round.RandomNumber.String())
	w.Write([]byte("OK!"))
}

func main() {
	ctx = context.Background()
	u0, err := GetTestUser(ctx, 0)
	if err != nil {
		log.Fatal(err)
	}
	u0.AddContract(common.HexToAddress(user.TEST_CONTRACT_ADDR))
	u = u0

	http.HandleFunc("/deploy", deployReqHandler)
	http.HandleFunc("/next", nextBlockHandler)
	http.HandleFunc("/add0", addStake0)
	http.HandleFunc("/add1", addStake1)
	http.HandleFunc("/stake", checkStakeHandler)
	http.HandleFunc("/round", checkRoundData)

	http.ListenAndServe(":3327", nil)

}
