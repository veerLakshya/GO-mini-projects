package infra

import (
	"webrtc-app/internal/domain"
)

type InMemoryClientRepository struct {
	clients map[string]*domain.Client
}

var _ domain.ClientRepository = (*InMemoryClientRepository)(nil)

func NewInMemoryClientRepository() *InMemoryClientRepository {
	return &InMemoryClientRepository{
		clients: make(map[string]*domain.Client),
	}
}

func (c *InMemoryClientRepository) Get(id string) (*domain.Client, bool) {
	client, ok := c.clients[id]
	return client, ok
}

func (c *InMemoryClientRepository) Create(isHost bool) *domain.Client {
	client := domain.NewClient(isHost)
	idStr := client.Id.String()
	c.clients[idStr] = client
	return client
}

func (c *InMemoryClientRepository) Delete(id string) bool {

	_, exists := c.clients[id]
	if !exists {
		return false
	}

	delete(c.clients, id)
	return true
}
