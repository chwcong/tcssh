package handler

import (
	"errors"
	"github.com/desertbit/grumble"
	"sync"
	"tcssh/db"
	"tcssh/model"
	"tcssh/util/config"
)

var (
	cdOnce sync.Once
	cdHan    Handler
)

type cdHandler struct {
	location *config.CurrentLocation
	printer  *printer
}

func NewCdHandler() Handler {
	cdOnce.Do(func() {
		cdHan = &cdHandler{
			location: config.GlobalLocation,
		}
	})
	return cdHan
}

func (h *cdHandler) Handle(c *grumble.Context)  error {
	name := c.Args.String("group")
	if name == "" {
		return errors.New("please input group name")
	}
	dentry,err := model.GetDentryIdByNameAndParentID(db.DB,name, h.location.GetLocation())
	if err != nil {
		return err
	}
	h.location.SetLocation(dentry.ID)
	return nil
}
