package cmd

import (
	"github.com/desertbit/grumble"
	"github.com/fatih/color"
)


func NewApp() (app *grumble.App) {
	app = grumble.New(&grumble.Config{
		Name:                  "tcssh",
		Description:           "An terminal ssh client",
		HistoryFile:           "/tmp/tcssh.hist",
		Prompt:                "tc » ",
		PromptColor:           color.New(color.FgGreen, color.Bold),
		HelpHeadlineColor:     color.New(color.FgGreen),
		HelpHeadlineUnderline: true,
		HelpSubCommands:       true,

		Flags: func(f *grumble.Flags) {
		},
	})
	app.AddCommand(lsCmd)
	return app
}