package utils

import "os"

func GetDevPath() string {
	return os.Getenv("DEV_PATH")
}
