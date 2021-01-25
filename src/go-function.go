package main

import "fmt"

func main() {

	c := add(1, 8)
	fmt.Printf("c=%d \n", c)

	//sum := add;
	//
	//d:= add(c,9);
	//fmt.Printf("d=%d \n",d)

	a, b := moreOper(1, 9)
	fmt.Printf("a=%d,b=%d \n", a, b)

	sum(1, 2, 3, 4, 5, 6, 7, 8, 9)

}

func add(a, b int) int {
	return a + b
}

func moreOper(a, b int) (int, int) {
	return a + b, a * b
}

func sum(args ...int) {
	total := 0
	for _, v := range args {
		total += v
	}

	fmt.Printf("total=%d \n", total)
}
