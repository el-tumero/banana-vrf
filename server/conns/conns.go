package conns

import "github.com/gorilla/websocket"

var conns []*websocket.Conn

func AddConn(conn *websocket.Conn) int {
	for i, v := range conns {
		if v == nil {
			conns[i] = conn
			return i
		}
	}

	conns = append(conns, conn)
	return len(conns) - 1
}

func RemoveConn(id int) {
	conns[id] = nil
}

func GetConns() []*websocket.Conn {
	return conns
}
