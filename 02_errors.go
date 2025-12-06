package main

import (
	"errors"
	"fmt"
)

// Sebelumnya kita sudah membahas tentang interface error yang merupakan representasi dari error di Go-Lang, dan membuat error menggunakan function errors.New()
// Sebenarnya masih banyak yang bisa kita lakukan menggunakan package errors, contohnya ketika kita ingin membuat beberapa value error yang berbeda
// https://pkg.go.dev/errors

var (
	ValidationError = errors.New("validation errror")
	NotFoundError   = errors.New("Not Found errror")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}
	if id != "ali" {
		return NotFoundError
	}

	return nil
}

func main() {
	err := GetById("asas")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("validation eror")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("not found eror")
		} else {
			fmt.Println("unknown eror")
		}
	}

	// last := "Ali"
	// first := "Muhammad"
	// fmt.Printf("Huu %s%s\n", first, last)
}
