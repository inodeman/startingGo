package main

import "fmt"


func values() {
	fmt.Println("**Example Values**")
	fmt.Println("go" + "lang")
	fmt.Println("1+1", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	fmt.Println(true && true)
	fmt.Println(true || false)
	fmt.Println(!true)
}

func helloWorld() {
	fmt.Println("**Examples Hello World**")
	fmt.Println("Hello, World!")
}


func main() {
	helloWorld()
	values()

}
