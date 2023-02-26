package utils

import (
	"os"
	"path/filepath"
)

func GetAppPath() string {
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}

	exeDir := filepath.Dir(exePath)

	appPath := filepath.Dir(exeDir)

	return appPath
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
