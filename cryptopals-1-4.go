/*
Copyright (c) 2017, AverageSecurityGuy
# All rights reserved.

Complete challenge 4 from Cryptopals set 1.
http://cryptopals.com/sets/1

Usage:

$ go run cryptopals-1-4.go
*/

package main

import (
    "os"
    "fmt"
    "bufio"
)

import ca "github.com/averagesecurityguy/cryptanalysis"


func check(e error) {
	if e != nil {
		fmt.Printf("Error: %s\n", e.Error())
		os.Exit(0)
	}
}


func open(filename string) *os.File {
	data, err := os.Open(filename)
	check(err)

	return data
}


func main() {
    fmt.Println("Cryptopals Set 1")
    fmt.Println("================")

    fmt.Println("Challenge 4")
    fmt.Println("-----------")
    crypts := open("data/4.txt")
    cscan := bufio.NewScanner(crypts)

	for cscan.Scan() {
        enc := ca.DecodeHexStr(cscan.Text())
		score, key, msg := ca.BreakSingleByteXor(enc)

        if score < 35.0 {
            fmt.Printf("%f: %d: %q\n", score, key, msg)
        }
    }
}
