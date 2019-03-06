package examples

import (
	"fmt"
	"os"
)

func Panic() {
	fmt.Println("**Examples Panic**")
	panic("a problem")
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
