package main

import (
	"log"
	"net/http"
	"webrtc-app/internal/application"
	"webrtc-app/internal/infra"
	"webrtc-app/internal/transport/handler"
)

func main() {
	roomRepo := infra.NewInMemoryRoomRepository()
	clientRepo := infra.NewInMemoryClientRepository()

	roomService := application.NewRoomService(roomRepo)
	clientService := application.NewClientService(clientRepo)

	clientHandler := handler.NewClientHandler(clientService)
	roomHandler := handler.NewRoomHandler(roomService)

	http.HandleFunc("/create", roomHandler.CreateRoomRequestHandler)
	http.HandleFunc("POST /join/{id}", roomHandler.JoinRoomRequestHandler)

	http.HandleFunc("/client/create", clientHandler.CreateClientHandler)

	log.Println("starting server on port 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
