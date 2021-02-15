package cmd

import (
	"github.com/desertbit/grumble"
	"tcssh/handler"
)

var sshCmd = &grumble.Command{
	Name: "ssh",
	Help: "ssh connect to the host.",
	Flags: func(f *grumble.Flags) {
	},
	Args: func(a *grumble.Args) {
		a.String("host", "host name")
	},
	Run: func(c *grumble.Context) error {
		h := handler.NewSSHHandler()
		err := h.Handle(c)
		return err
	},
}