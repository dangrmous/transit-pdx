package main

import (
	"github.com/dangrmous/transit-pdx/client"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	appClient := client.New()
}
