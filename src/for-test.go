package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf(" i=%d \n", i)
	}

	str := "123456789"

	fmt.Println()
	for i, char := range str {
		fmt.Printf(" i=%d，char=%d \n", i, char)
	}

	m := map[string]int{"one": 1, "two": 2}

	for k, v := range m {
		fmt.Printf(" key=%s,value=%d \n", k, v)
	}

	numbers := []int{10, 20, 30, 40}

	for i, v := range numbers {
		fmt.Printf(" 第%d次，value=%d \n", i, v)
	}

}
