package main

import (
	"github.com/desertbit/grumble"
	"tcssh/cmd"
	"tcssh/util/config"
)

func main() {
	config.Init()
	app := cmd.NewApp()
	grumble.Main(app)
}
