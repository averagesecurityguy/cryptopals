/*
Copyright (c) 2017, AverageSecurityGuy
# All rights reserved.

Complete challenge 2 from Cryptopals set 1.
http://cryptopals.com/sets/1

Usage:

$ go run cryptopals-1-2.go
*/

package main

import (
    "fmt"
)

import ca "github.com/averagesecurityguy/cryptanalysis"


func main() {
    fmt.Println()
    fmt.Println("Challenge 2")
    fmt.Println("-----------")
    b1 := ca.DecodeHexStr("1c0111001f010100061a024b53535009181c")
    b2 := ca.DecodeHexStr("686974207468652062756c6c277320657965")
    str := ca.EncodeHexStr(ca.XOR(b1, b2))

    fmt.Println(str)
}
