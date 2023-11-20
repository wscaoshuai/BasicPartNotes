package main

import (
	_ "errors"
	"fmt"
	_ "math/rand"
)

func main() {
	//编程流程：先易后难，先死后活
	//基本语法:if语句
	var a int
	fmt.Println("请输入年龄")
	fmt.Scanln(&a)
	if a > 18 {
		fmt.Println("请进")
	} else {
		fmt.Println("get out")
	}

	// //练习
	// var b int
	// var c int
	// fmt.Println("请输入两个数字")
	// fmt.Scanln(&b)
	// fmt.Scanln(&c)
	// if b+c >= 50 {
	// 	fmt.Println("hello world")
	// }

	//if练习1
	// var d float32
	// var e float32
	// fmt.Println("请输入两个数字")
	// fmt.Scanln(&d)
	// fmt.Scanln(&e)
	// if d > 10.0 && e < 20.0 {
	// 	fmt.Println(e + d)
	// }

	//if练习2
	// var f int32
	// var g int32
	// fmt.Println("请输入两个数字")
	// fmt.Scanln(&f)
	// fmt.Scanln(&g)
	// if (f+g)%3 == 0 && (f+g)%5 == 0 {
	// 	fmt.Println("两个数字和能同时被3和5整除")
	// }

	// // if嵌套
	// var h int32
	// fmt.Println("请输入成绩")
	// fmt.Scanln(&h)
	// if h == 100 {
	// 	fmt.Println("car")
	// } else if h <= 99 && h >= 80 {
	// 	fmt.Println("phone")
	// } else if h <= 79 && h > 60 {
	// 	fmt.Println("pad")
	// } else {
	// 	fmt.Println("nothing")
	// }

	//if练习3求ax²+bx+c=0得方程根，若b²-4ac>0,则有两个解，若b²-4ac=0，则有一个解，若b²-4ac<0，则无解
	//提示：x1=(-b+sqrt(b²-4ac))/2a;  x2=(-b-sqrt(b²-4acj)/2ak

	// var i float64
	// var j float64
	// var k float64
	// fmt.Println("请输入要测试的数值")
	// fmt.Scanln(&i)
	// fmt.Scanln(&j)
	// fmt.Scanln(&k)
	// l := j*j - 4*i*k
	// if l > 0 {
	// 	m1 := (-j + math.Sqrt(l)) / 2 * i
	// 	m2 := (-j - math.Sqrt(l)) / 2 * i
	// 	fmt.Printf("m1=%v,m2=%v", m1, m2)
	// } else if l == 0 {
	// 	m1 := (-j + math.Sqrt(l)) / 2 * i
	// 	fmt.Printf("m1=%v", m1)
	// } else {
	// 	fmt.Println("无解")
	// }

	//if练习4
	// var n int64
	// fmt.Println("请输入年份")
	// fmt.Scanln(&n)
	// if n%400 == 0 || n/4 == 0 && n%100 != 0 {
	// 	fmt.Println(n, "是闰年")
	// } else {
	// 	fmt.Println("平年")
	// }

	// //嵌套
	// var o int32
	// var p string
	// fmt.Println("请输入成绩")
	// fmt.Scanln(&o)
	// fmt.Println("请输入性别")
	// fmt.Scanln(&p)
	// if o <= 8 {
	// 	if p == "男" {
	// 		fmt.Println("恭喜进入男子总决赛")
	// 	} else if p == "女" {
	// 		fmt.Println("恭喜进入女子总决赛")
	// 	}
	// } else {
	// 	fmt.Println("您未进入总决赛")
	// }

	// //if练习5
	// var q byte
	// var r int8
	// var s float64
	// fmt.Println("请输入当前月份")
	// fmt.Scanln(&q)
	// fmt.Println("请输入用户年龄")
	// fmt.Scanln(&r)
	// fmt.Println("请输入票价")
	// fmt.Scanln(&s)
	// if q >= 4 && q <= 10 {
	// 	if r < 18 {
	// 		fmt.Println("当前为季节儿童票为：", 0.5*s)
	// 	} else if r >= 18 && r < 60 {
	// 		fmt.Println("当前季节成人票为：", s)
	// 	} else if r >= 60 {
	// 		fmt.Println("当前季节老人票为：", 0.3*s)
	// 	}
	// } else if q <= 12 && q > 10 || q >= 1 && q < 4 {
	// 	if r >= 18 && r < 60 {
	// 		fmt.Println("当前季节成人票为：", 0.7*s)
	// 	} else if r < 18 || r >= 60 {
	// 		fmt.Println("当前季节其他票价为：", 0.3*s)
	// 	}
	// } else {
	// 	fmt.Println("请输入正确的月份")
	// }

	//基本语法：
	//细节：1case、switch后可以使表达式（常量、变量、有返回值的函数等）；2case的表达式值类型必须和switch一致；3case后面可以使多个表达式；
	//4表达式如果是常量值则不能重复；5switch语句,默认有 break结束已执行语句；6defalut不是必须；7switch后面也可不跟表达式，类似if语句；
	//8switch后面可以接一个声明（定义）分号结束（不推荐使用）；9switch穿透—fallthrough,接在case语句后面则会执行下一个case。
	//10switch中可以使用 type-switch来判断interface变量中实际指向的变量类型

	//基本使用
	// var key byte
	// fmt.Println("请输入一个字符“a,b,c,d,e”")
	// fmt.Scanf("%c", &key)
	// switch key {
	// case 'a':
	// 	fmt.Println("星期一")
	// case 'b':
	// 	fmt.Println("星期二")
	// 	fallthrough //穿透到下一层的case（只可以穿透一层）
	// case 'c':
	// 	fmt.Println("星期三")
	// default:
	// 	fmt.Println("输入有误")

	// }

	// //type-switch使用
	// var t interface{}
	// var u = 10.0
	// t = u
	// switch v := t.(type) {
	// case nil:
	// 	fmt.Printf("t的类型：%T", v)
	// case int:
	// 	fmt.Printf("t的int类型")
	// case float64:
	// 	fmt.Printf("t是float64类型")
	// case func(int) float64:
	// 	fmt.Printf("t是func(int)类型")
	// case bool, string:
	// 	fmt.Printf("t是bool,string类型")
	// default:
	// 	fmt.Printf("未知类型")
	// }

	// //switch练习1
	// var x byte
	// fmt.Println("请输入一个小写英文字母：")
	// fmt.Scanf("%c", &x)
	// switch x {
	// case 'a':
	// 	fmt.Printf("A")
	// case 'b':
	// 	fmt.Printf("B")
	// case 'c':
	// 	fmt.Printf("C")
	// case 'd':
	// 	fmt.Printf("D")
	// case 'e':
	// 	fmt.Printf("E")
	// default:
	// 	fmt.Println("other")
	// }

	// //switch练习2//修改完成
	// var y int8 = 60
	// var min int8 = 0
	// var max int8 = 100
	// var score int8
	// fmt.Println("请输入成绩")
	// fmt.Scanln(&score)
	// switch {
	// case score >= y && score <= max:
	// 	fmt.Println("成绩合格")
	// case score < y && score >= min:
	// 	fmt.Println("成绩不合格")
	// default:
	// 	fmt.Println("请输入正确的成绩")
	// }

	// //switch练习3
	// var z int8
	// fmt.Println("输入当前月份")
	// fmt.Scanln(&z)
	// switch z {
	// case 3, 4, 5:
	// 	fmt.Println("Spring")
	// case 6, 7, 8:
	// 	fmt.Println("Summer")
	// case 9, 10, 11:
	// 	fmt.Println("Autumn")
	// case 12, 1, 2:
	// 	fmt.Println("Winter")
	// default:
	// 	fmt.Println("请输入正确的月份")
	// }

	//if和switch的比较：具体数字（较少）用switch，判断区间以及bool的用if更好

	// //for循环语句
	// //循环条件返回是一个bool值，
	// var A int8
	// for A = 0; A <= 10; A++ {
	// 	fmt.Println("hello")
	// }

	// //for循环的其他写法
	// var B int8 = 0 //初始化赋值
	// for B < 10 {   //循环判断
	// 	fmt.Println("~hello")
	// 	B++ //变量迭代
	// }

	// //for循环的其他写法(通常需要配合break使用)
	// var C int8 = 0
	// for {

	// 	if C < 10 {
	// 		fmt.Println("hello")
	// 	} else {
	// 		break
	// 	}
	// 	C++
	// }

	//遍历字符串、数组
	//传统方式遍历;按照字节遍历，utf-8编码，汉字占三个字节(英文和数字没问题)，将汉字所在的数组再转换成[]rune切片类型
	// var D string = "helloworld,北京"
	// var D1 = []rune(D) //不进行此步骤转换无法输出汉字
	// for E := 0; E < len(D1); E++ {
	// 	fmt.Printf("%c\n", D1[E])
	// }

	//for-range遍历;是按照字符的方式遍历的，相对方便，不用进行切片，即可输出汉字
	// var F string = "hello~world,上海"
	// for index, val := range F {
	// 	fmt.Printf("index=%d,val=%c\n", index, val)
	// }

	//for循环练习1
	// var G int64 = 100
	// var H int64 = 0
	// var sum int64 = 0
	// var count int64 = 0
	// for ; H <= G; H++ {
	// 	if H%9 == 0 {
	// 		fmt.Printf("%v\n", H)
	// 		count++
	// 		sum = sum + H
	// 	}
	// }
	// fmt.Printf("count=%v\n sum=%v", count, sum)

	//for循环练习2
	// var K int64 = 6
	// var L int64
	// for L = 0; L <= K; L++ {
	// 	fmt.Printf("%v + %v = %v\n", L, K-L, K)
	// }

	//go语言内没有while；do-while，但是可以用if语句实现同样的效果（如下）
	//while:
	// var M int64 = 1
	// for {
	// 	if M > 10 {
	// 		break //跳出for循环
	// 	}
	// 	fmt.Println("hello", M)
	// 	M++
	// }
	//do-while:
	// var N int64 = 1
	// for N <= 10 {
	// 	fmt.Println("hello~", N)
	// 	N++
	// 	if N > 10 {
	// 		break
	// 	}
	// }

	//多重循环控制练习1

	// var allscore int   //班级总成绩
	// var count int = 0  //及格人数计数器
	// var clanum int = 0 //班级数量
	// var stunum int = 0 //学生数量

	// fmt.Println("请输入班级数量")
	// fmt.Scanln(&clanum)
	// fmt.Println("请输入学生数量")
	// fmt.Scanln(&stunum)
	// for O := 1; O <= clanum; O++ {
	// 	var sum int = 0 //学生成绩计数器
	// 	for P := 1; P <= stunum; P++ {
	// 		var score int //学生成绩
	// 		fmt.Printf("请输入%d班 第%d个学生的成绩\n", O, P)
	// 		fmt.Scanln(&score)
	// 		sum += score
	// 		if score >= 60 {
	// 			count++
	// 		}
	// 	}
	// 	fmt.Printf("第%d班的平均成绩是%v\n", O, sum/stunum)

	// 	allscore += sum
	// }
	// fmt.Println("所有班的平均成绩为：", allscore/clanum)
	// fmt.Println("所有班的及格人数为：", count)

	//多重循环控制练习2(金字塔)
	//规律：星号（层数*2—1） 空格（总层数——当前层数）
	// var I int64
	// var J int64
	// var K int64
	// var L int64 = 10
	// for I = 1; I <= L; I++ { //打印层数
	// 	for J = 1; J <= L-I; J++ {
	// 		fmt.Print(" ")
	// 	}
	// 	for K = 1; K <= 2*I-1; K++ { //打印空格数
	// 		if K == 1 || K == 2*I-1 || I == L { //将中间判断为空的条件
	// 			fmt.Print("*")
	// 		} else {
	// 			fmt.Print(" ")
	// 		}
	// 	}
	// 	fmt.Println()
	// }

	//多重循环控制练习3(九九乘法表)
	// var A int64   //行
	// var B int64   //列
	// var sum int64 //和
	// for A = 1; A <= 9; A++ {
	// 	for B = 1; B <= A; B++ {
	// 		sum = A * B
	// 		fmt.Printf("%v*%v=%v \t", A, B, sum)
	// 	}
	// 	fmt.Printf("\n")
	// }

	//流程控制作业1（视频99）
	//判断一个数的范围
	// var num1 float64
	// fmt.Println("请输入要判断的数字")
	// fmt.Scanln(&num1)
	// if num1 > 0 {
	// 	if num1 != 0 {
	// 		fmt.Println("这个数字大于0")
	// 	} else {
	// 		fmt.Println("这个数字等于0")
	// 	}
	// } else {
	// 	fmt.Println("这个数字小于0")
	// }

	//判断一个年份是否是闰年
	// var year int64
	// fmt.Println("请输入一个年份")
	// fmt.Scanln(&year)
	// if year%100 == 0 || year%4 == 0 && year%400 != 0 {
	// 	fmt.Println("该年份是闰年")
	// } else {
	// 	fmt.Println("该年是平年")
	// }

	//流程控制作业2（视频99）
	//水仙花数
	// var num2 int64
	// var a int64 = 0
	// var b int64 = 0
	// var c int64 = 0
	// fmt.Println("请输入一个数字")
	// fmt.Scanln(&num2)
	// a = num2 / 100              //百位
	// b = (num2 - (100 * a)) / 10 //十位
	// c = num2 - 100*a - 10*b     //个位
	// if num2 == (a*a*a)+(b*b*b)+(c*c*c) {
	// 	fmt.Println("该数字是水仙花数")
	// } else {
	// 	fmt.Println("该数字不是水仙花数")
	// }
	// fmt.Printf("%v %v %v", a, b, c)

	//流程控制作业3（视频99）
	//登录判断(待修改，使用数组完成)
	// var username string
	// var password string
	// fmt.Println("请输入用户名")
	// fmt.Scanln(&username)
	// fmt.Println("请输入密码")
	// fmt.Scanln(&password)
	// if username == "admin" {
	// 	if password == "123456" {
	// 		fmt.Println("恭喜登录成功")
	// 	} else {
	// 		fmt.Println("请输入正确的密码")
	// 	}
	// } else {
	// 	fmt.Println("未找到该注册用户")
	// }

	//流程控制作业4（视频99）
	//输入年月判断该月份的天数（switch）
	// var year int64
	// var month int64
	// fmt.Println("请输入年月")
	// fmt.Scanln(&year)
	// fmt.Scanln(&month)
	// if year%100 != 0 && year%4 == 0 || year%400 == 0 {
	// 	switch month {
	// 	case 1, 3, 5, 7, 8, 10, 12:
	// 		fmt.Printf("%d年%d月有31天", year, month)
	// 	case 4, 6, 9, 11:
	// 		fmt.Printf("%d年%d月有30天", year, month)
	// 	case 2:
	// 		fmt.Printf("%d年%d月有28天", year, month)
	// 	}
	// } else {
	// 	switch month {
	// 	case 1, 3, 5, 7, 8, 10, 12:
	// 		fmt.Printf("%d年%d月有31天", year, month)
	// 	case 4, 6, 9, 11:
	// 		fmt.Printf("%d年%d月有30天", year, month)
	// 	case 2:
	// 		fmt.Printf("%d年%d月有29天", year, month)
	// 	}

	// }

	//流程控制作业5（视频99）
	//体重估算器【（身高-108）*2=体重】
	// var weight float64 = 0.0
	// var height float64 = 0.0
	// var w1 float64 = 0.0
	// var w2 float64 = 0.0
	// fmt.Println("请输入您的身高")
	// fmt.Scanln(&height)
	// weight = (height - 108) * 2
	// if weight > 0 {
	// 	w1 = weight + 10
	// 	w2 = weight - 10
	// 	fmt.Printf("正常体重范围在%v~%v之间，您的体重估计为：%v", w1, w2, weight)
	// } else {
	// 	fmt.Println("请输入正确的体重")
	// }

	//流程控制作业6（视频99）
	//考试等级判断(if-else解决)
	// var score int64
	// fmt.Println("请输入你的成绩")
	// fmt.Scanln(&score)
	// if score > 0 && score < 100 {
	// 	if score >= 60 {
	// 		if score <= 69 {
	// 			fmt.Println("成绩为及格")
	// 		} else if score <= 79 && score >= 70 {
	// 			fmt.Println("成绩为良好")
	// 		} else if score <= 89 && score >= 80 {
	// 			fmt.Println("成绩为优良")
	// 		} else {
	// 			fmt.Println("成绩为优秀")
	// 		}
	// 	} else {
	// 		fmt.Println("成绩为不合格")
	// 	}
	// }

	//流程控制作业6（视频99）
	//考试等级判断(switch解决)
	// var score2 int64
	// fmt.Println("请输入你的成绩")
	// fmt.Scanln(&score2)
	// switch score2 / 10 {
	// case 1, 2, 3, 4, 5:
	// 	fmt.Println("该生成绩不合格")
	// case 6:
	// 	fmt.Println("该生成绩合格")
	// case 7, 8:
	// 	fmt.Println("该生成绩良好")
	// case 9, 10:
	// 	fmt.Println("该生成绩优秀")
	// default:
	// 	fmt.Println("请输入正确的成绩")
	// }

	// var num3 int64
	// var num4 int64
	// fmt.Println("请输入两个成绩")
	// fmt.Scanln(&num3)
	// fmt.Scanln(&num4)
	// if num3%4 == 0 || num3+num4 > 1000 {
	// 	fmt.Println(num3)
	// } else {
	// 	fmt.Println(num4)
	// }

	//跳转控制语句-break：用于中断某个语句块，中断当前for循环（默认跳出最近的for循环），跳出switch语句。
	//如何随机生成（0-n)之间的一个伪随机数，使用（math)包下的rand函数。
	//如果生成一个随机数，则需要给rand设置一个种子：time.Now().Unix()即返回一个Unix时间（从1790年1月1日0时0分0秒到现在的秒数）

	//rand.Seed(time.Now().Unix())//设置种子，返回真正的随机数（Unix：秒 UnixNano:纳秒）
	// n := rand.Intn(100) + 1//生成范围内的伪随机数
	// fmt.Println(n)

	/*语法：
		label：{...
	 		label2{...
				label3{...
					break label2  //此处的“label2”标签(标签名字任意)表明跳出label2的循环，而不是跳出label3
	 			}

	 		}

		 }
	*/

	//跳转控制语句-break案例分析：
	// var count int64 = 0
	// for I := 1; I <= 100; I++ {
	// 	rand.Seed(time.Now().UnixNano())
	// 	n := rand.Intn(100) + 1
	// 	fmt.Println(n)
	// 	count++
	// 	if n == 99 {
	// 		break
	// 	}
	// }
	// fmt.Println("生成随机的99共用了所少次：", count)

	//跳转控制语句-break练习1
	//100之内的数求和，输出第一次加和大于20的两个随机数
	// label4:
	// for I := 1; I <= 100; I++ {
	// 	rand.Seed(time.Now().UnixNano())
	// 	I := rand.Intn(100) + 1
	// 	for J := 1; J <= 100; J++ {
	// 		rand.Seed(time.Now().UnixNano())
	// 		J := rand.Intn(100) + 1
	// 		if I+J >= 20 {
	// 			fmt.Println(I, J)
	// 			break label4
	// 		}
	// 	}
	// }

	//跳转控制语句-break练习2
	//实现登录验证，三次机会，
	// var username1 string
	// var password1 string
	// var loginnum int64 = 3
	// for K := 1; K <= 3; K++ {
	// 	fmt.Println("请输入用户名")
	// 	fmt.Scanln(&username1)
	// 	fmt.Println("请输入密码")
	// 	fmt.Scanln(&password1)
	// 	if username1 == "admin" && password1 == "123456" {
	// 		fmt.Println("登录成功")
	// 		break
	// 	} else {
	// 		loginnum--
	// 		fmt.Printf("还有%d次机会\n", loginnum)
	// 	}
	// }
	// if loginnum == 0 {
	// 	fmt.Println("今日登录次数已用完")
	// }

	//跳转控制语句-continue
	//用于结束本次循环，继续执行下一次循环
	//同break使用规则相同

	//跳转控制语句-continue案例分析
	// for i := 0; i < 4; i++ {
	// 	for j := 0; j < 10; j++ {
	// 		if j == 2 {
	// 			continue
	// 		}
	// 		fmt.Println("j=", j)
	// 	}
	// }

	//跳转控制语句-continue练习1(视频106)
	// var num5 int64
	// for num5 = 0; num5 <= 100; num5++ {
	// 	if num5%2 == 0 {
	// 		fmt.Println("偶数", num5)
	// 	} else {
	// 		continue
	// 	}
	// }

	//跳转控制语句-continue练习2(视频106)
	// var positive int64
	// var negative int64
	// var num6 int64
	// for i := 0; i <= 100; i++ {
	// 	fmt.Println("请输入一个数字")
	// 	fmt.Scanln(&num6)
	// 	if num6 == 0 {
	// 		break
	// 	}
	// 	if num6 > 0 {
	// 		positive++
	// 		continue
	// 	}
	// 	negative++
	// }
	// fmt.Printf("正数有：%d个,负数有：%d个", positive, negative)

	//跳转控制语句-continue练习3(视频106)(待重做)
	// var cash float64 = 100000
	// var count int64 = 0
	// for cash > 50000 {
	// 	cash = cash * ((1 - 0.05) * (1 - 0.05))
	// 	count++
	// 	if cash <= 50000 {
	// 		cash = cash - 1000
	// 		count++
	// 	} else if cash < 1000 {
	// 		break
	// 	}
	// }

	// fmt.Println("该人可以经过的路口：", count)

	//跳转控制语句-goto
	//通常配合条件语句使用，但一般不主张使用-goto语句,容易造成程序混乱
	// 	fmt.Println("1")
	// 	goto label5
	// 	fmt.Println("2") //2、3、4不执行
	// 	fmt.Println("3")
	// 	fmt.Println("4")
	// label5:
	// 	fmt.Println("5")

	//return 一般使用在函数或者方法中，跳出程序。

}
