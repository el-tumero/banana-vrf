package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/el-tumero/banana-vrf-server/conns"
	"github.com/gorilla/websocket"
)

const MAX_MESSAGE_LEN = 128

var broadcast = make(chan []byte, 10)

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
			log.Println("Can't read message! ", err)
			break
		}

		if err != nil {
			log.Println("read:", err)
			break
		}
		broadcast <- message
	}
	conns.RemoveConn(id)
}

func handleBroadcast(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Broadcast closed")
			return
		case message := <-broadcast:
			for _, c := range conns.GetConns() {
				if c != nil {
					err := c.WriteMessage(1, message)
					if err != nil {
						fmt.Println(err)
					}
				}
			}
		}
	}

}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	mux := http.NewServeMux()
	mux.HandleFunc("/ws", handler)
	server := &http.Server{Addr: ":3333", Handler: mux}
	go handleBroadcast(ctx)

	go func() {
		fmt.Println("Server is running on :3333")
		err := server.ListenAndServe()
		if err != nil {
			fmt.Println(err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig)

	<-sig
	cancel()

	serverCtx, cancelCtx := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelCtx()
	if err := server.Shutdown(serverCtx); err != nil {
		fmt.Println(err)
	}

}
