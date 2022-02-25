package main

import (
	"log"

	logger "github.com/DanielPickens/logger"
)

func main() {
	clilog.Env = "STANDARD_APP_LOG"
	clilog.SetOutput()

	log.Printf("[INFO] any function")
}
