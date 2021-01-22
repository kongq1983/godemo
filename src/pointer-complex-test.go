package main

import "fmt"

func main() {
	a := 10
	ap := &a
	app := &ap

	//10
	fmt.Printf("a: %d \n", a)
	// a的地址
	fmt.Printf("a: %x \n", &a)
	// c0000a2058  a的地址
	fmt.Printf("ap: %x \n", ap)
	// 10 ap所指向的对象 10
	fmt.Printf("*ap: %d \n", *ap)
	// c0000cc018 ap指针的地址
	fmt.Printf("app: %x \n", app)
	// c0000a2058 ap指针所指向对象=a的地址
	fmt.Printf("*app: %x \n", *app)
	// 10
	fmt.Printf("**app: %d \n", **app)

	pointerArray()
	swapDemo()

}

func pointerArray() {
	a := []int{100, 200, 300}
	var ptr [3]*int

	for i := 0; i < 3; i++ {
		ptr[i] = &a[i]
		fmt.Printf("a[%d]的地址：%x\n", i, ptr[i])
	}

	for i := 0; i < 3; i++ {
		fmt.Printf("a[%d]的值：%d\n", i, *ptr[i])
	}

}

func swapDemo() {
	a := 100
	b := 200

	fmt.Printf("交换之前a的值为：%d\n", a)
	fmt.Printf("交换之前b的值为：%d\n", b)

	swap(&a, &b)

	fmt.Printf("交换之后a的值为：%d\n", a)
	fmt.Printf("交换之后b的值为：%d\n", b)

}

func swap(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
