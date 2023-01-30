package main

import "fmt"

func main() {
	//声明方式一
	var ints1 []int
	//声明方式二
	ints2 := make([]int, 0)

	fmt.Println(ints1)
	fmt.Println(ints2)
}