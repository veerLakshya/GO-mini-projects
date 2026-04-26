package application

import "webrtc-app/internal/domain"

type RoomService struct {
	store domain.RoomRepository
}

func NewRoomService(store domain.RoomRepository) *RoomService {
	return &RoomService{store: store}
}

func (s *RoomService) CreateRoom() *domain.Room {
	room := domain.NewRoom()
	s.store.Add(room)
	return room
}

func (s *RoomService) GetRoom(id string) (*domain.Room, bool) {
	return s.store.Get(id)
}
