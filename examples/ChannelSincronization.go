package examples

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func ChannelSincronization() {
	fmt.Println("**Examples Channel Sincronization**")

	done := make(chan bool, 1)
	go worker(done)

	<-done
}
