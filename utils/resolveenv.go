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
