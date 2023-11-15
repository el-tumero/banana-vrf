package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/el-tumero/banana-vrf-server/conns"
	"github.com/gorilla/websocket"
)

const MAX_MESSAGE_LEN = 1024

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

func handler(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer c.Close()

	c.SetReadLimit(MAX_MESSAGE_LEN)

	id := conns.AddConn(c)

	for {
		mt, message, err := c.ReadMessage()
		if mt == -1 {
			log.Println("Connection closed!")
			break
		}

		if err != nil {
			log.Println("read:", err)
			break
		}
		broadcast(message)
	}

	conns.RemoveConn(id)
}

func broadcast(msg []byte) {
	for _, c := range conns.GetConns() {
		if c != nil {
			err := c.WriteMessage(1, msg)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func main() {
	http.HandleFunc("/ws", handler)
	err := http.ListenAndServe("localhost:3333", nil)
	log.Fatal(err)
}
