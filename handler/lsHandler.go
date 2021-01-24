package handler

import (
	"github.com/desertbit/grumble"
	"github.com/fatih/color"
	"sync"
	"tcssh/db"
	"tcssh/model"
	"tcssh/util/config"
)

var (
	once sync.Once
	h    Handler
)

type lsHandler struct {
	location *config.CurrentLocation
	color *color.Color
}

func NewLsHandler() Handler {
	once.Do(func() {
		h = &lsHandler{
			location: config.GlobalLocation,
			color: color.New(color.FgBlue),
		}
	})
	return h
}

func (h *lsHandler) Handle(c *grumble.Context) (err error) {
	groups := model.GetGroupByParentID(db.DB,h.location.GetLocation())
	h.printGroup(c,groups)
	return
}

func (h *lsHandler)printGroup(c *grumble.Context,groups []model.Group) {
	nameString := ""
	var names []string
	for i:=0;i< len(groups);i++ {
		names = append(names,groups[i].Name)
		nameString = nameString + h.color.Sprint(groups[i].Name) +"\t"
	}
	c.App.Println(nameString)
}