package env

import (
	"log"
	"os"
)

func GetApiKey() string {
	key := os.Getenv("DESIGNPILOT_API_KEY")
	if key == "" {
		log.Fatal("In order to use the CLI, you need to set the `DESIGNPILOT_API_KEY` environment variable. You can generate a key through the web app at designpilot.ai/profile")
	}
	return key
}
