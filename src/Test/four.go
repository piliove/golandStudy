package main

import "fmt"

func Factorial(x int) (result int) {
	if x == 0 {
		result = 1;
	} else {
		result = x * Factorial(x - 1);
	}
	return;
}

func main() {
	// var i int = 15
	// fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(i))
	var a int8 = 12
	fmt.Printf("a 是%d\n", a)
}