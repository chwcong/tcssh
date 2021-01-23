package config

import (
	"log"
	"os"
	"path"
)

var (
	WorkDir         = ".tcssh"
	HomeDir         = ""
	WorkPath        = ""
	HistoryFileName = "tcssh.hist"
)

func InitWorkPath() {
	dir, err := os.UserHomeDir()
	if err != nil {
		// TODO 支持更多的目录
		// TODO 实现日志组件
		log.Fatal(err)
	}
	HomeDir = dir
	WorkPath = path.Join(HomeDir, WorkDir)
	// test if the work path is exist ,if not ,create the dir
	_, err = os.Stat(WorkPath)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(WorkPath, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}
	}
}
