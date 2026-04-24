package main

import (
	"log"
	"net/http"
	"webrtc-app/server"
)

func main() {
	server.AllRooms.Init()

	log.Println("all rooms", server.AllRooms.Map)

	http.HandleFunc("/create", server.CreateRoomRequestHandler)
	http.HandleFunc("/join", server.JoinRoomRequestHandler)

	log.Println("starting server on port 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
