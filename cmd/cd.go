package cmd

import (
	"github.com/desertbit/grumble"
	"tcssh/handler"
)

var cdCmd = &grumble.Command{
	Name: "cd",
	Help: "cd change current group",
	Args: func(a *grumble.Args) {
		a.String("group", "group name")
	},
	Run: func(c *grumble.Context) error {
		h := handler.NewCdHandler()
		err := h.Handle(c)
		return err
	},
}
