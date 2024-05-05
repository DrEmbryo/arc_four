package main

import (
	"fmt"
	"os"

	ArcFour "github.com/DrEmbryo/arc_four/lib"
)


func main() {
	key := "someSecretKey"
	source, err := os.ReadFile("./cmd/test.txt")
	if err != nil {
		panic(err)
	}

	enc := &ArcFour.RC4{}
	
	enc.Init(key)
	encrypted := enc.Encrypt(string(source))
	enc.Init(key)
	decrypted := enc.Decrypt(encrypted)
	fmt.Println(decrypted)
}