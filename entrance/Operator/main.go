package main

import (
	"fmt"
)

// func test() bool { //声明一个函数
// 	fmt.Println("test...")
// 	return true
//}
func main() {
	//golang中会根据参与运算的数字的字符类型决定结果的数值类型
	// var a int = 10 / 4
	// fmt.Println("a=", a)
	// var b float32 = 10.0 / 4
	// fmt.Println("b=", b)

	//取余公式：a % b = a - a / b *b
	// fmt.Println("10 % 3=", 10%3)
	// fmt.Println("-10 % 3=", -10%3)
	// fmt.Println("10 % -3=", 10%-3)
	// fmt.Println("-10 % -3=", -10%-3)

	//golang中自增自减只能作为单独的渔具使用，
	//错误写法：b:=a++;  b:=a--
	//并且没有后++a;++b;--a,--b

	//练习
	// var c = 97 //假期时间
	// var d = 7  //一周几天
	// var e int = (c / d)
	// var f int = (c - (d * e))
	// fmt.Println(e, "周", f, "天")

	// var g float32 = 30
	// var h float32 = (g - 100) * 5.0 / 9
	// fmt.Println("华氏温度：", g, "°", "摄氏温度：", h, "°")

	//逻辑运算：与，或，非：&&，||，！
	// var i int = 10
	// if i > 9 && test() { //如果该出i<9,则不判断test函数
	// 	fmt.Println("ok")
	// }

	// var j int = 10
	// var k int = 20
	// j, k = k, j
	// fmt.Println("j=", j, "k=", k)

	//其他运算符：&取地址符；*指针
	// var l int = 100
	// var m *int = &l
	// fmt.Println("l的地址：", &l, "m的值：", *m)

	//练习
	var n int = 99
	var o int = 98
	var p int = 97
	if n > o {
		fmt.Println("最大值：", n)
	} else {
		fmt.Println("最大值：", o)
	}

	if n > o {
		if n > p {
			fmt.Println("第一：", n)
		} else if o > p {
			fmt.Println("第二：", o)
		} else {
			fmt.Println("第三：", p)
		}
	}

}
