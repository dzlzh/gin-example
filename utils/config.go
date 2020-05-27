package utils

import (
	"fmt"
	"os"

	"github.com/go-ini/ini"
)

var Config *ini.File
var RootPath string

func InitConfig() {
	var err error
	Config, err = ini.Load(RootPath + "/.config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
}
