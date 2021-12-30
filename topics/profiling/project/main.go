package main

import (
	"log"
	"os"

	"github.com/learning_golang/topics/profiling/project/service"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	log.SetOutput(os.Stdout)
}

func main() {
	service.Start()
}
