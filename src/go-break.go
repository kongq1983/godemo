package main

import "fmt"

func main() {
	breakdemo1()
	breakdemo2()
	breakdemo3()
}

func breakdemo1() {

	for i := 0; i < 10; i++ {
		if i > 5 {
			break
		}
		fmt.Printf("current i=%d \n", i)
	}

}

func breakdemo2() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j > 3 {
				break
			}
			fmt.Printf("current i=%d,j=%d \n", i, j)
		}
	}
}

func breakdemo3() {
	fmt.Println()
OUTER:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j > 3 {
				break OUTER
			}
			fmt.Printf("current i=%d,j=%d \n", i, j)
		}
	}
}
