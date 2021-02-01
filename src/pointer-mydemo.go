package main

import "fmt"

func main() {
	var a int = 100
	var str string = "welcome to golang!"

	fmt.Printf("a.address=%p,str.address=%p \n", &a, &str)

	var str1 = "welcome to you !"
	ptr := &str1

	fmt.Printf("ptr.address=%p ptr.value=%s \n", ptr, *ptr)

	var ptr1 *string

	ptr1 = &str1

	fmt.Printf("ptr1.address=%p ptr1.value=%s \n", ptr1, *ptr1)

}
