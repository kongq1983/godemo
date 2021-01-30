package main

import "fmt"
import "strconv"

func main() {

	// 长度和容量都是5个元素
	array := make([]string, 5)

	for i := 0; i < len(array); i++ {
		array[i] = strconv.Itoa(i) // int to string
		fmt.Printf("%d=%s \n", i, array[i])
	}

	fmt.Println()
	//长度为3个  容量为5个
	array1 := make([]string, 3, 5)
	for i := 0; i < len(array1); i++ { //可以访问3个元素
		array1[i] = strconv.Itoa(i) // int to string
		fmt.Printf("%d=%s \n", i, array1[i])
	}

	myslice()

	myappend()

}

func myslice() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	slice1 := slice[2:5]
	fmt.Printf("\n")
	fmt.Printf("slice1.length=%d \n", len(slice1))
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("%d=%d \n", i, slice1[i])
	}

	slice2 := slice[2:3:7]
	fmt.Printf("\n")
	fmt.Printf("slice2.length=%d \n", len(slice2))
	for i := 0; i < len(slice2); i++ {
		fmt.Printf("%d=%d \n", i, slice2[i])
	}

}

func myappend() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("start myappend =============================\n")
	// 3,4,5
	slice1 := slice[2:5]
	fmt.Printf("\n")
	fmt.Printf("slice1.length=%d \n", len(slice1))
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("%d=%d \n", i, slice1[i])
	}

	//slice1 = append(slice1,10,11,12,13,14,15,16,17,18,19)
	slice1 = append(slice1, 10, 11, 12, 13, 14)
	fmt.Printf("\n")

	fmt.Printf("slice.length=%d \n", len(slice))
	for i := 0; i < len(slice); i++ {
		fmt.Printf("%d=%d \n", i, slice[i])
	}
	fmt.Printf("\n")
	fmt.Printf("slice1.length=%d \n", len(slice1))
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("%d=%d \n", i, slice1[i])
	}
}
