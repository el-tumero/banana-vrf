package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/el-tumero/banana-vrf-client/api"
	"github.com/el-tumero/banana-vrf-client/coordinator"
	"github.com/el-tumero/banana-vrf-client/user"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gorilla/websocket"
)

var u *user.User
var conn *websocket.Conn

func main() {
	userApiPort := flag.Int("port", 3030, "port for user api")
	rpcAddr := flag.String("rpc", user.TEST_RPC, "rpc address")
	chainId := flag.Int64("chain_id", 1337, "chain id")
	contractAddr := flag.String("contract", user.TEST_CONTRACT_ADDR, "contract address")
	privateKey := flag.String("priv", "567eade5964411e5c837c03de980e0e006cfab066f1faffee2b82dea5969a942", "private key address")

	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())

	addr := url.URL{Scheme: "ws", Host: "127.0.0.1:3333", Path: "/ws"}

	c, _, err := websocket.DefaultDialer.Dial(addr.String(), nil)
	if err != nil {
		log.Fatal("dial error", err)
	}
	conn = c

	// user setup
	// u, err = user.New()
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

	if err := coordinator.Init(u, conn); err != nil {
		log.Fatal("coordinator init err ", err)
	}

	go coordinator.DecisionProc(ctx, u, c)
	go coordinator.Read(ctx, c, u)
	// go coordinator.NextRoundActions(ctx, u)

	// create http server
	server := api.CreateHttpServer(*userApiPort)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	fmt.Println("Stop!")
	cancel()
	c.Close()

	ctxServer, cancelServer := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelServer()
	if err := server.Shutdown(ctxServer); err != nil {
		log.Println("Server shutdown error ", err)
		return
	}
}
