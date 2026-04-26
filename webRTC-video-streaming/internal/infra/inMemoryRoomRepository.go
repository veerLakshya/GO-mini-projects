package infra

import (
	"sync"

	"webrtc-app/internal/domain"
)

type InMemoryRoomRepository struct {
	mu    sync.RWMutex
	rooms map[string]*domain.Room
}

var _ domain.RoomRepository = (*InMemoryRoomRepository)(nil)

func NewInMemoryRoomRepository() *InMemoryRoomRepository {
	return &InMemoryRoomRepository{
		rooms: make(map[string]*domain.Room),
	}
}

func (r *InMemoryRoomRepository) Get(id string) (*domain.Room, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	room, ok := r.rooms[id]
	return room, ok
}

func (r *InMemoryRoomRepository) Add(room *domain.Room) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.rooms[room.GetID().String()] = room
}

func (r *InMemoryRoomRepository) Delete(id string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.rooms[id]; !exists {
		return false
	}

	delete(r.rooms, id)
	return true
}
