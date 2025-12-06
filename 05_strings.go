package main

import (
	"fmt"
	"strings"
)

// Package strings adalah package yang berisikan function-function untuk memanipulasi tipe data String
// Ada banyak sekali function yang bisa kita gunakan
// https://golang.org/pkg/strings/


func main() {
	fmt.Println(strings.Contains("Ali Syahbana", "Ali"))
	fmt.Println(strings.Split("Ali Syahbana", " "))
	fmt.Println(strings.ToLower("Ali Syahbana"))
	fmt.Println(strings.ToUpper("Ali Syahbana"))
	fmt.Println(strings.Trim("      Ali Syahbana      ", " "))
	fmt.Println(strings.ReplaceAll("Ali Ferdynan Ali Ferdynan", "Ali", "Muhammad"))
}
