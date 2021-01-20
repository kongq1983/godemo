package main

import (
	"fmt"
)

type (
	mystring string
)

func main() {

	sum := 100
	count := 6

	// int
	average1 := sum / count
	// float32
	average2 := float32(sum) / float32(count)

	fmt.Printf("average1 type =%T, value=%d \n", average1, average1)
	fmt.Printf("average2 type =%T, value=%f \n", average2, average2)

	var str mystring
	str = "类别别名demo"

	str2 := str + "!!!"

	fmt.Printf("str type =%T \n", str)
	fmt.Printf("str2 type =%T \n", str2)

	var str3 string
	str3 = "GoLang"
	fmt.Printf("str3 type =%T \n", str3)

}
