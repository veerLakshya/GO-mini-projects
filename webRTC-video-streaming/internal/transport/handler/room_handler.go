package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"webrtc-app/internal/application"
	ws "webrtc-app/internal/transport/ws"
)

type roomResp struct {
	RoomId string `json:"room_id"`
	Error  bool   `json:"error"`
}

type RoomHandler struct {
	roomService *application.RoomService
}

func NewRoomHandler(roomService *application.RoomService) *RoomHandler {
	return &RoomHandler{roomService: roomService}
}

// create a room and return a roomId
func (h *RoomHandler) CreateRoomRequestHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("CreateRoomRequestHandler:")

	room := h.roomService.CreateRoom()

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(roomResp{
		RoomId: room.GetID().String(),
		Error:  false,
	})
}

func (h *RoomHandler) JoinRoomRequestHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("JoinRoomRequestHandler:")

	roomId := r.PathValue("id")

	if roomId == "" {
		log.Println("Room id missing in request")
		http.Error(w, "Room cannot be empty", http.StatusBadRequest)
		return
	}

	upgrader := ws.NewUpgrader()

	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Web Socket Upgrade Error:", err)
		http.Error(w, "websocket upgrade failed", http.StatusBadRequest)
		return
	}

	defer wsConn.Close()

	_, exists := h.roomService.GetRoom(roomId)
	if !exists {
		log.Println("Invalid Room ID")
		http.Error(w, "Invalid Room Id", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
