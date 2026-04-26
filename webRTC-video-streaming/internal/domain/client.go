package domain

import "github.com/google/uuid"

type Client struct {
	Id     uuid.UUID
	RoomId string
	Host   bool
}

func NewClient(isHost bool) *Client {
	return &Client{
		Id:   uuid.New(),
		Host: isHost,
	}
}
