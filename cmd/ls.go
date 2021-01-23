package cmd

import (
	"github.com/desertbit/grumble"
	"tcssh/handler"
)

var lsCmd = &grumble.Command{
	Name: "ls",
	Help: "list node and group",
	Flags: func(f *grumble.Flags) {
	},
	Run: func(c *grumble.Context) error {
		h := handler.NewLsHandler()
		err := h.Handle(c)
		return err
	},
}
