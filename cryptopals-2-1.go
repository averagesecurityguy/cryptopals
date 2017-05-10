/*
Copyright (c) 2017, AverageSecurityGuy
# All rights reserved.

Complete challenge 1 from Cryptopals set 2.
http://cryptopals.com/sets/2

Usage:

$ go run cryptopals-2-1.go
*/

package main

import (
    "fmt"
    ca "github.com/averagesecurityguy/cryptanalysis"
)


func main() {
    fmt.Println("Cryptopals Set 2")
    fmt.Println("================")

    fmt.Println("Challenge 1")
    fmt.Println("-----------")
    str := "YELLOW SUBMARINE"
    padded := ca.PadPkcs7([]byte(str), 20)
    fmt.Printf("%x\n", padded)
}
