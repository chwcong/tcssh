package main

import (
	"github.com/desertbit/grumble"
	"tcssh/cmd"
	"tcssh/db"
	"tcssh/util/config"
)

func main() {
	Init()
	app := cmd.NewApp()
	grumble.Main(app)
}

func Init() {
	config.InitWorkPath()
	db.InitDB()
}
