package main

import (
	"log"
	"os"
	"strconv"
)

func GetStringEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	log.Println("Environment variable", key, "not found, using fallback", fallback)
	return fallback
}

func GetIntEnv(key string, fallback int) int {
	if strValue, ok := os.LookupEnv(key); ok {
		value, err := strconv.Atoi(strValue)
		if err != nil {
			log.Panicln("Error reading the environment variable", key, ", not an int!")
		}
		return value
	}
	log.Println("Environment variable", key, "not found, using fallback", fallback)
	return fallback
}

func GetBoolEnv(key string, fallback bool) bool {
	if strValue, ok := os.LookupEnv(key); ok {
		value, err := strconv.ParseBool(strValue)
		if err != nil {
			log.Panicln("Error reading the environment variable", key, ", not an boolean!")
		}
		return value
	}
	log.Println("Environment variable", key, "not found, using fallback", fallback)
	return fallback
}
