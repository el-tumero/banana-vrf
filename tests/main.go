package main

import (
	"context"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/el-tumero/banana-vrf-client/user"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/holiman/uint256"
)

var u0 *user.User
var u1 *user.User

var ctx context.Context

func deployReqHandler(w http.ResponseWriter, r *http.Request) {
	_, _, err := DeployContract(ctx, u0)
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
	if err := u0.AddStake(ctx, uint256.NewInt(200)); err != nil {
		fmt.Println("Can't add stake :(")
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("Stake added!"))
}

func addStake1(w http.ResponseWriter, r *http.Request) {
	if err := u1.AddStake(ctx, uint256.NewInt(200)); err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("Stake added!"))
}

func checkStakeHandler(w http.ResponseWriter, r *http.Request) {
	stake, err := u0.CheckStake(u0.GetAddress2())
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Println(u0.GetAddress2().Hex())
	fmt.Println(stake.String())
	w.Write([]byte(stake.String()))
}

func roundNext(w http.ResponseWriter, r *http.Request) {
	if err := u0.FinalizeRound(ctx); err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("OK!"))
}

func checkRoundData(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	idint, err := strconv.Atoi(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	round, err := u0.GetRoundData(uint32(idint))
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	// fmt.Println(round.RandomNumber.String())
	res := fmt.Sprintf("[op: %s] [num: %s] [hash: %s] [state: %d]\n", round.Proposer.String(), round.RandomNumber, hex.EncodeToString(round.RandomNumberHash[:]), round.State)
	w.Write([]byte(res))
}

func main() {

	rpcAddr := flag.String("rpc", user.TEST_RPC, "rpc address")
	chainId := flag.Int64("chain_id", 1337, "chain id")
	privateKey0 := flag.String("priv0", PRIVATE_KEYS[0], "#0 private key address")
	privateKey1 := flag.String("priv1", PRIVATE_KEYS[1], "#1 private key address")
	contractAddr := flag.String("contract", user.TEST_CONTRACT_ADDR, "contract address")

	flag.Parse()

	ctx = context.Background()

	user.CHAIN_ID = *chainId

	priv0, err := crypto.HexToECDSA(*privateKey0)
	if err != nil {
		log.Fatal(err)
	}
	u00, err := user.NewFromPrivateKey(priv0)
	if err != nil {
		log.Fatal(err)
	}
	u0 = u00
	err = u0.ConnectToBlockchain(ctx, *rpcAddr)
	if err != nil {
		log.Fatal(err)
	}
	err = u0.AddContract(common.HexToAddress(*contractAddr))
	if err != nil {
		log.Fatal(err)
	}

	priv1, err := crypto.HexToECDSA(*privateKey1)
	if err != nil {
		log.Fatal(err)
	}
	u11, err := user.NewFromPrivateKey(priv1)
	if err != nil {
		log.Fatal(err)
	}
	u1 = u11
	err = u1.ConnectToBlockchain(ctx, *rpcAddr)
	if err != nil {
		log.Fatal(err)
	}
	err = u1.AddContract(common.HexToAddress(*contractAddr))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected!")
	fmt.Println(u0.GetAddress(), u1.GetAddress())

	http.HandleFunc("/deploy", deployReqHandler)
	http.HandleFunc("/next", nextBlockHandler)
	http.HandleFunc("/add0", addStake0)
	http.HandleFunc("/add1", addStake1)
	http.HandleFunc("/stake", checkStakeHandler)
	http.HandleFunc("/round", checkRoundData)
	http.HandleFunc("/roundnext", roundNext)

	err = http.ListenAndServe(":3327", nil)
	if err != nil {
		log.Fatal(err)
	}

}
