package main
import (
	// "fmt"
	// //_unsafe//不使用的包在前方打上下划线
	// "strconv"

)

func main(){
	// 基本的整形数据类型
	// var a int = 100//int的表示范围与系统位数有关
	// var b int8 = 127//int8表示范围-128到127
	// var i uint16 = 255//uint16无符号类型0到255
	// var j byte = 255//byte表示范围0-255
	// fmt.Printf("a的数据类型%T\n",a)
	// fmt.Printf("b的数据类型%T\n",b)
	// fmt.Printf("i所占用的字节数是%d\n",i,unsafe.Sizeof(i))
	// fmt.Println("j=",j)


	// 小数类型使用(float32占用4个字节，64则占用8个字节，表示的范围和精度不同，均不受操作系统的影响)
	// var price float32 = 52.432
	// fmt.Println("price=",price)


	//字符类型直接用byte,当字符大小超过255使用int，需要格式化输出使用Printf
	// var c byte = 'C'
	// var d byte = 'D'
	// var e int64 = '海'
	// fmt.Println("这样输出的是码值:",c)
	// fmt.Printf("这样输出的是本身格式：%c\n",d)
	// fmt.Printf("汉字输出则用int类型：%d",e)


	//bool类型常用于判断（if，while等）占用一个字节的空间，只能用true和false，不能用0和非0表示真假。
	// var f bool = true
	// fmt.Println("f=",f)
	

	//string字符串类型：UTF-8编码表示Unicode文本,字符串一旦赋值不可更改，常用双引号包裹字符串(不能包含特殊字符)
	//					建议用反引号可以原生输出（实现防攻击，输出源码等）
	//字符串的拼接方式：可以使用’+’链接，多个链接目标可以将’+’留在上一行的末尾，
	// var g string = "广西北海18377966432"
	// fmt.Println("adress:",g)
	// var h string = `var g string = "广西北海18377966432" 
	// 				fmt.Println("adress:",g)`
	// fmt.Println("输出原生格式以及特殊字符等",h)


	//基本数据类型的转换:低精度转高精度同样要注明，转换并不会改变数值本身的数据类型.
	// var k int32 = 100
	// var l float32 = float32(k)
	// var m int8 = int8(k)
	// var n int64 = int64(k)
	// var o uint = uint(k)
	//fmt.Println("k=",k,"l=",l,"m=",m,"n=",n,"o=",o)
	// 强制转换会按照溢出等方式输出，不会编译报错
	// var p int64 = 99999
	// var q int8 = int8(p)
	// fmt.Println("q=",q)

	//基本数据类型和string类型的转换 
	// var r int = 99
	// var s float64 =23.456
	// var t bool = true 
	// // var u byte = 'u'
	// var str string //定义空的字符串
	// var str1 string //定义空的字符串
	// var str2 string //定义空的字符串
	//方法一:fmt.Sprintf
	// str = fmt.Sprintf("%d\n",r)
	// fmt.Println("str",str)
	// str = fmt.Sprintf("%f\n",s)
	// fmt.Println("str",str)
	// str = fmt.Sprintf("%t\n",t)
	// fmt.Println("str",str)
	// str = fmt.Sprintf("%d\n",u)
	// fmt.Println("str",str)
	//方法二strconv函数
	// str = strconv.FormatInt(int64(r),10)//10:是保留小数10位；64：表示小数是float64位
	// fmt.Printf("str type %T str=%q\n",str,str)
	// str1 = strconv.FormatFloat(s,'f',10,64)
	// fmt.Printf("str1 type %T str1=%q\n",str1,str1)
	// str2 = strconv.FormatBool(t)
	// fmt.Printf("str2 type %T str2=%q\n",str2,str2)


	//string类型转换基本数据类型
	// var str3 string = "true"
	// var u bool
	// u , _ =strconv.ParseBool(str3)//其中”b,_“解释：1.strconv.ParseBool(tstr)函数会返回两个值（value bool,err error）,本次使用前者，所以用“_”忽略后者
	// fmt.Printf("u type %T b=%v\n",u,u)
	// var str4 string = "1234567890"
	// var v int64
	// var w int
	// v,_ = strconv.ParseInt(str4,10,64)
	// fmt.Printf("v type %T v=%v\n",v,v)
	// fmt.Printf("w type %T w=%v\n",w,w)


	// var str5 string = "123.456"
	// var n float64
	// n,_ =strconv.ParseFloat(str5,64)
	// fmt.Printf("n type %T n=%v",n,n)


}