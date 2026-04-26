package domain

type ClientRepository interface {
	Get(id string) (*Client, bool)
	//creates and adds client to the inmemory clients
	Create(isHost bool) *Client
	Delete(id string) bool
}
