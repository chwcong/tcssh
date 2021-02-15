package handler

import (
	"fmt"
	"github.com/desertbit/grumble"
	"sync"
	"tcssh/db"
	"tcssh/model"
	"tcssh/util/config"
)

var (
	sshOnce sync.Once
	sshHan    Handler
)

type sshHandler struct {
	location *config.CurrentLocation
}

func NewSSHHandler() Handler {
	rmOnce.Do(func() {
		rmHan = &sshHandler{
			location: config.GlobalLocation,
		}
	})
	return rmHan
}

func (h *sshHandler) Handle(c *grumble.Context) (err error) {
	hostName := c.Args.String("host")
	node,err := model.GetNodeOfCurrentIdByName(db.DB,h.location.GetLocation(),hostName)
	fmt.Println(node)
	return
}

