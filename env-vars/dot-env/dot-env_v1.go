package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	firstEnvVar := os.Getenv("FIRST_ENV_VAR")
	secondEnvVar := os.Getenv("SECOND_ENV_VAR")

	log.Println("FIRST:", firstEnvVar)
	log.Println("SECOND:", secondEnvVar)
}
