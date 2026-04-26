package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"webrtc-app/internal/application"
)

type clientResp struct {
	ClientId string `json:"client_id"`
	Error    bool   `json:"error"`
}

type ClientHandler struct {
	clientService *application.ClientService
}

func NewClientHandler(clientService *application.ClientService) *ClientHandler {
	return &ClientHandler{
		clientService: clientService,
	}
}

func (c *ClientHandler) CreateClientHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("CreateClientHandler: ")

	newClient := c.clientService.CreateClient(false)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(clientResp{
		ClientId: newClient.Id.String(),
		Error:    false,
	})
}
