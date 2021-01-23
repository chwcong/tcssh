package handler

import (
	"github.com/desertbit/grumble"
	"sync"
	"tcssh/model"
)

var (
	once sync.Once
	h    Handler
)

type lsHandler struct {
	currentId int
}

func NewLsHandler() Handler {
	once.Do(func() {
		h = &lsHandler{
			currentId: 0,
		}
	})
	return h
}

func (h *lsHandler) Handle(c *grumble.Context) (err error) {
	model.GetGroupByParentID(h.currentId)
	return
}