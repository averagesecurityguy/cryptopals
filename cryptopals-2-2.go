/*
Copyright (c) 2017, AverageSecurityGuy
# All rights reserved.

Complete challenge 2 from Cryptopals set 2.
http://cryptopals.com/sets/2

Usage:

$ go run cryptopals-2-2.go
*/

package main

import (
    "os"
    "fmt"
    "bytes"
    "bufio"
    ca "github.com/averagesecurityguy/cryptanalysis"
)


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
    fmt.Println("Cryptopals Set 2")
    fmt.Println("================")

    fmt.Println("Challenge 2")
    fmt.Println("-----------")

    key := []byte("YELLOW SUBMARINE")
    data := get_data("data/10.txt")
    crypt := ca.PadPkcs7([]byte(data), 16)

    iv := make([]byte, 16)
    plain, err := ca.DecryptCbc(crypt, key, iv)

    if err != nil {
        panic(err)
    }

    fmt.Println(string(plain))
}
