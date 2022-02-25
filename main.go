package main

import (
	"log"

	"Logger/clilog"
)

func main() {
	clilog.Env = "STANDARD_APP_LOG"
	clilog.SetOutput()

	log.Printf("[INFO] any function")
}
