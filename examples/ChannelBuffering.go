package examples

import "fmt"

func ChannelBuffering() {
	fmt.Println("**Examples Channel Buffering**")

	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

}
