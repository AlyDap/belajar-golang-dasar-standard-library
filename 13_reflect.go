package main

import (
	"fmt"
	"reflect"
)
// Dalam bahasa pemrograman, biasanya ada fitur Reflection, dimana kita bisa melihat struktur kode kita pada saat aplikasi sedang berjalan
// https://golang.org/pkg/reflect/

type Sample struct {
	Name string `required:"true" max:"10"`
}

type Person struct {
	Name    string `required:"true" max:"10"`
	Address string `required:"true" max:"30"`
	Email   string `required:"false" max:"20"`
	Age     int    `required:"false" max:"2"`
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)
	// valueType := reflect.TypeOf(value)
	fmt.Println("Type Name", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		var structField reflect.StructField = valueType.Field(i)
		//  structField := valueType.Field(i)
		fmt.Println(structField.Name, "with type", structField.Type)
		fmt.Println(structField.Tag.Get("required"))
		fmt.Println(structField.Tag.Get("max"))
	}
}

func IsValid(value any) (result bool) {
	result = true
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			result = data != ""
			if result == false {
				return result
			}
		}
	}
	return result
}

func main() {
	readField(Sample{"Ali"})
	fmt.Println("----------------------------")
	readField(Person{"Ali", "", "", 4})
	fmt.Println("----------------------------")

	person1 := Person{
		Name:    "ada",
		Address: "ada",
		Email:   "ada",
		Age:     23,
	}
	person2 := Person{Name: "ada", Address: "ada"}
	person3 := Person{Name: "ada", Age: 1}
	fmt.Println("person1", IsValid(person1))
	fmt.Println("person2",IsValid(person2))
	fmt.Println("person3",IsValid(person3))
}
