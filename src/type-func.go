package main

import "fmt"

/**
  是否基数
*/
func isOdd(num int) bool {
	if num%2 == 1 {
		return true
	}

	return false
}

/**
  是否偶数
*/
func isEven(num int) bool {
	if num%2 == 0 {
		return true
	}

	return false
}

// 声明函数类型
type myBoolFunc func(int) bool

func filter(array []int, myFunc myBoolFunc) []int {
	var result []int

	for _, v := range array {
		if myFunc(v) {
			result = append(result, v)
		}
	}

	return result

}

func main() {
	array := []int{1, 2, 3, 4, 5, 6}
	oddArray := filter(array, isOdd)
	fmt.Println("oddArray=", oddArray)
	evenArray := filter(array, isEven)
	fmt.Println("evenArray=", evenArray)
	//isOddNum :=
}
