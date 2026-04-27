package server

import (
	"sync"

	"github.com/gorilla/websocket"
)

type RoomsMap struct {
	Mutex sync.RWMutex
	Map   map[string]*Room
}

func (r *RoomsMap) Init() {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	r.Map = make(map[string]*Room)
	r.Map["testRoom"] = NewRoom()
}

// Get room by a specific id
func (r *RoomsMap) Get(roomId string) (*Room, bool) {
	r.Mutex.RLock()
	defer r.Mutex.RUnlock()

	room, ok := r.Map[roomId]

	return room, ok
}

// Add a new room in rooms
func (r *RoomsMap) AddRoom(newRoom *Room) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	roomIdStr := newRoom.GetID().String()

	r.Map[roomIdStr] = newRoom
}

// delete a room from rooms
func (r *RoomsMap) DeleteRoom(roomId string) bool {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	_, exists := r.Map[roomId]

	if !exists {
		return false
	}

	delete(r.Map, roomId)
	return true
}

func (r *RoomsMap) InsertIntoRoom(roomId string, host bool, conn *websocket.Conn) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	c := NewClient(host, conn)

	room, ok := r.Map[roomId]
	if !ok {
		// room not found
		return
	}

	room.AddClient(*c)
}
