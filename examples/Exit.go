package examples

import (
	"fmt"
	"os"
)

func Exit() {
	fmt.Println("**Examples Exit**")

	defer fmt.Println("This will not be printed")
	os.Exit(3)

}
