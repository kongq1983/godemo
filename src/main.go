package main

import "fmt"
import "mycom"
import . "mycom/entity"
import "reflect"

func main()  {

	empty1 := mycom.IsBlank("hello") // tool.go
	empty2 := mycom.IsEmpty("") //util.go

	//mycom.isEmpty("1");  //

	fmt.Printf("empty1=%t \n",empty1)
	fmt.Printf("empty2=%t \n",empty2)

	//不需要，可以用_代替
	//name,age,_ := mycom_entity.Empoyee();
	//name,age,_ := mycom_entity.Empoyee();
	//name,age,_ := mycom_entity.Empoyee();
	name,age,_ := Empoyee();

	fmt.Printf("name=%s,age=%d \n",name,age)

	//var a int16 = 10;

	//var b int32 = 10;
	//var b int16 = 10;

	//if(a==b) {
	//
	//}

	//var c int8;

	//if(c==128) {
	//
	//}

	//var a1 int32 = 10;
	//var a2 int64 = 20;
	//
	//var a3 int64 ;
	//// 类型不一致
	//a3 =  a1+a2;
	//
	//fmt.Printf("a3=%d",a3);

	aa := 1800;

	fmt.Printf("aa type is =%s \n",reflect.TypeOf(aa))

	strdemo();
	readStr();
	modifyStr();

}

func strdemo()  {
 	//str := "aA 你好 8"
 	str := "Go 你好!"
 	fmt.Printf("str.length=%d \n",len(str))

 	for i:=0;i<len(str);i++ {
		//fmt.Printf("str[%d]=%v u=%c,\n",i,str[i],str[i])
		fmt.Printf("str[%d]=%v,unicode=%U  \n",i,str[i],str[i])
	}

	t:= str;

	str += "欢迎您！"
    // 原字符串没有变
	fmt.Printf("string t = %s \n",t)
	// 字符串可以连接，但原字符串不会变
	fmt.Printf("string str = %s \n",str)

	//str[0]=71    G
	//str[1]=111   o
	//str[2]=32    空格
	//str[3]=228
	//str[4]=189
	//str[5]=160
	//str[6]=229
	//str[7]=165
	//str[8]=189
	//str[9]=33    !
}

func readStr()  {
	fmt.Println()
	str := "abcd你"  // 汉字占3个字节  一共7个字节

	for i :=0;i<len(str);i++ {
		fmt.Printf("str[%d]=%v,unicode=%U  \n",i,str[i],str[i])
	}

	// 编译不会报错，但是执行时会报错 汉字占3个字节
	//c := str[len(str)]
	//fmt.Println(len(str),c)

	// 你
	fmt.Println(str[4:])
	//abcd
	fmt.Println(str[:4])

	//join

	str1 := "你好,"+"King"
	fmt.Printf("str1=%s \n",str1)

	for i :=0;i<len(str1);i++ {
		fmt.Printf("%c",str1[i]);
	}
	fmt.Println()
	for _, v:= range str1 {
		fmt.Printf("%c",v);
	}
	fmt.Println()
	for i, v:= range str1 {
		fmt.Printf("index=%d,%c \n",i,v);
	}

}

func modifyStr() {
	fmt.Println()
	str := "Golang  欢迎你！";
	// str转化为[]byte数组，自动复制数据
	bs :=[]byte(str);
	bs[6] = ','

	fmt.Printf("%s\n",str)
	// 修改后的数据
	fmt.Printf("%s\n",bs)

	str1 := "Golang  欢迎你！";
	// 转化为[]rune，自动复制数据
	rs := []rune(str1);
	rs[8] = '恭'
	rs[9] = '喜'
	fmt.Println(str1)
	// 转化为字符串，又一次复制数据
	fmt.Println(string(rs))
}