package server

import (
	"encoding/json"
	"log"
	"net/http"
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

	log.Panicln("All Rooms: ", AllRooms.Map)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(resp{
		RoomId: room.GetID().String(),
		Error:  false,
	})
}

func JoinRoomRequestHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("JoinRoomRequestHandler: ")
}
