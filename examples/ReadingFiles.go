package examples

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
)

func checkRead(e error) {
	if e != nil {
		panic(e)
	}
}

func createFiles() {
	cmd := exec.Command("echo", "hello\ngo")

	testFile, err := os.Create("/tmp/dat")
	if err != nil {
		panic(err)
	}
	defer testFile.Close()

	cmd.Stdout = testFile

	err = cmd.Start()
	if err != nil {
		panic(err)
	}
	cmd.Wait()
}

func ReadingFiles() {
	fmt.Println("**Examples Reading Files**")

	createFiles()
	dat, err := ioutil.ReadFile("/tmp/dat")
	checkRead(err)
	fmt.Print(string(dat))

	f, err := os.Open("/tmp/dat")
	checkRead(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	checkRead(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	o2, err := f.Seek(6, 0)
	checkRead(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	checkRead(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	o3, err := f.Seek(6, 0)
	checkRead(err)

	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	checkRead(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, 0)
	checkRead(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	checkRead(err)
	fmt.Printf("5 bytes: %s\n", string(b4))
	f.Close()

}
