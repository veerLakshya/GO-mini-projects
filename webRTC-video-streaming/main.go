package main

import (
	"log"
	"net/http"
	"webrtc-app/server"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	server.AllRooms.Init()

	mux := http.NewServeMux()

	log.Println("all rooms", server.AllRooms.Map)

	mux.HandleFunc("/create", server.CreateRoomRequestHandler)
	mux.HandleFunc("/join", server.JoinRoomRequestHandler)

	log.Println("starting server on port 8000")
	err := http.ListenAndServe(":8000", corsMiddleware(mux))
	if err != nil {
		panic(err)
	}
}
