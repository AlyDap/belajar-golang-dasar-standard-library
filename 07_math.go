package main

import (
	"fmt"
	"math"
)

// Package math merupakan package yang berisikan constant dan fungsi matematika
// https://golang.org/pkg/math/

func main() {
	fmt.Println(math.Ceil(1.40))  // bulat ke atas
	fmt.Println(math.Floor(1.60)) // bulat ke bawah
	fmt.Println(math.Round(1.60)) //  bulat terdekat
	fmt.Println(math.Max(10, 11)) // cari nilai maks
	fmt.Println(math.Min(10, 11)) // cari nilai min
}
