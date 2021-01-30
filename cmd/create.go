package cmd

import (
	"github.com/desertbit/grumble"
	"tcssh/handler"
)

var createCmd = &grumble.Command{
	Name: "create",
	Help: "create an new node",
	Args: func(a *grumble.Args) {
		a.String("address", "remote address [user@]hostname")
	},
	Flags: func(f *grumble.Flags) {
		f.String("n", "name", "", "set the name of this host")
		f.String("d", "description", "", "set the description  of this host")
		f.Int("p", "port", 22, "set the port of ssh host")
	},
	Run: func(c *grumble.Context) error {
		h := handler.NewCreateHandler()
		err := h.Handle(c)
		return err
	},
}
