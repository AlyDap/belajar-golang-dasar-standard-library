package main

import (
	"fmt"
	"strconv"
)

// Sebelumnya kita sudah belajar cara konversi tipe data, misal dari int32 ke int34
// Bagaimana jika kita butuh melakukan konversi yang tipe datanya berbeda? Misal dari int ke string, atau sebaliknya
// Hal tersebut bisa kita lakukan dengan bantuan package strconv (string conversion)
// https://golang.org/pkg/strconv/


func main() {
	
	result, err := strconv.ParseBool("true") // ke boolean
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(result)
	}

	resultInt, err := strconv.Atoi("1000") // ke integer
	if err != nil {
		fmt.Println("Error", err.Error()) 
	} else {
		fmt.Println(resultInt)
	}

	binary := strconv.FormatInt(999, 2) // ke binary
	fmt.Println(binary)

	var stringInt string = strconv.Itoa(999) // ke string
	fmt.Println(stringInt)
}
