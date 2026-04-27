package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type resp struct {
	RoomId string `json:"room_id"`
	Error  bool
}

var AllRooms RoomsMap

// create a room and return a roomId
func CreateRoomRequestHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("CreateRoomRequestHandler: ")

	room := NewRoom()
	AllRooms.AddRoom(room)

	log.Println("All Rooms: ", AllRooms.Map)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(resp{
		RoomId: room.GetID().String(),
		Error:  false,
	})
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func JoinRoomRequestHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("JoinRoomRequestHandler: ")

	roomId := r.PathValue("id")
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "web socket upgrade error: ", http.StatusInternalServerError)
		return
	}

	AllRooms.InsertIntoRoom(roomId, false, ws)

	go Broadcaster()

	for {
		var msg BroadcastMsg

		err := ws.ReadJSON(&msg.Message)
		if err != nil {
			log.Fatal("Read Error: ", err)
		}
		msg.Conn = ws
		msg.RoomId = roomId

		Broadcast <- msg
	}

}
