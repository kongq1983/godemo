package main

import "fmt"

func main() {
	continuedemo1()
	//continuedemo2()
	continuedemo3()
}

func continuedemo1() {

	for i := 0; i < 10; i++ {
		fmt.Printf("start current i=%d \n", i)
		if i > 5 {
			continue
		}
		fmt.Printf("end current i=%d \n", i)
	}

}

func continuedemo2() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j > 3 {
				break
			}
			fmt.Printf("current i=%d,j=%d \n", i, j)
		}
	}
}

func continuedemo3() {
	fmt.Println()
OUTER:
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("start current i=%d,j=%d \n", i, j)
			if j > 2 {
				continue OUTER
			}
			fmt.Printf("end current i=%d,j=%d \n", i, j)
		}
	}
}
