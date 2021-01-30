package main

import "fmt"

/**
匿名函数
*/
func main() {

	myadd := madd
	sum := myadd(3, 5)
	fmt.Printf("sum=%d \n", sum)

	multiply := func(x, y int) int { return x * y }
	multiplyvalue := multiply(3, 6)
	fmt.Printf("multiplyvalue=%d \n", multiplyvalue)

	sum1()

}

func madd(x int, y int) int {
	return x + y
}

func sum1() {

	total := func(num int) int {
		sum := 0
		for i := 1; i <= num; i++ {
			sum += i
		}
		fmt.Printf("sum=%d \n", sum)
		return sum
	}(100)

	fmt.Printf("total=%d \n", total)

}
