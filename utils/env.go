package utils

import "os"

const (
	Dev = iota
	Prod
)

func ResolveEnv() int {
	if env := os.Getenv("ENV"); env == "dev" {
		return Dev
	}
	return Prod
}

func GetDevPath() string {
	return os.Getenv("DEV_PATH")
}

func ResolveAppRoot() string {
	if ResolveEnv() == Dev {
		return GetDevPath()
	}
	return GetAppPath()
}
