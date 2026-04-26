package domain

import (
	"sync"
	"time"

	"github.com/google/uuid"
)

type Room struct {
	ID        uuid.UUID
	Clients   []Client
	CreatedAt time.Time
	UpdatedAt time.Time

	Mutex sync.RWMutex
}

func NewRoom() *Room {
	now := time.Now()
	return &Room{
		ID:        uuid.New(),
		Clients:   make([]Client, 0),
		CreatedAt: now,
		UpdatedAt: now,
		Mutex:     sync.RWMutex{},
	}
}

func (r *Room) GetID() uuid.UUID {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	return r.ID
}

func (r *Room) GetClients() []Client {
	r.Mutex.RLock()
	defer r.Mutex.RUnlock()

	return r.Clients
}

func (r *Room) AddClient(c Client) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	r.Clients = append(r.Clients, c)
	r.UpdatedAt = time.Now().UTC()
}

func (r *Room) RemoveClient(id uuid.UUID) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	for i, c := range r.Clients {
		if c.Id == id {
			r.Clients = append(r.Clients[:i], r.Clients[i+1:]...)
			r.UpdatedAt = time.Now().UTC()
			return
		}
	}
}
