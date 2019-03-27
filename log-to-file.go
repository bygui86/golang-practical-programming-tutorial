package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("./log-to-file.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Panicln("Error opening log file:", err.Error)
	}
	defer file.Close()

	log.SetOutput(file)
	log.Println("Logging to a file in go ;)")
}
