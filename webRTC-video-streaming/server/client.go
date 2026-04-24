package server

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	Id   uuid.UUID
	Host bool
	Conn *websocket.Conn
}

func NewClient(isHost bool, conn *websocket.Conn) *Client {
	return &Client{
		Id:   uuid.New(),
		Host: isHost,
		Conn: conn,
	}
}
