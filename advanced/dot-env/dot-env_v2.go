package main

import (
	"log"
	"os"

	// If you're very lazy, you can just take advantage of the autoload package which will read in .env on import
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	firstEnvVar := os.Getenv("FIRST_ENV_VAR")
	secondEnvVar := os.Getenv("SECOND_ENV_VAR")

	log.Println("FIRST:", firstEnvVar)
	log.Println("SECOND:", secondEnvVar)
}
