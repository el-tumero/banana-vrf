package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/el-tumero/banana-vrf-client/proposals"
	"github.com/el-tumero/banana-vrf-client/user"
	"github.com/ethereum/go-ethereum/common"

	"github.com/gorilla/websocket"
)

var u *user.User
var conn *websocket.Conn

func main() {
	userApiPort := flag.Int("port", 3030, "port for user api")
	rpcAddr := flag.String("rcp", user.TEST_RPC, "rpc address")
	chainId := flag.Int64("chain_id", 1337, "chain id")
	contractAddr := flag.String("contract", user.TEST_CONTRACT_ADDR, "contract address")

	flag.Parse()

	ctx := context.Background()

	addr := url.URL{Scheme: "ws", Host: "127.0.0.1:3333", Path: "/ws"}

	c, _, err := websocket.DefaultDialer.Dial(addr.String(), nil)
	if err != nil {
		log.Fatal("dial error", err)
	}
	conn = c

	// user setup
	u, err = user.New()
	if err != nil {
		log.Fatal("user creation error", err)
	}
	if err := u.ConnectToBlockchain(ctx, *rpcAddr); err != nil {
		log.Fatal("can't connect to blockchain", err)
	}
	user.CHAIN_ID = *chainId

	if err := u.AddContract(common.HexToAddress(*contractAddr)); err != nil {
		log.Fatal("can't add contract", err)
	}

	// debug api
	http.HandleFunc("/read", readReqHandler)

	go Read(c)

	err = http.ListenAndServe(fmt.Sprintf(":%d", *userApiPort), nil)
	if err != nil {
		log.Fatal("can't start user api")
	}

	defer c.Close()
}

// func sendReqHandler(w http.ResponseWriter, r *http.Request) {
// 	Send(conn)
// 	w.Write([]byte("OK!"))
// }

func readReqHandler(w http.ResponseWriter, r *http.Request) {
	var out []byte = []byte{}
	for _, p := range proposals.GetStorage() {
		out = append(out, []byte(p.Num.String())...)
		out = append(out, 10) // \n
	}
	w.Write(out)
}

// func Send(conn *websocket.Conn) {
// 	vrf, err := u.GenerateVrf(mock.GetRandomNumber())
// 	if err != nil {
// 		log.Println("send error", err)
// 		return
// 	}
// 	err = u.Propose(conn, vrf)
// 	if err != nil {
// 		log.Println("propose error", err)
// 	}

// 	log.Println("sent: ", new(uint256.Int).SetBytes(vrf[16:48]).String())
// }

func Read(conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read error", err)
			return
		}
		p, err := proposals.CastBytes(message)
		if err != nil {
			log.Println("cast error", err)
		}
		log.Println("Received propsal!")
		proposals.AddProposalToStorage(p)
	}
}
