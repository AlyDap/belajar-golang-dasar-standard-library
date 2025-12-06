package main

import "fmt"

// Sebelumnya kita sudah sering menggunakan package fmt dengan menggunakan function Println
// Selain Println, masih banyak function yang terdapat di package fmt, contohnya banyak digunakan untuk melakukan format
// https://pkg.go.dev/fmt

func main() {
	last := "Ali"
	first := "Muhammad"
	fmt.Printf("Huu %s%s\n", first, last)
}
