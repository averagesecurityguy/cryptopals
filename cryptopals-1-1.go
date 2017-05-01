/*
Copyright (c) 2017, AverageSecurityGuy
# All rights reserved.

Complete challenge 1 from Cryptopals set 1.
http://cryptopals.com/sets/1

Usage:

$ go run cryptopals-1-1.go
*/

package main

import (
    "fmt"
)

import ca "github.com/averagesecurityguy/cryptanalysis"


func main() {
    fmt.Println("Cryptopals Set 1")
    fmt.Println("================")

    fmt.Println("Challenge 1")
    fmt.Println("-----------")
    hex_str := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
    b64 := ca.EncodeB64Str(ca.DecodeHexStr(hex_str))
    fmt.Println(b64)
}
