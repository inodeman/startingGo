package examples

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkfileContents() {
	dat, err := ioutil.ReadFile("/tmp/dat1")
	check(err)
	fmt.Println("content of /tmp/dat1")
	fmt.Print(string(dat))

	dat, err = ioutil.ReadFile("/tmp/dat2")
	check(err)
	fmt.Println("content of /tmp/dat1")
	fmt.Print(string(dat))

}

func WritingFiles() {
	fmt.Println("**Examples Writing Files**")

	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	f, err := os.Create("/tmp/dat2")
	check(err)
	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()

	checkfileContents()

}
