package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "Hello"
	str2 := string("World")

	// fmt.Println(str1+ " " +str2)

	/**
	  strings的常见函数
	*/
	index := strings.Index(str1, "e")        //查找“e”在str1中的位置
	contains := strings.Compare(str1, str2)  //Compare返回一个按字典顺序比较两个字符串的整数。 如果a==b，则结果为0，如果a < b则为-1，如果a > b则为+1。
	fold := strings.EqualFold(str1, str2)    //类似于Java里的equals（）
	contain := strings.Contains(str1, "llo") //查找“llo”是否在str1中

	index02 := strings.Index(str1, "o")
	fmt.Println(index02)

	fmt.Println(index)
	fmt.Println(contains)
	fmt.Println(fold)
	fmt.Println(contain)
}