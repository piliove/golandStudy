package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string `json:name`
	Age int `json:age`
}

func main() {
	// m := map[int]string{1: "a", 2: "b", 3: "c"}

	m := Person{
		Name: "HaiCoder",
		Age:109,
	}

	// len := len(m)
	// fmt.Println(len)

	fmt.Println(m)

	personValue := reflect.ValueOf(m)
	nameValue := personValue.FieldByName("Name").String()

	fmt.Println(nameValue)
}
