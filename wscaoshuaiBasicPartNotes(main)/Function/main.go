package main

import (
	"fmt"
	_ "progect1/closure"
	_ "progect1/defer1"
	_ "progect1/init" //(待修改)
	_ "progect1/multiplication"
	_ "progect1/pascal"
	_ "progect1/suffix"
	"progect1/usualfuncation/exeo"

	/*该出放置别名，使用后后面的地址失效*/ /*"go_code/progect1/mode03"*/
	_ "progect1/Fibonacci"
	_ "progect1/monkey"
	_ "progect1/swap"
	_ "progect1/usualfuncation/error"
	_ "time"
)

// 	var (
// 		Fun1 = func(n1 int, n2 int) int {//接119行，此时fun1就是一个全局匿名函数
// 			return n1 * n2
// 		}
//	 )

func main() {
	fmt.Println("Test,测试,Test,测试,Test,测试,Test,测试,Test,测试,Test,测试")
	//函数的基本介绍
	//基本语法
	//	func 函数名 （形参列表） （返回值列表）{
	//			执行语句
	//			return 返回值列表
	//	}
	//函数的使用细节
	//1、函数的（形参、返回值列表）可以是多个
	//2、形参和返回值列表的数据类型可以使值类型，也可以是引用类型
	//3、函数名大写则可以被引用，其余命名规则不变
	//4、函数中的变量是局部变量
	//5、基本数据类型和数组都是值的传递，在函数内修改，不影响原来的值
	//6、如果希望改变函数外的变量，则需要引入地址&，通过指针的方式操作变量
	//7、不支持函数重载
	//8、函数也是一种数据类型，可以赋值给变量，则该变量就成为一个函数类型的变量，可以通过该变量调用函数
	//9、函数也可以作为形参，并且被调用
	//10、支持自定义数据类型；基本语法； type 自定义数据类型名 数据类型 （相当于起一个别名）
	//11、支持对函数返回值命名
	//12、使用下划线忽略返回值
	//13、支持可变参数用法：（0-n）个参数：func sum(args... int)sum int{}
	//                     (1-n)个参数:func sum(n1 int,args... int)sum int{}
	//                     其中args就是slice切片，通过args[index]可以访问到各个值; 形参列表内的可变参数要放在最后

	//函数传参的方式：
	//1、值传递
	//2、引用传递
	//两者传递的都是变量的副本，前者是对值的拷贝，后者是对地址的拷贝，后者效率更高
	//值类型：基本数据类型int系列，float系列，bool,string、数组和结构体struck
	//引用类型：指针、slice切片、map、管道chan、interface等都是引用类型

	// var num1 float64
	// var num2 float64
	// var operator byte = '+'
	// fmt.Println("请输入两个数字以及一个运算符")
	// fmt.Scanln(&num1)
	// fmt.Scanln(&num2)
	// fmt.Scanln(&operator)
	// result := mode03.Cal(num1, num2, operator)
	// fmt.Println("result=", result)

	//包的基本介绍：go的每一个文件都属于一个包，go是以包的形式管理文件和目录结构的。
	//1、区分相同名字的函数、变量、标识符等
	//2、可以很好的管理大量的文件
	//3、控制函数、变量访问范围，即作用域
	//4、为了使其他包的文件可以访问，必须使改文件名的首字母需要大写
	//5、如果包名过长，可以取一个别名，原来的名字不能使用
	//6、同一个包下，不允许有相同的函数名，也不可以有相同的全局变量名

	//go语言函数的调用机制

	//return语句
	//基本语法
	// 	func 函数名（形参列表）（返回值列表）{
	// 		语句。。。。
	// 		return （返回值列表）
	// 	}
	//1、返回多个值时用“_”占位符可以忽略该值
	//2、如果只返回一个值，可以不写括号

	//递归调用
	//1、执行函数时，就会创建一个独立的空间
	//2、函数的局部变量是独立的，不会相互影响
	//3、递归必须向退出递归的条件逼近，否则就会无限递归
	//4、当函数执行完毕，就会return，遵循谁调用，返回给谁。

	//递归课堂练习1
	//斐波那契数列
	//此处调用progect1文件夹下Fibonacci包下Fibonacci函数
	// result := Fibonacci.Fibonacci(10)
	// fmt.Println("result=", result)

	//递归课堂练习2
	//猴子吃桃问题
	//此处调用progect1文件夹下monkey包下monkey函数
	// result := monkey.Monkey(1)
	// fmt.Println("reslut=", result)

	//此处调用progect1文件夹下swap包下swap函数
	// a := 10
	// b := 20
	// swap.Swap(&a, &b)
	// fmt.Printf("a=%v,b=%v", a, b)

	//init函数（待修改）
	//每一个源文件都会有一个init函数，在main之前被go运行框架调用执行
	//通常在init函数中完成初始化
	//全局变量>init函数>main函数
	//此处调用progect1文件夹下init包下init函数,来实现init函数下的全局变量的初始化
	// age := 90
	// fmt.Println("main...age=", age)
	// fmt.Println("Age=", init.Age, "Name=", Init.name)

	//匿名函数:没有名字的函数（希望使用一次）
	//使用方式：直接定义调用（调用一次之后就不能再用）；赋值给全局变量成为全局匿名函数（可以反复调用）；全局匿名函数
	//方式1：
	// 	res:=func (n1 int ,n2 int)int{
	// 		return n1+n2
	// 	}(n1的数值,n2的数值)
	//	fmt.println("res1=",res1)
	//方式2：
	// 	a:=func (n1 int,n2 int )int{//此处将函数赋值给变量a
	// 		return n1-n2
	// 	}
	// 	res2:=a(n1的数值,n2的数值)//这里就可以通过调用a来实现函数
	// 	fmt.Println("res2=",res2)
	//方式3：
	// 	res4:= Fun1(n1的数值,n2的数值)
	//	fmt.Println("res4=", res4)

	//闭包：一个函数与其相关的引用环境组合的一个整体
	//此处调用progect1文件夹下closure包下closure函数
	// m := closure.AddUpper()
	// fmt.Println(m(1)) //结果为：11
	// fmt.Println(m(2)) //结果为：13
	// fmt.Println(m(3)) //结果为：16

	//闭包课堂练习1
	//此处调用progect1文件夹下suffix包下Makesuffix函数
	//使用 string.HasSuffix  函数，判断某个字符串是否有指定的后缀名
	// n := suffix.Makesuffix(".jpg")   //此处指定一个后缀名
	// fmt.Println("处理后的文件名=", n("ok")) //将函数赋值给n后调用n函数

	//defer:
	//1、需要创建资源（数据库连接、文件句柄、锁等）防止函数执行完毕后，及时释放资源，（延时机制）
	//2、defer后可以继续使用创建资源
	//3、当函数完毕后，系统会一次从defer的栈中，取出语句，关闭资源。
	//4、好处:不用关心何时关闭资源
	//此处调用progect1文件夹下defer1包下Sum函数
	// res := defer1.Sum(10, 20)
	// fmt.Println("res=", res)

	//函数综合练习1
	//九九乘法表
	//此处调用progect1文件夹下multiplication包下Multiplication函数
	// var S int
	// fmt.Println("请输入行高")
	// fmt.Scanln(&S)
	// multiplication.Multiplication1(S)
	// fmt.Printf("%d,%d乘法表", S, S)

	//函数综合练习2
	//打印金字塔
	//此处调用progect1文件夹下pascal包下Pascal函数
	// var num int
	// fmt.Println("请输入要打印的金字塔的高度")
	// fmt.Scanln(&num)
	// pascal.Pascal(num)

	//函数综合练习3(待修改)
	//将一个（3*3）的二维数组进行转置

	//字符串中常用的函数(系统函数)：
	//1、统计字符串的长度、按照字节：len(str)[此处这个是内建函数]
	//此处调用progect1文件夹下usualfuncation文件夹下len_str包下Len_str1函数
	//len_str.Len_str1()

	//2、字符串遍历、同时处理有中文问题r:=[]rune(str)
	//此处调用progect1文件夹下usualfuncation文件夹下len_str包下Len_str2函数
	//len_str.Len_str2()

	//3、字符串转整数、 n,err:=strconv.Atoi("字符串")
	//此处调用progect1文件夹下usualfuncation文件夹下len_str包下Len_str3函数
	//len_str.Len_str3()

	//4、整数转字符串、 str=strconv.Itoa(整数)
	//此处调用progect1文件夹下usualfuncation文件夹下len_str包下Len_str4函数
	//len_str.Len_str4()

	//5、字符串转、 []byte: var bytes = []byte("hello go")
	//此处调用progect1文件夹下usualfuncation文件夹下len_str包下Len_str5函数
	//len_str.Len_str5()

	//6、[]byte 转 字符串：str=string([]byte{97,98,99})
	//此处调用progect1文件夹下usualfuncation文件夹下len_str包下Len_str6函数
	//len_str.Len_str6()

	//7、10进制转2,8,16进制： str=strconv.FormatInt(123,2)//2->8,16
	//此处调用progect1文件夹下usualfuncation文件夹下len_str包下Len_str7函数
	//len_str.Len_str7()

	//8、查找子串是否在指定的字符串中：string.Contains("seafood","foo")//结果为：true
	//此处调用progect1文件夹下usualfuncation文件夹下len_str包下Len_str8函数
	// len_str.Len_str8()

	//9、统计一个字符串有几个指定的子串：string.Count("ceheese","e")//结果为：4
	//此处调用progect1文件夹下usualfuncation文件夹下len_str包下Len_str9函数
	//len_str.Len_str9()

	//10、不区分大小写的字符串比较(==是区分字母大小写的):fmt.Println(string.EqualFold("abc","Abc"))//结果为：true
	//此处调用progect1文件夹下usualfuncation文件夹下len_str包下Len_str10函数
	// len_str.Len_str10()

	//11、返回子串在字符串第一次出现的index值，如果没有返回-1：string.Index("NLT_abc","abc")
	//此处调用progect1文件夹下usualfuncation文件夹下len_str包下Len_str11函数
	// len_str.Len_str11()

	//12、返回子串在字符串最后一次出现的index，如没有返回-1：string.LastIndex("go golang","go")
	//此处调用progect1文件夹下usualfuncation文件夹下len_str包下Len_str12函数
	// len_str.Len_str12()

	//13、将制定的子串替换成另外一个子串：string.Replace("go go hello","go","go语言",n) n可以指定你希望替换几个，如果n=-1表示全部替换
	//此处调用progect1文件夹下usualfuncation文件夹下len_str包下Len_str13函数
	//len_str.Len_str13()

	//14、按照指定的某个字符，为分隔标识符，将一个字符串拆分成字符串数组：string.Split("hello,world,ok",",")
	//此处调用progect1文件夹下usualfuncation文件夹下len_str包下Len_str14函数
	//len_str.Len_str14()

	//15、将字符串的字母进行大小写的转换：string.ToLower("Go")//go strings.ToUpper("Go")//Go
	//此处调用progect1文件夹下usualfuncation文件夹下len_str包下Len_str15函数
	//len_str.Len_str15()

	//16、将字符串左右两边的空格去掉：string.TrimSpace("tn a lon gopher ntrn")
	//此处调用progect1文件夹下usualfuncation文件夹下len_str包下Len_str16函数
	//len_str.Len_str16()

	//17、将字符串左右两边指定的字符去掉：string.Trim("!hello!","!") //["hello"] //将左边的！和""去掉
	//此处调用progect1文件夹下usualfuncation文件夹下len_str包下Len_str17函数
	//len_str.Len_str17()

	//18、将字符串左边指定的字符串去掉：string.TrimLeft("!hello!","!") //将左边的！和"" 去掉
	//19、将字符串右边指定的字符串去掉：string.TrimRight("!hello!","!")//将右边的！和"" 去掉

	//20、判断字符串是否以指定的字符串开头：string.HasPrefi("ftp://192.168.10.1","ftp")//true
	//此处调用progect1文件夹下usualfuncation文件夹下len_str包下Len_str20函数
	//len_str.Len_str20()

	//21、判断字符串是以指定的字符串结束：string.HasSuffix("NLT_abc.jpg","jpg")//false

	//使用时间函数（引入time包）：提供了时间的显示测量用的函数，采用公历

	//time包内专门用于表示时间的数据类型Time
	//now := time.Now() //now的类型就是time.Time
	// fmt.Printf("type =%T val =%v\n", now, now)
	// fmt.Printf("年=%v\n", now.Year())
	// fmt.Printf("月=%v\n", int(now.Month()))
	// fmt.Printf("日=%v\n", now.Day())
	// fmt.Printf("时=%v\n", now.Hour())
	// fmt.Printf("分=%v\n", now.Minute())
	// fmt.Printf("秒=%v\n", now.Second())

	//时间日期的格式化输出
	//方式一：固定输出    方式二：使用now.Format方法
	// fmt.Printf("当前的年月日时分秒 %d-%d-%d %d:%d:%d \n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	// dateStr := fmt.Sprintf("当前的年月日时分秒 %d-%d-%d %d:%d:%d \n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	// fmt.Printf("dateStr=%v", dateStr)

	// fmt.Printf(now.Format("2006/01/02 15:04:05")) //里面的时间格式是固定的格式

	//每隔一秒钟就打印一个数字
	// i := 0
	// for {
	// 	i++
	// 	fmt.Println(i)
	// 	time.Sleep(time.Second)
	// 	//time.Sleep(time.Millisecond * 100)//取0.1秒
	// 	if i == 10 {
	// 		break
	// 	}
	// }

	//获取Unix时间戳，Unix nano时间戳（使用此方法获取随机数）
	// fmt.Printf("unix时间戳=%v\n unix nano时间戳=%v", now.Unix(), now.UnixNano())

	//编写一段代码来统计某个函数的执行时间
	//此处调用progect1文件夹下usualfuncation文件夹下len_str包下Len_str21函数
	// start := time.Now().Unix()
	// len_str.Len_str21()
	// end := time.Now().Unix()
	// fmt.Printf("执行len_str21()所耗费的时间为%v\n", end-start)

	//错误处理
	//默认情况下，当错误发生后（panic）,程序自动退出（崩溃）
	//也可以捕捉错误，并行处理，保证程序的正常运行，还可以给给管理员一个提示（邮件，短信等）
	//Go语言不支持 try...catch...finally
	//Go语言引入的处理方法：defer,panic,recover
	//简单描述为：Go中抛出一个panic异常（既接受一个interface{}类型的值），然后在defer中通过recover捕获这个异常，然后在处理
	//举例说明：
	//此处调用progect1文件夹下usualfuncation文件夹下error包下Error函数
	// error.Error()
	// fmt.Println("main下面的代码...") //error函数内的函数如果不执行异常捕获，则不会执行该句

	//此处调用progect1文件夹下usualfuncation文件夹下error包下Error函数
	//用还函数读取配置文件init.config的信息
	//如果传入的信息不符合函数内定义的，则返回一个自定义的错误名称
	// err := error.ReadConf("confi.ini")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("后面代码继续执行")

	//函数循环练习1
	//循环打印输入的月份的天数[continue实现]
	//此处调用progect1文件夹下usualfuncation文件夹下exeo包下Exeo01函数
	//exeo.Exeo01()

	//函数循环练习2
	//生成一个1-100的随机数，反复判断是否猜中
	//此处调用progect1文件夹下usualfuncation文件夹下exeo包下Exeo02函数
	//exeo.Exeo02()

	//函数循环练习3
	//输入并判断是否是100以内的素数并且相加，且每行显示5个
	//此处调用progect1文件夹下usualfuncation文件夹下exeo包下Exeo03函数
	//exeo.Exeo03()

	//函数循环练习4
	//此处调用progect1文件夹下usualfuncation文件夹下exeo包下Exeo04函数
	//从1990.1.1起判断其中某一天是打渔还是晒网
	exeo.Exeo04()

}
