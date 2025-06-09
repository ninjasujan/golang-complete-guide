package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {

	// sha1 and crypto

	hash := sha1.New()

	fmt.Println("[Hash sh1]", hash)

	sha1Hash := hash.Sum([]byte("some string"))

	fmt.Println("[String sha1]", sha1Hash)

	fmt.Println("[converted string]", string(sha1Hash))

}
