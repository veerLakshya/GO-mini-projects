package domain

type RoomRepository interface {
	Get(id string) (*Room, bool)
	Add(room *Room)
	Delete(id string) bool
}
