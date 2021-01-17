package main

import "fmt"
import "os"

const (
	OPEN = 1
	CLOSE = 0
)

const (
	ONE,TWO,THREE = 1,2,3
	FIRST,SECOND,THIRD = 1,2,3
)

var (
	HOME = os.Getenv("HOME")
	USER = os.Getenv("USER")
	GOROOT = os.Getenv("GOROOT")
	PATH = os.Getenv("PATH")
)

func main1()  {
	fmt.Println("welcome to golang!")


	n :=0;
	p := &n;
	n++;
	*p++;
	fmt.Println(n)
	fmt.Println(&n)
	fmt.Println(*p)
	fmt.Println(p)

	var name0 string = "golang0";
	//
	var name1 = "golang1";
	//常量
	const name2 string = "golang2";
	const name3 = "golang3";

	fmt.Printf("name is %s \n",name0)
	fmt.Printf("name is %s \n",name1)
	fmt.Printf("name is %s \n",name2)

	var num1 int = 18;
	var num2 = 20;
	num3 := 20;

	const pi float64 = 3.1415926

	fmt.Printf("num01=%d \n",num1)
	fmt.Printf("num1=%d \n",num1)
	fmt.Printf("num2=%d \n",num2)
	fmt.Printf("num3=%d \n",num3)
	fmt.Printf("pi=%f \n",pi)

	var limit bool = false;
	var over = true;
	fmt.Printf("limit=%t \n",limit)
	fmt.Printf("over=%t \n",over)

	var a1, a2, a3 = 1, 2, 3 ;
	fmt.Printf("a1=%d,a2=%d,a3=%d\n",a1,a2,a3)

	aa1, aa2, aa3 := 1, 2, 3 ;
	fmt.Printf("aa1=%d,aa2=%d,aa3=%d\n",aa1,aa2,aa3)

	string1,int1 := "golang",10;
	fmt.Printf("string1=%s,int1=%d\n",string1,int1)

	// 带0x的指针 0xc00000a0e0p
	fmt.Printf("p address = %p \n",p);
	// 不带0x的指针 c00000a0e0
	fmt.Printf("p address = %#p \n",p);

	fmt.Printf("open=%d,close=%d \n",OPEN,CLOSE)
	fmt.Printf("one=%d,two=%d \n",ONE,TWO)
	fmt.Printf("first=%d,second=%d \n",FIRST,SECOND)

	fmt.Printf("home=%s,user=%s,goroot=%s,path=%s \n",HOME,USER,GOROOT,PATH)

	filedemo()

}

func filedemo() {
	//file, err := os.Stat(`D:\log\a.txt`)
	//file, err := os.Stat(`helloworld.go`)
	// 目前只能写helloworld.go 和helloworld.go同目录的其他文件都不行 其他都会报The system cannot find the file specified，目前还不知原因
	file, err := os.Stat(`helloworld.go`)

	if err != nil {
		fmt.Println("err = ", err)
	}

	if (file != nil) {
		fmt.Println(file.Name()) //获取文件名
		fmt.Println(file.IsDir())   //判断是否是目录，返回bool类型
		fmt.Println(file.ModTime()) //获取文件修改时间
		fmt.Println(file.Mode())
		fmt.Println(file.Size()) //获取文件大小
		fmt.Println(file.Sys())
	}
}