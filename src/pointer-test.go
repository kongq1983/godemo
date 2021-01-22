package main

import "fmt"

func main() {
	num := 10
	fmt.Printf("num.address=%x \n", &num)

	num1 := 18
	num1pointer := &num1

	fmt.Printf("num1.address=%x \n", &num1)
	fmt.Printf("num1pointer.address=%x \n", num1pointer)
	fmt.Printf("*num1pointer.val=%d \n", *num1pointer)

	var ptr *int
	fmt.Printf("ptr.address=%x \n", ptr)
	// 会报错
	//*ptr = 10

	//var p *int = nil
	// 会报错
	//*p = 10

	// 定义1个指针  执行nil  0
	// 1个是指针地址   1个是指针锁指向的地址
	var aa *int
	// 0
	aap := &aa

	fmt.Println()

	fmt.Printf("aa指针自己地址=%x \n", &aa)
	fmt.Printf("aa指针所指向的内存地址=%x \n", aa)
	fmt.Printf("aap.address=aa指针地址=%x \n", aap)
	// aap指针指向aa的内存地址
	fmt.Printf("aa指针所指向的内存地址=%x \n", *aap)
	fmt.Printf("&aap=aap的内存地址 %x \n", &aap)

}
