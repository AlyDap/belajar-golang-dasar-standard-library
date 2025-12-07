package main

import (
	"container/list"
	"fmt"
)

// Package container/list adalah implementasi struktur data double linked list di Go-Lang
// https://golang.org/pkg/container/list/

func main() {
	// var data *list.List = list.New()
	data := list.New()
	data.PushBack("Muhammad")
	data.PushBack("Ferdynan")
	data.PushBack("ALi")

	var head *list.Element = data.Front()
	fmt.Println(head.Value) // Muhammad

	next := head.Next() // Ferdynan
	fmt.Println(next.Value)

	next = next.Next() // ALi
	fmt.Println(next.Value)

	fmt.Println("--------------")

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}
