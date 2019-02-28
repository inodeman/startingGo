package examples

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func MultipleReturnValues() {
	fmt.Println("**Examples Multiple Return Values**")
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}
