package main

import (
	"encoding/json"
	"fmt"
)

type ProductInfo struct {
	Name  string  `json:"name"`
	Price string `json:"price"`
}

func main() {

	str := "{\"name\":\"AppleWatchS8\",\"price\":\"3199\"}"
	data := ProductInfo{}

	err := json.Unmarshal([]byte(str), &data)
	fmt.Println(str)
	// fmt.Println(data)
	fmt.Println(err)

	name := ProductInfo{}.name
	fmt.Println(name)
	fmt.Printf("%+v", data)

	return

	if err := json.Unmarshal([]byte(str), &data); err != nil {
		fmt.Println("error: " + err.Error())
	} else {
		fmt.Println(data)
	}
}
