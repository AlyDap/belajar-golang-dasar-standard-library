package main

import (
	"fmt"
	"regexp"
)

// Package regexp adalah utilitas di Go-Lang untuk melakukan pencarian regular expression
// Regular expression di Go-Lang menggunakan library C yang dibuat Google bernama RE2
// https://github.com/google/re2/wiki/Syntax
// https://golang.org/pkg/regexp/

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`a([a-z])i`)

	fmt.Println(regex.MatchString("ali"))
	fmt.Println(regex.MatchString("api"))
	fmt.Println(regex.MatchString("aLi"))

	fmt.Println(regex.FindAllString("ali api apa awi a2i asi aLi", 10))
}
