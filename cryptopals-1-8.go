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
)

import ca "github.com/averagesecurityguy/cryptanalysis"

func check(e error) {
	if e != nil {
		fmt.Printf("Error: %s\n", e.Error())
		os.Exit(0)
	}
}


func get_data(filename string) [][]byte {
    file, err := os.Open(filename)
    check(err)

    var data [][]byte
    scan := bufio.NewScanner(file)

	for scan.Scan() {
        temp := ca.DecodeHexStr(scan.Text())
        data = append(data, temp)
    }

    return data
}


func main() {
    fmt.Println("Cryptopals Set 1")
    fmt.Println("================")

    fmt.Println("Challenge 8")
    fmt.Println("-----------")

    block_size := 16
    crypts := get_data("data/8.txt")
    low := len(crypts[0]) / block_size
    msg := ""

    for _, crypt := range crypts {
        chunks := ca.Chunk(crypt, block_size)
        temp := make(map[string] int)

        for _, c := range chunks {
            temp[string(c)] = 0
        }

        // Assume that the crypt with the least unique chunks is our ecb
        // encrypted data.
        if len(temp) < low {
            low = len(temp)
            msg = string(crypt)
        }
    }

    for _, c := range ca.Chunk([]byte(msg), 16) {
        fmt.Printf("%x\n", c)
    }
}
