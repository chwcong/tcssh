package config

import "os"

var (
	WorkDir = ""
)

func Init() {
	WorkDir,err := os.UserHomeDir()
	if err != nil {

	}

}
