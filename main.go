package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
)

func main() {

	var input string
	flag.StringVar(&input, "i", "fubu for us by us", "a string")

	var height int
	flag.IntVar(&height, "n", 12, "a number")

	flag.Parse()

	h := sha256.New()

	h.Write([]byte(input))

	fmt.Printf("%x", h.Sum(nil))

}
