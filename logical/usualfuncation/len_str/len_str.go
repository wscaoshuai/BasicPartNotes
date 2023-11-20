package len_str

import (
	_ "fmt"
	"strconv"
	_ "strings"
)

// func Len_str1() {
// 	str := "hello北京"                 //golang的编码是按照utf-8的，汉字占用3个字节
// 	fmt.Println("str len", len(str)) //此处是按照字节返回
// }

// func Len_str2() {
// 	str := "hello上海"
// 	r := []rune(str)
// 	for i := 0; i < len(str); i++ {
// 		fmt.Printf("字符=%c\n", r[i])
// 	}
// }

// func Len_str3() {
// 	n, err := strconv.Atoi("12345")
// 	if err != nil {
// 		fmt.Println("转换错误", err)
// 	} else {
// 		fmt.Println("转换结果是：", n)
// 	}
// }

// func Len_str4() {
// 	str := strconv.Itoa(12345)
// 	fmt.Printf("str=%v,str=%T", str, str)
// }

// func Len_str5() {
// 	var bytes = []byte("hello go")
// 	fmt.Printf("bytes=%v\n", bytes)
// }

// func Len_str6() {
// 	str := string([]byte{97, 98, 99})
// 	fmt.Printf("str=%v\n", str)
// }

// func Len_str7() {
// 	str := strconv.FormatInt(123, 2)
// 	fmt.Printf("123对应的二进制是=%v\n", str)
// 	str1 := strconv.FormatInt(123, 16)
// 	fmt.Printf("123对应的16进制是=%v\n", str1)
// }

// func Len_str8() {
// 	b := strings.Contains("seafood", "mary")
// 	fmt.Printf("b=%v\n", b)
// }

// func Len_str9() {
// 	num := strings.Count("ceheese", "e")
// 	fmt.Printf("num=%v\n", num)
// }

// func Len_str10() {
// 	b := strings.EqualFold("abc", "Abc")
// 	fmt.Printf("b=%v\n", b)           //true
// 	fmt.Println("结果", "abc" == "ABC") //false  区分字母大小写
// }

// func Len_str11() {
// 	index := strings.Index("NLT_abcabcabc", "abc") //结果为：4
// 	fmt.Printf("index=%v\n", index)
// }

// func Len_str12() {
// 	index := strings.LastIndex("go golang", "go")
// 	fmt.Printf("index=%v\n", index)
// }

// func Len_str13() {
// 	str2 := "go go hello"
// 	str3 := strings.Replace(str2, "go", "北京", -1)
// 	fmt.Printf("str=%v str=%v\n", str2, str3)
// }

// func Len_str14() {
// 	strArr := strings.Split("hello,world,ok", ",")
// 	for i := 0; i < len(strArr); i++ {
// 		fmt.Printf("strArr=%v\n", strArr)
// 	}
// 	fmt.Printf("strArr=%v\n", strArr)
// }

// func Len_str15() {
// 	var str string
// 	str = "golang Hello"
// 	str = strings.ToLower(str)
// 	str = strings.ToUpper(str)
// 	fmt.Printf("str%v\n", str)
// }

// func Len_str16() {
// 	str := strings.TrimSpace("tn a lone gopher ntrn   ")
// 	fmt.Printf("str=%q\n", str)
// }

// func Len_str17() {
// 	str := strings.Trim("!hello!", "!")
// 	fmt.Printf("str=%q\n", str)
// }

// func Len_str20() {
// 	b := strings.HasPrefix("ftp://192.168.10.1", "hsp")
// 	fmt.Printf("b=%v\n", b)
// }

func Len_str21() { //编写一段代码来统计某个函数的执行时间
	str := ""
	for j := 0; j < 100000; j++ {
		str += "hello" + strconv.Itoa(j)
	}
}
