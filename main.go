package main

import (
	"tcssh/db"
	"tcssh/util/config"
)

func main() {
	Init()
	//app := cmd.NewApp()
	//grumble.Main(app)
}

func Init() {
	config.InitWorkPath()
	db.InitDB()
}
