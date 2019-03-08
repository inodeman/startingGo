package examples

import (
	"crypto/sha1"
	"fmt"
)

func Sha1Hashes() {
	fmt.Println("**Examples Sha1 Hashes**")
	s := "sha1 this string"
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	fmt.Println(s)
	fmt.Printf("%x\n", bs)

}
