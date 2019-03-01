package examples

import "fmt"

func Channels() {
	fmt.Println("**Examples Channels**")
	messages := make(chan string)
	go func() { messages <- "ping" }()
	msg := <-messages
	fmt.Println(msg)
}
