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

// Init init config and db
func Init() {
	config.InitWorkPath()
	config.InitGlobalLocation()
	db.InitDB()
}
