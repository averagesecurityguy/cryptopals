/*
Copyright (c) 2017, AverageSecurityGuy
# All rights reserved.

Complete challenge 6 from Cryptopals set 1.
http://cryptopals.com/sets/1

Usage:

$ go run cryptopals-1-6.go
*/

package main

import (
    "os"
    "fmt"
    "bufio"
    "bytes"
)

import ca "github.com/averagesecurityguy/cryptanalysis"


func check(e error) {
	if e != nil {
		fmt.Printf("Error: %s\n", e.Error())
		os.Exit(0)
	}
}


func get_data(filename string) []byte {
    file, err := os.Open(filename)
    check(err)

    var buffer bytes.Buffer
    scan := bufio.NewScanner(file)

	for scan.Scan() {
        buffer.WriteString(scan.Text())
    }

    return ca.DecodeB64Str(buffer.String())
}


func main() {
    fmt.Println("Cryptopals Set 1")
    fmt.Println("================")

    fmt.Println("Challenge 6")
    fmt.Println("-----------")

    crypt := get_data("data/6.txt")
    size, err := ca.KeyLength(crypt)
    check(err)

    chunks := ca.Chunk(crypt, size)
    trans := ca.Transpose(chunks)

    var key []byte
    for _, t := range trans {
        _, k, _ := ca.BreakSingleByteXor(t)
        key = append(key, k)
    }

    fmt.Println(string(key))
}
