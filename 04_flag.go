package main

import (
	"flag"
	"fmt"
)

// Package flag berisikan fungsionalitas untuk memparsing command line argument
// https://pkg.go.dev/flag

func main() {
	// nama flag = host, default valuenya = localhost, deskripsi = put ...
	host := flag.String("host", "localhost", "Put yout database host")
	username := flag.String("username", "root", "Put yout database username")
	password := flag.String("password", "root", "Put yout database password")
	var port *int = flag.Int("port", 0, "database port")

	flag.Parse()

	// balikannya pointer, jadi pake asteris operator
	fmt.Println("host", *host)
	fmt.Println("username", *username)
	fmt.Println("password", *password)
	fmt.Println("port", *port)

	// go run .\04_flag.go -username=ali -password="rahasia bgt" -host=123.231.23.3 -port=5505

	// last := "Ali"
	// first := "Muhammad"
	// fmt.Printf("Huu %s%s\n", first, last)
}
