package application

import "webrtc-app/internal/domain"

type ClientService struct {
	clients domain.ClientRepository
}

func NewClientService(clients domain.ClientRepository) *ClientService {
	return &ClientService{clients: clients}
}

func (c *ClientService) CreateClient(isHost bool) *domain.Client {
	newClient := c.clients.Create(isHost)
	return newClient
}
