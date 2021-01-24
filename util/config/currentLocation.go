package config

import "sync"

// GlobalLocation hold the global location
var GlobalLocation *CurrentLocation

// CurrentLocation is the struct of location
type CurrentLocation struct {
	parentId int
	rwLock   sync.RWMutex
}

// InitGlobalLocation init the GlobalLocation of the app
func InitGlobalLocation() {
	GlobalLocation = &CurrentLocation{
		parentId: 0,
	}
}

func (l *CurrentLocation) GetLocation() int {
	return l.parentId
}

func (l *CurrentLocation) SetLocation(id int) {
	l.parentId = id
}
