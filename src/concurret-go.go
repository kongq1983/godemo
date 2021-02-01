package main

import (
	"fmt"
	"time"
)

func add3(x, y int) {
	z := x + y
	fmt.Printf("z=%d \n", z)
}

func main() {

	for i := 0; i < 10; i++ {
		//add3(i,i)
		go add3(i, i) // 如果没有下面的time.Sleep  不会输出printf  因为main已经退出了  相当于java主线程已经退出了
	}

	// 等待1s
	time.Sleep(5 * time.Second)
}
