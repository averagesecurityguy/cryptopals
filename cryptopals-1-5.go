/*
Copyright (c) 2017, AverageSecurityGuy
# All rights reserved.

Complete challenge 5 from Cryptopals set 1.
http://cryptopals.com/sets/1

Usage:

$ go run cryptopals-1-5.go
*/

package main

import (
    "os"
    "fmt"
)

import ca "github.com/averagesecurityguy/cryptanalysis"


func check(e error) {
	if e != nil {
		fmt.Printf("Error: %s\n", e.Error())
		os.Exit(0)
	}
}


func main() {
    fmt.Println("Cryptopals Set 1")
    fmt.Println("================")

    fmt.Println("Challenge 5")
    fmt.Println("-----------")

    plain := []byte(`Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`)
    key := []byte("ICE")
    cipher := ca.EncryptXor(plain, key)

    fmt.Printf("%x\n", cipher)
}
