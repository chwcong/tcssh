package cmd

import (
	"github.com/desertbit/grumble"
	"tcssh/handler"
)

var mkdirCmd = &grumble.Command{
	Name: "mkdir",
	Help: "make an new group",
	Args: func(a *grumble.Args) {
		a.String("group", "group name")
	},
	Flags: func(f *grumble.Flags) {
		f.String("d", "description", "", "set the description of this group")
	},
	Run: func(c *grumble.Context) error {
		h := handler.NewMkdirHandler()
		err := h.Handle(c)
		return err
	},
}
