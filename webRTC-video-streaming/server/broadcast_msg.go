package server

import (
	"log"

	"github.com/gorilla/websocket"
)

type BroadcastMsg struct {
	Message map[string]interface{}
	RoomId  string
	Conn    *websocket.Conn
}

var Broadcast = make(chan BroadcastMsg)

func Broadcaster() {
	for {
		msg := <-Broadcast

		room, ok := AllRooms.Get(msg.RoomId)
		if !ok {
			continue
		}

		for _, client := range room.GetClients() {
			if client.Conn == msg.Conn {
				continue
			}
			err := client.Conn.WriteJSON(msg.Message)
			if err != nil {
				log.Println("broadcast write error:", err)
				client.Conn.Close()
			}
		}
	}
}
