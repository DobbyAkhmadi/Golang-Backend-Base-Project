package seeders

import (
	"os"
	"strings"
)

func MainSeed(modules string) {
	configEnv := strings.ToLower(os.Getenv("SERVER_MODE"))
	if configEnv == "dev" {
		if modules == "region" {

		}
	}
}
