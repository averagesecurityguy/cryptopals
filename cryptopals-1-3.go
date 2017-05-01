/*
Copyright (c) 2017, AverageSecurityGuy
# All rights reserved.

Complete challenge 3 from Cryptopals set 1.
http://cryptopals.com/sets/1

Usage:

$ go run cryptopals-1-3.go
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
    /*
    Take a hex string and determine the single byte key used to encrypt it
    with XOR. This is equivalent to the Ceaser Cipher and can be cracked using
    letter frequency analysis.
    */

    fmt.Println("Cryptopals Set 1")
    fmt.Println("================")

    fmt.Println("Challenge 3")
    fmt.Println("-----------")
    enc := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
    dec := ca.DecodeHexStr(enc)
    _, msg := ca.BreakSingleByteXor(dec)

    fmt.Println(msg)
}
