package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Package bufio atau singkatan dari buffered io
// Package ini digunakan untuk membuat data IO seperti Reader dan Writer
// https://pkg.go.dev/bufio

func main() {
	input := strings.NewReader("this is long string\nwith new line\n")
	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine() // baca 1 baru pakai readline sampai enter
		if err == io.EOF {
			break
		}

		fmt.Println(string(line))
	}
}
