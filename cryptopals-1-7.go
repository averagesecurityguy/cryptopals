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
    "crypto/aes"
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

    fmt.Println("Challenge 7")
    fmt.Println("-----------")

    block_size := 16
    ecb, err := aes.NewCipher([]byte("YELLOW SUBMARINE"))
    check(err)

    ciphertext := get_data("data/7.txt")
    plaintext := make([]byte, 0)
    chunks := ca.Chunk(ciphertext, block_size)

    for _, chunk := range chunks {
        chunk = ca.PadPkcs7(chunk, block_size)
        temp := make([]byte, block_size)

        ecb.Decrypt(temp, chunk)

        plaintext = append(plaintext, temp...)
    }

    fmt.Println(string(plaintext))
}
