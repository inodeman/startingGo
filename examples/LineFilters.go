package examples

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LineFilters() {
	fmt.Println("**Examples Line Filters**")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)

		if ucl == "EXIT" {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

}
