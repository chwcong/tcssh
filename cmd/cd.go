package cmd

import (
	"github.com/desertbit/grumble"
	"tcssh/handler"
)

var cdCmd = &grumble.Command{
	Name: "cd",
	Help: "cd change current group",
	Flags: func(f *grumble.Flags) {
	},
	Run: func(c *grumble.Context) error {
		h := handler.NewLsHandler()
		err := h.Handle(c)
		return err
	},
}
