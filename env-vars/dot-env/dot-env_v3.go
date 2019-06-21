package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("dot-env_v3_one.env", "dot-env_v3_two.env")
	if err != nil {
		log.Fatal("Error loading .env files")
	}

	firstEnvVar := os.Getenv("FIRST_ENV_VAR")
	secondEnvVar := os.Getenv("SECOND_ENV_VAR")
	thirdEnvVar := os.Getenv("THIRD_ENV_VAR")
	fourthEnvVar := os.Getenv("FOURTH_ENV_VAR")

	log.Println("FIRST:", firstEnvVar)
	log.Println("SECOND:", secondEnvVar)
	log.Println("THIRD:", thirdEnvVar)
	log.Println("FOURTH:", fourthEnvVar)
}
