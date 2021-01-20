package main

import "fmt"
import "strconv"

func main() {

	strnum := "123"
	fmt.Printf("strnum type=%T, 当前占%d位 \n", strnum, strconv.IntSize)

	// 把字符串变为十进制
	num, err := strconv.Atoi(strnum)
	fmt.Printf("num type=%T, 且数字是%d \n", num, num)
	fmt.Printf("%s \n", err)

	num += 45
	// Itoa把十进制变为字符串 string
	newNum := strconv.Itoa(num)
	fmt.Printf("newNum type=%T, value=%s \n", newNum, newNum)

	strnum1 := "0101"
	toint, err1 := strconv.ParseInt(strnum1, 2, 32)
	fmt.Printf("toint type=%T, 且数字是%d \n", toint, toint)
	fmt.Printf("%s \n", err1)

	strnum2 := "22"
	toint2, err2 := strconv.ParseInt(strnum2, 8, 32)
	fmt.Printf("toint2 type=%T, 且数字是%d \n", toint2, toint2)
	fmt.Printf("%s \n", err2)

	var num3 int64
	num3 = 7
	str2num := strconv.FormatInt(num3, 2)
	fmt.Printf("str2num=%s \n", str2num)

}
