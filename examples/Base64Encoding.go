package examples

import (
	"encoding/base64"
	"fmt"
)

func Base64Encoding() {
	fmt.Println("**Examples Base64 Encoding**")
	data := "abc123!?$*&()'-=@~"

	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	sDec, _ := base64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	uEnc := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)

	uDec, _ := base64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))

}
