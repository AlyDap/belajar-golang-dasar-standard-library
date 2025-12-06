package main

import (
	"fmt"
	"os"
)

// Go-Lang telah menyediakan banyak sekali package bawaan, salah satunya adalah package os
// Package os berisikan fungsionalitas untuk mengakses fitur sistem operasi secara independen (bisa digunakan  disemua sistem operasi)
// https://golang.org/pkg/os/

func main() {
	// Mengambil data argumen ketika aplikasinya dijalankan, pakai args:
	// go run .\03_os.go Muhammad Ferdynan Ali Syahbana
	// go run .\03_os.go "Muhammad Ferdynan Ali Syahbana"

	args := os.Args
	for _, arg := range args {
		fmt.Println(arg)
	}

	// mengambil nama hostnama
	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("Error", err.Error())
	}
}
