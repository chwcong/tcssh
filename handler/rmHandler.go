package handler

import (
	"github.com/desertbit/grumble"
	"sync"
	"tcssh/util/config"
)

var (
	rmOnce sync.Once
	rmHan    Handler
)

type rmHandler struct {
	location *config.CurrentLocation
}

func NewRmHandler() Handler {
	rmOnce.Do(func() {
		rmHan = &rmHandler{
			location: config.GlobalLocation,
		}
	})
	return rmHan
}

func (h *rmHandler) Handle(c *grumble.Context) (err error) {
	return
}
