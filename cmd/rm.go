package cmd

import (
	"github.com/desertbit/grumble"
	"tcssh/handler"
)

var rmCmd = &grumble.Command{
	Name: "rm",
	Help: "remove dentry.if remove a group,all nodes and groups of this group will be remove",
	Flags: func(f *grumble.Flags) {
	},
	Run: func(c *grumble.Context) error {
		h := handler.NewRmHandler()
		err := h.Handle(c)
		return err
	},
}