package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/el-tumero/banana-vrf-client/coordinator"
	"github.com/el-tumero/banana-vrf-client/user"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var u *user.User

func main() {
	rpcAddr := flag.String("rpc", user.TEST_RPC, "rpc address")
	relay := flag.String("relay", "127.0.0.1:3333", "relay address")
	chainId := flag.Int64("chain_id", 1337, "chain id")
	contractAddr := flag.String("contract", user.TEST_CONTRACT_ADDR, "contract address")
	privateKey := flag.String("priv", "567eade5964411e5c837c03de980e0e006cfab066f1faffee2b82dea5969a942", "private key address")

	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())

	priv, err := crypto.HexToECDSA(*privateKey)
	if err != nil {
		log.Fatal("private key conversion err", err)
	}
	u, err = user.NewFromPrivateKey(priv)
	if err != nil {
		log.Fatal("user creation error ", err)
	}
	if err := u.ConnectToBlockchain(ctx, *rpcAddr); err != nil {
		log.Fatal("can't connect to blockchain ", err)
	}
	user.CHAIN_ID = *chainId

	if err := u.AddContract(common.HexToAddress(*contractAddr)); err != nil {
		log.Fatal("can't add contract ", err)
	}
	fmt.Println("Connected as:", u.GetAddress())

	go coordinator.Start(ctx, u, relay)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	fmt.Println("Stop!")
	cancel()

}
