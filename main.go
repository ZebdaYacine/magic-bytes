package main

import (
	"fmt"

	"github.com/ZebdaYacine/magic-bytes/magicbytes/magicbytes"
)

func main() {
	// base64 representation of "Hello, World!"
	str := "SGVsbG8sIFdvcmxkIQ=="
	fmt.Println(magicbytes.SaveBase64ToFile(str, "greeting"))
}
