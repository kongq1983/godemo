package main

import "fmt"

func main() {
	demo1()
	demo2()
	demo3()
	demo4()
}

func demo1() {
	defer fmt.Printf("defer execute. \n")
	fmt.Printf("welcome to golang ! \n")
}

var i = 0
var j = 0

func printI() {
	fmt.Printf("i=%d \n", i)
}
func demo2() {

	for ; i < 10; i++ {
		defer printI()
	}

}

func printJ(j int) {
	fmt.Printf("j=%d \n", j)
}
func demo3() {

	for ; j < 10; j++ {
		defer printJ(j)
	}

}

func demo4() {
	defer fmt.Printf("defer execute 1. \n")
	defer fmt.Printf("defer execute 2. \n")
	defer fmt.Printf("defer execute 3. \n")
	defer fmt.Printf("defer execute 4. \n")
}
