package main

import "fmt"

func main() {
	a := 88

	if a < 10 {
		fmt.Printf("a < 10 \n")
	} else {
		fmt.Printf("a >= 10 \n")
	}

	fmt.Printf("a的值=%d \n", a)

	if a >= 100 {
		fmt.Printf("a > 100 \n")
	} else if a >= 80 {
		fmt.Printf("a >= 80 and a < 100 \n")
	} else {
		fmt.Printf("a < 80 \n")
	}

	ifdemo1()

}

func ifdemo1() {
	// 子语句只能有1个表达式
	if a := 18; a < 20 {
		fmt.Printf(" a <20 \n")
	} else {
		fmt.Printf(" a >=20 \n")
	}

	// a编译不通过
	//fmt.Printf(" a.value=%d \n",a)

}
