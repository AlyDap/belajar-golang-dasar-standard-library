package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Dir("hello/world.go"))             // hello
	fmt.Println(filepath.Base("hello/world.go"))            // world.go
	fmt.Println(filepath.Ext("hello/world.go"))             // .go
	fmt.Println(filepath.IsAbs("hello/world.go"))           // apakah absolute? false
	fmt.Println(filepath.IsLocal("hello/world.go"))         // apakah local? true
	fmt.Println(filepath.Join("hello", "world", "main.go")) // hello/world/main.go
}
