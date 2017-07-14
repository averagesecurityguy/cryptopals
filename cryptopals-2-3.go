/*
Copyright (c) 2017, AverageSecurityGuy
# All rights reserved.

Complete challenge 3 from Cryptopals set 2.
http://cryptopals.com/sets/2

Usage:

$ go run cryptopals-2-3.go
*/

package main

import (
    "os"
    "fmt"
    ca "github.com/averagesecurityguy/cryptanalysis"
)


func check(e error) {
	if e != nil {
		fmt.Printf("Error: %s\n", e.Error())
		os.Exit(0)
	}
}


func encrypt(data []byte) []byte {
    var encrypted []byte

    // Generate a random pad size of 5-10 characters and append it to the
	// data.
	size, err := ca.RandomIntRange(5, 11)
	check(err)

    pad := make([]byte, size)
    data = append(pad, data...)
    data = append(data, pad...)

	// Generate a new key and a random choice. Encrypt the data using the
	// key and either ECB or CBC depending on the choice.
	key, err := ca.RandomBytes(16)
	check(err)

    choice, err := ca.RandomIntRange(0, 2)
	check(err)

    if choice == 0 {
		fmt.Printf("ECB - ")
        enc, err := ca.EncryptEcb(data, key)
        check(err)

		encrypted = enc
    } else {
		fmt.Printf("CBC - ")
        iv, err := ca.RandomBytes(16)
		check(err)

        enc, err := ca.EncryptCbc(data, key, iv)
        check(err)

		encrypted = enc
    }

    return encrypted
}


func main() {
    fmt.Println("Cryptopals Set 2")
    fmt.Println("================")

    fmt.Println("Challenge 3")
    fmt.Println("-----------")

	// The chosen plaintext should yield at least four identical blocks no
	// matter what padding is chosen. We can score the encrypted value to see
    // if it is ECB.
    plain := []byte("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")

	for i := 0; i < 100; i++ {
		encrypted := encrypt(plain)
		score := ca.ScoreEcb(encrypted, 16)

		if score < float64(1) {
			fmt.Println("Guess ECB")
		} else {
			fmt.Println("Guess CBC")
		}
	}
}
