package Logger

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	CLI_LOG := "CLI_LOG"
	CLI_LOG_PATH := "CLI_LOG_PATH"
	log.SetOutput(os.Stdout)
	log.Println("This is a standard message")
	log.SetOutput(ioutil.Discard)
	x:= os.Getenv(CLI_LOG)
	y:= os.Getenv(CLI_LOG_PATH)
	log.Println(x)
	log.Println(y)
}
