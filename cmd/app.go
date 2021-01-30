package cmd

import (
	"github.com/desertbit/grumble"
	"github.com/fatih/color"
	"path"
	"tcssh/util/config"
)

func NewApp() (app *grumble.App) {
	app = grumble.New(&grumble.Config{
		Name:                  "tcssh",
		Description:           "An terminal ssh client",
		HistoryFile:           path.Join(config.WorkPath, config.HistoryFileName),
		Prompt:                "tc » ",
		PromptColor:           color.New(color.FgGreen, color.Bold),
		HelpHeadlineColor:     color.New(color.FgGreen),
		HelpHeadlineUnderline: true,
		HelpSubCommands:       true,
	})
	app.AddCommand(lsCmd)
	app.AddCommand(cdCmd)
	app.AddCommand(mkdirCmd)
	app.AddCommand(rmCmd)
	app.AddCommand(createCmd)
	return app
}
