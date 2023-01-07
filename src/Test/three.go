package main

import "fmt"

func main() {
	var arr = []int{1,2,3}

	printSlice(arr)
}

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}