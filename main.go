package main

import (
	"log"
	"os"
)

func main() {
	log.Print("logging in go!")
	file, err := os.OpenFile("/home/mcrap/info.log",
		os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.Print("Logging to a file in go!")
}
