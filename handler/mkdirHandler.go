package handler

import (
	"github.com/desertbit/grumble"
	"sync"
	"tcssh/db"
	"tcssh/model"
	"tcssh/util/config"
	constant "tcssh/util/const"
)

var (
	mkdirOnce sync.Once
	mkdirHan    Handler
)

type mkdirHandler struct {
	location *config.CurrentLocation
}

func NewMkdirHandler() Handler {
	mkdirOnce.Do(func() {
		mkdirHan = &mkdirHandler{
			location: config.GlobalLocation,
		}
	})
	return mkdirHan
}

func (m *mkdirHandler) Handle(c *grumble.Context) (err error) {
	name := c.Args.String("group")
	description := c.Flags.String("description")
	dentry := &model.Dentry{
		ParentID:    m.location.GetLocation(),
		Name:        name,
		Description: description,
		NodeId:      0,
		Type:        constant.DENTRY_TYPE_GROUP,
	}
	err = model.CreateDentry(db.DB,dentry)
	return
}