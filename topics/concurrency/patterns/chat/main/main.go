package main

import (
	"fmt"
	"github.com/learning_golang/topics/concurrency/patterns/chat"
	"os"
	"os/signal"
)

func main() {
	rs := chat.New()

	signCh := make(chan os.Signal, 1)
	signal.Notify(signCh, os.Interrupt)

	// Wait to stop
	<-signCh
	fmt.Println("Shutdown Down started")
	rs.Close()
	fmt.Println("Shutdown Down Completed")
}
