package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {

	str := "/usr/local/redis/bin/redis.sh"
	fmt.Printf("start with /usr result=%t \n", strings.HasPrefix(str, "/usr1"))
	fmt.Printf("start with /usr result=%t \n", strings.HasPrefix(str, "/usr"))

	fmt.Printf("end with .sh result=%t \n", strings.HasSuffix(str, ".sh"))

	fmt.Printf("contains with bin result=%t \n", strings.Contains(str, "bin"))

	fmt.Printf("ContainsAny with bbb result=%t \n", strings.ContainsAny(str, "fa"))

	// 返回0
	fmt.Printf("Index with usr result=%d \n", strings.Index(str, "/"))

	// 返回20
	fmt.Printf("Index with usr result=%d \n", strings.LastIndex(str, "/"))

	// 返回9
	fmt.Printf("Index with usr result=%d \n", strings.Index("你好abc中国", "中"))

	// 返回9
	fmt.Printf("Index with usr result=%d \n", strings.IndexRune("你好abc中国", '中'))

	str1 := "name one, name two, name three, name four"
	new := "Go"
	old := "name"
	result := strings.Replace(str1, old, new, 5)

	fmt.Printf("result=%s \n", result)

	str2 := strings.Repeat("abc", 3)
	fmt.Printf("str2=%s \n", str2)

	str3 := "one two three four one one two"
	// 3
	oneTimes := strings.Count(str3, "one")
	fmt.Printf("oneTimes = %d \n", oneTimes)

	twoTimes := strings.Count(str3, "two")
	fmt.Printf("twoTimes = %d \n", twoTimes)

	//str5 := "golang你好!"
	str5 := "你好!"
	// 3
	fmt.Printf("%s字符数-one=%d \n", str5, len([]rune(str5)))
	// 7 字节统计
	fmt.Printf("%s字节数-two=%d \n", str5, len(str5))
	// 3
	fmt.Printf("%s字符数-three=%d \n", str5, utf8.RuneCountInString(str5))

	str6 := "aBcDeF你好！"

	fmt.Printf("%s小写=%s \n", str6, strings.ToLower(str6))
	fmt.Printf("%s大写=%s \n", str6, strings.ToUpper(str6))

	// Trim 裁剪字符串开头或者结尾的字符，也就是不匹配中间的字符
	// GoLang
	fmt.Printf("%q\n", strings.Trim(" !!! GoLang !!! ", "! "))
	// GoLang
	fmt.Printf("%q\n", strings.Trim(" !!! GoLang !!! ", " ! "))
	//  !!! GoLang !!!
	fmt.Printf("%q\n", strings.Trim(" !!! GoLang !!! ", "!"))

	// GoLang !!!   TrimLeft 修改字符串开头的字符
	fmt.Printf("%q\n", strings.TrimLeft(" !!! GoLang !!! ", " ! "))
	//  !!! GoLang !!!
	fmt.Printf("%q\n", strings.TrimLeft(" !!! GoLang !!! ", "!"))

	//  !!! GoLang   TrimRight 修改字符串结尾的字符
	fmt.Printf("%q\n", strings.TrimRight(" !!! GoLang !!! ", " ! "))
	//  !!! GoLang !!!
	fmt.Printf("%q\n", strings.TrimRight(" !!! GoLang !!! ", "!"))

	// GoLang
	fmt.Printf("%q\n", strings.TrimSpace(" \t\n GoLang \t\n\r\n "))

	// 天是几号
	fmt.Printf("%q\n", strings.Trim("今天是几号", "今"))
	// 今天是几号
	fmt.Printf("%q\n", strings.Trim("今天是几号", "天"))
	// 气不错，我   今天会被拆分成2个字符分别匹配裁剪
	fmt.Printf("%q\n", strings.Trim("今天天气不错，我天", "今天"))

	// ["one" "two" "three" "four" "five"]
	fmt.Printf("%q\n", strings.Split("one two three four five", " "))

	str7 := "one two three four five 你好"
	array := strings.Fields(str7) // 按空格分割

	str8 := strings.Join(array, ";")
	fmt.Printf("str8=%s \n", str8)

	for i, val := range array {
		fmt.Printf("index=%d,val=%s \n", i, val)
	}

	num9 := 1000

	fmt.Printf("num9 二进制=%b \n", num9)
	fmt.Printf("num9 八进制=%o \n", num9)
	fmt.Printf("num9 十进制=%d \n", num9)
	fmt.Printf("num9 小写十六进制=%x \n", num9)
	fmt.Printf("num9 大写十六进制=%X \n", num9)

}
