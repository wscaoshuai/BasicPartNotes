package main

import (
	"fmt"
)


func main() {
	//数组
	// //1.定义数组
	// var hens [6]float64
	// //2.给数组的元素赋值，下标从0开始，不要数组越界
	// hens[0] = 3.0 //hens数组的第1个元素
	// hens[1] = 5.0 //hens数组的第2个元素
	// hens[2] = 1.0 //hens数组的第3个元素
	// hens[3] = 3.4 //hens数组的第4个元素
	// hens[4] = 2.0 //hens数组的第5个元素
	// hens[5] = 6.0 //hens数组的第6个元素
	// //3遍历数组求出总体重
	// totalweight := 0.0
	// for i := 0; i < len(hens); i++ {
	// 	totalweight += hens[i]
	// }
	// //4.求出平均值
	// avgweight := fmt.Sprintf("%.2f", totalweight/float64(len(hens)))
	// fmt.Printf("totalweight=%v avgweght=%v", totalweight, avgweight)

	//示例：
	// var intArr [3]int
	// fmt.Println(intArr)
	// fmt.Printf("intArr的地址=%p intArr[0]的地址=%p", &intArr, &intArr[0]) //整个数组的地址0xc0420480a0,首位地址0xc0420480a0

	//数组的使用  视频p146
	//从终端循环输入5个成绩，保存到float64中，并输出
	// var score [5]float64
	// for i := 0; i < len(score); i++ {
	// 	fmt.Printf("请输入第%d个元素值", i)
	// 	fmt.Scanln(&score[i])
	// }
	// //遍历数组
	// for j := 0; j < len(score); j++ {
	// 	fmt.Println("第%d数字为：\n", j, score[j])
	// }

	//四种数组初始化的方式
	//1.
	// var numArr01 [3]int = [3]int{1, 2, 3}
	// fmt.Println("numArr01=", numArr01)
	// //2.
	// var numArr02 = [3]int{1, 2, 3}
	// fmt.Println("numArr02=", numArr02)
	// //3.这里【。。。】是必须的写法
	// var numArr03 = [...]int{1, 2, 3}
	// fmt.Println("numArr03=", numArr03)
	// //4.
	// var numArr04 = [...]int{0: 1, 1: 2, 2: 3}
	// fmt.Println("numArr04=", numArr04)
	// //5.
	// strArr05 := [...]string{1: "jack", 2: "tom", 3: "mary"}
	// fmt.Println("strArr05=", strArr05)

	//数组遍历
	//1.常规遍历
	//2.for-range
	//基本语法：for index ,value:= range array01{
	//  			...
	//			}
	// 1.第一个返回值index是数组的下标
	// 2.第二个value是在该下标位置的值
	// 3.他们都是仅在for循环内部可见的局部变量
	// 4.遍历数组元素的时候，如果不想使用下标index，可以直接把下标index表为下划线_
	// 5.index和value的名称不是固定的，即程序员可以自定义，一般命名为index和value
	//案例演示：
	// heroes := [...]string{"宋江","吴用","卢俊义"}

	// for index , value := range heroes{
	// 	fmt.Printf("index=%v,value=%v",index,value)
	// }

	// 数组的使用注意事项和细节
	// 1.数组是多个相同类型的数据的组合，一个数组一旦声明/定义了，其长度是固定的，不能动态变化
	// 2.var arr []int 这时arr就是一个slice切片
	// 3.数组中的元素可以是任何数据类型，包括值类型和引用类型，但是不能混用
	// 4.数组创建后，如果没有赋值，就有默认值
	// 	数值类型数组：默认值为0
	// 	字符串数组：默认值为""
	// 	bool数组：默认值为false
	// 5.使用数组的步骤：1.声明数组并开辟空间；2.给数组的各个元素赋值；3.使用数组
	// 6.使用数组的下标是从零开始的
	// 7.数组下标必须在指定的范围内使用，否则会报panic：数组越界，比如var arr[5]int 则有效下标为0-4
	// 8.Go的数组属性值类型，在默认情况下是值传递，因此会进行值拷贝，数组间不会相互影响
	// 9.如想在其他函数中，去修改原来的数组，可以使用引用传递（指针方式）
	
	
	//切片
	//切片是数组的一个引用，因此切片是引用类型，在进行传递时，遵守引用传递的机制
	//切片的使用和数组类似，遍历切片、访问数组的元素和求切片的长度都是一样的
	//切片的长度是可以变化的，因此切片是一个可以动态变化的数组
	//切片的定义
	//var 变量名 []类型
	//切片的基本用法
	// var intArr [5]int = [...]int{11,22,33,44,55}
	// slice := intArr[1:3]  //表示slice(切片名,任意)引用到intArr这个数组：
	// 					  //下标为1开始到3结束(不包含3),如果3为最后一个元素则包含
	// fmt.Println("slice的元素是：",slice)
	// fmt.Println("slice的大小：",len(slice))
	// fmt.Println("slice的容量：",cap(slice))//切片得容量是可以动态变化的

	//切片使用的三种方式：
	//1.定义一个切片，然后让切片去引用（数组随着切片改变）一个已经创建好的数组，不然则为空不使用，

	//2.通过make来创建切片
		//基本语法：var 切片名 []type = make([]type数据类型,len切片大小,[cap]切片容量)

	//3.定义一个切片，直接指定具体数组，使用原理类似make的方式
		//基本语法：var 切片名 []string = []string{"tom","jack","mary"}
		//切片仍然可以切片

	
	//切片遍历的两种方式
	//1.for循环常规方式遍历
	
	//2.for-range结构遍历切片

	//切片的使用注意事项和细节
	//切片初始化时仍不能越界，范围在[0-lrn(arr)]之间，但是可以动态增长
	//var slice = arr[0:end] 可以简写：var slice = arr[:end]
	//var slice = arr[strat:len(arr)]可以简写成: var slice = arr[start:]
	//var slice = arr[0:len(arr)]可以简写成： var slice = arr[:]

	//切片的动态增加
	//用append内置函数实现对切片的动态追加元素或切片（同类型多个切不能是数组） 示例：
	// var slice2 []int = []int{100,200,300}
	// slice2 = append(slice2,400,500,600)
	// fmt.Println("slice2:",slice2)

	// slice2 = append(slice2,slice2...)
	// fmt.Println("slice2:",slice2)
	//切片append操作的底层原理分析：
		//1.切片append操作的本质是对数组扩容
		//2.go底层会创建一下新的数组newArr（安装扩容后大小）
		//3.将slice原来包括的元素拷贝到新的数组newArr
		//4.slice重新引用到数组newArr
		//5.newArr是在底层来维护，程序员不可见
	
	//切片的拷贝
	//切片使用copy内置函数完成拷贝
	// var slice3[]int = []int{1,2,3,4}
	// var slice4 = make([]int,10)
	// fmt.Println(slice3)
	// copy(slice4,slice3)
	// fmt.Println(slice4)
	//同为切片类型才可以拷贝
	//slice3以及slice4是独立的数据空间，互不影响
	//拷贝操作不会因为切片的容量大小而报错


	//string和slice的关系
	//1.string的底层是一个byte数组，因此string也可以进行切片处理
	//2.string和切片在内存的形式，以“ABCD”
	//3.string是不可变的，也就说不能通过str[0]='z'方式来修改字符串
	//4.如果需要修改字符串，可以先将string->[]byte/或者[]rune修改重新转成string
	

	//多维数组---二维数组
	// 1.基本语法：var 数组名[大小][大小]类型
	// 2.先声明再赋值:var arr[2][3]int,在赋值（有默认值比如int为0）
	// 3.直接进行初始化：var 数组名[大小][大小]类型 = [大小][大小]类型{{初值...},{初值...}}
	//先声明定义
	var arr [4][6]int 
	//再进行复制（4行6列的二维数组初始值为0）
	arr[1][2] = 1
	arr[2][1] = 2
	arr[1][3] = 1
	
	for i:=0; i<4; i++{
		for j:=0; j<6; j++{
			fmt.Print(arr[i][j]," ")
		}
		fmt.Println()
		fmt.Println("")
	}

	//二维数组的遍历
	//1.双层for循环完成遍历

	//2.for——range方式完成遍历
	for i,v := range arr{
		for j,v2 := range v{
			fmt.Printf("arr[%v][%v]=[%v]\t",i,j,v2)
		}
		fmt.Println()
	}


	
}
