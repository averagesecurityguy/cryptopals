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
	"strings"
    ca "github.com/averagesecurityguy/cryptanalysis"
)


const LEN = 144


func check(e error) {
	if e != nil {
		fmt.Printf("Error: %s\n", e.Error())
		os.Exit(0)
	}
}


func oracle(data, key []byte) []byte {
    pad := "Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK"
    decoded := ca.DecodeB64Str(pad)
    data = append(data, decoded...)

	encrypted, err := ca.EncryptEcb(data, key)
    check(err)

	return encrypted
}


func get_block_size(key []byte) int {
	data := make([]byte, 0)
	enc := oracle(data, key)
	length := len(enc)

	for i := 1; i < 64; i++ {
		data = make([]byte, i)
		enc = oracle(data, key)
		diff := len(enc) - length

		if diff > 0 {
			return diff
		}
	}

	return 0
}


func is_ecb(key []byte, bs int) bool {
	ecb := false

	for i := 0; i <= 64; i++ {
		data := make([]byte, i)
		enc := oracle(data, key)
		score := ca.ScoreEcb(enc, bs)

		if score < float64(1) {
			ecb = true
		}
	}

	return ecb
}


func build_dict(data, key []byte) map[[LEN]byte]byte {
	d := make(map[[LEN]byte]byte)
	var block [LEN]byte

	for i := 0; i < 256; i++ {
		plain := append(data, byte(i))
		enc := oracle(plain, key)

		copy(block[:], enc[:LEN])
		d[block] = byte(i)
	}

	return d
}


func main() {
    fmt.Println("Cryptopals Set 2")
    fmt.Println("================")

    fmt.Println("Challenge 4")
    fmt.Println("-----------")

	key, err := ca.RandomBytes(16)
	check(err)

    // Find the block size of the oracle.
	bs := get_block_size(key)
	fmt.Printf("Block Size: %d\n", bs)

	// Is the oracle ECB?
	if is_ecb(key, bs) {
		fmt.Println("Oracle is ECB")
	} else {
		fmt.Println("Oracle is not ECB")
	}

	// Use a long string of repeating characters to identify the next byte in
	// the plaintext. Make the list of repeating characters smaller as more
	// plaintext bytes are identified.
	a := []byte(strings.Repeat("A", LEN - 1))
	g := []byte("")
	var block [LEN]byte

	for i:= 0; i < LEN - 1; i++ {
		p := append(a, g...)
		dict := build_dict(p, key)
		enc := oracle(a, key)

	    copy(block[:], enc[:LEN])
		char := dict[block]
		fmt.Print(string(char))

		g = append(g, char)
		a = a[1:]
	}

	fmt.Println("")
}
