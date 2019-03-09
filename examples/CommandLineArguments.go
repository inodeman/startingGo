package examples

import (
	"fmt"
	"os"
)

func CommandLineArguments() {
	argsWithProg := os.Args
	argsWithouProg := os.Args[1:]
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithouProg)
	fmt.Println(arg)
}
