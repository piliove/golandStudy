package main

import (
	"encoding/json"
	"fmt"
)

type ProductInfo struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

func main() {
	fmt.Println("111")
	return

	str := "{'name':'AppleWatchS8','price':'3199'}"
	data := ProductInfo{}
	if err := json.Unmarshal([]byte(str), &data); err != nil {
		fmt.Println("error: " + err.Error())
	} else {
		fmt.Println(data)
	}
}
