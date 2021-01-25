package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j > 3 {
				goto GOBREAK
			}
			fmt.Printf("current i=%d,j=%d \n", i, j)
		}
	}

GOBREAK:
	fmt.Printf("goto end ---------------")

}
