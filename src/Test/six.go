package main

import "fmt"

func main() {
	//声明方式一
	var ints1 []int
	//声明方式二
	ints2 := make([]int, 0)

	// fmt.Println(ints1)
	// fmt.Println(ints2)

	ints1 = append(ints1,1,2,3,4,5)
	ints2 = append(ints2,1,2,3,4,5,6,7)

	fmt.Println(ints1)
	fmt.Println(ints2)
}