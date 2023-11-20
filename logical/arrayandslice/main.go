package main

import (
	"fmt"
	_"math/rand"
	_"time"

)

//exeo 01 输出A-Z
// func main()  {
// 	var array01 = [26]byte{}
// 	for i:=0;i<26;i++{
// 		array01[i] = 'A'+ byte(i)
// 	}
// 	for i:=0;i<26;i++{
// 		fmt.Printf("%c\n",array01[i])
// 	}
// }

// //exeo 02 找出数组内最大值并打印出该下标号
// func main(){
// 	var array02 = [10]int64{1,-1,22,91,-34,25,19}
// 	// for a:=0;a<=len(array02);a++{
// 	// 	fmt.Scanln("请输入数据",&array02[a])
// 	// }//待修改
// 	max :=array02[0]
// 	maxIndex := 0
// 	for a:=0;a<len(array02);a++{
// 		if array02[a] > max{
// 			max = array02[a]
// 			maxIndex = a
// 		}
// 	}
// 	fmt.Printf("max：%v\nmaxIndex:%v",max ,maxIndex)
// }

// //exeo 03 求数组的和以及平均值 用for-range
// func main(){
// 	var array03 = [10]int{1,-1,22,91,-34,25,19}
// 		sum:=0
// 		for _, value := range array03{
// 		sum += value
// 	}
// 	fmt.Printf("sum=%v,avg=%v",sum,float64(sum)/float64(len(array03)))
// }

//随机生成5个数，并将其反转打印
// func main()  {
// 	var array04 = [5]int{}
// 	rand.Seed(time.Now().Unix())
// 	for x:=0; x<len(array04); x++{
// 		array04 [x] = rand.Intn(100)
// 	}
// 	fmt.Println(array04)
// 	var temp = 0 
// 	for x:=0; x<len(array04)/2; x++{
// 		temp = array04[len(array04)-1-x]
// 		array04[len(array04)-1-x] = array04[x]
// 		array04[x] = temp
// 	}
// 	fmt.Println(array04)
// }


//遍历切片方式1：
// func main()  {
// 	var arr [5]int = [...]int{10,20,30,40,50}
// 	slice := arr[1:3] //
// 	for i:=0; i<len(slice); i++{
// 		fmt.Printf("slice[%v]=%v",i,slice[i])
// 	}
// }

//遍历切片方式2：
// func main()  {
// 	var slice = [...]int{10,20,30,40,50}
// 	for index,value := range slice{
// 		fmt.Printf("i=%v,v=%v\n",index,value)
// 	}
// }

//字符串与切片的应用
// func main(){
// 	//string底层是一个byte数组，因此string也可以进行切片处理
// 	str := "hello@163.com"
// 	//使用切片处理获取到163.com
// 	slice := str[6:]
// 	fmt.Println(slice)
// }


// func main(){
// 	//string是不可变的，也就说不能通过str[0]='z'方式来修改字符串
// 	str := "hello"
// 	str[0] = 'z'//此处报错编译不通过，原因：string是不可变的
// }

// func main(){
// 	//如果需要修改string，可以先将string转换成[]byte/或者[]rune

// 	//将string转换成[]byte，可以装换成英文或者数字，不能转换成汉字，汉字占用两个字节
// 	str := "hello"
// 	arr := []byte(str)
// 	arr[0]='z'			
// 	str = string(arr)
// 	fmt.Println(str)
	
// 	//将string转换成[]rune就可以转换成汉字，因为[]rune是按字符处理的
// 	str1 := "hello"
// 	arr1 := []rune(str1)
// 	arr1[0]='好'			
// 	str1 = string(arr1)
// 	fmt.Println(str1)

// }

//切片练习1
//编写一个函数fbn(n int)将斐波那契数列放入切片内
// func fbn(n int) ([]uint64){
// 	fbnSlice := make([]uint64,n)
// 	fbnSlice[0] = 1
// 	fbnSlice[1] = 1
// 	for i:=2;i<n;i++{
// 		fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
// 	}
// 	return fbnSlice

// }

// func main(){
// 	fbnSlice := fbn(1000)
// 	fmt.Println(fbnSlice)
// }


// //二维数组练习
// func main(){
// 	var sorces [3][5]float64
// 	for i:=0;i<len(sorces);i++{
// 		for j:=0;j<len(sorces[i]);j++{
// 			fmt.Printf("请输入第%v个班的第%v个同学的成绩",i+1,j+1)
// 			fmt.Scanln(&sorces[i][j])
// 		}
// 	}
// 	fmt.Println(sorces)
// }


//Homework
//题目1.随机生成十个整数（一到100的范围）保存到数组，并顺序打印以及求平均值。
		//最大值的下标，并查找里面是否有55？
// func main()  {
// 	var array05 = [10]int{}
// 	rand.Seed(time.Now().Unix())
// 	for x:=0; x<len(array05); x++{
// 		array05 [x] = rand.Intn(100)
// 	}
// 	avg := 0 
// 	sum := 0 
// 	max := 0
// 	for y:=0; y<len(array05); y++{
// 		sum += array05[y]
// 		avg = sum / len(array05)
// 		fmt.Println(array05[y])
// 		if array05[y]>max {
// 			max = array05[y]
// 		}
// 		if array05[y] == 55{
// 			fmt.Println("该数组内含有数字55")
// 		}
// 	}
// 	fmt.Println("总和=",sum)
// 	fmt.Println("平均值=",avg)
// 	fmt.Println("最大值=",max)
// 	fmt.Printf("最大值下标=%v",array05)//待续该
// }



		
//题目2.一直有个排好序的数组，要求插入一个元素，最后打印该数组，顺序依然是升序

// func main(){
// 	var array06 = [5]int{}
// 	max := 0 
// 	for i:=1;i<=len(array06);i++{
// 		fmt.Printf("请输入该数组第%v个元素",i)
// 		fmt.Scanln(&array06[i-1])	
// 		for j:=0;j<len(array06);j++{
// 			if array06[j]>max{
// 				max = array06[j]
// 			}
// 		}
// 	}
// 	// slice := array06[:]
// 	// slice = append(slice,600)rfvfcdft
	
// 	fmt.Println("array06:",array06)

// 	// fmt.Println("slice:",slice)

// }
//题目3.定义一个3行4列的二维数组，逐个从键盘输入值，编写程序将四周的0去掉
// func main(){
// 	var i = 0 
// 	var j = 0
// 	var array07 = [3][4]int{}
// 	for i=0;i<3;i++{
// 		for j=0;j<4;j++{
// 			fmt.Scanln(&array07[i][j])
// 		}
// 	}

// 	for i=0;i<3;i++{
// 		for j=0;j<4;j++{
// 			fmt.Print(array07[i][j],"    ")
// 		}
// 		fmt.Println("   ")
// 	}
// }

//4.定义一个4行4列的二维数组，逐个从键盘输入值，然后将第一行和第四行的数据进行交换，将第二行与第三行交换。
// func main(){
// 	var i = 0 
// 	var j = 0
// 	// var v = 0
// 	var array08 = [4][4]int{}
// 	for i=0;i<4;i++{
// 		for j=0;j<4;j++{
// 			fmt.Scanln(&array08[i][j])
// 		}
// 	}

// 	for i=0;i<4;i++{
// 		array08[3][i] , array08[0][i] = array08[0][i] , array08[3][i]
// 	}

// 	for i=0;i<4;i++{
// 		for j=0;j<4;j++{
// 		fmt.Print(array08[i][j],"    ")
// 	}
// 	fmt.Println("   ")
// }

// }
//5.试保存13579五个奇数到数组，并且倒序打印

func main(){
	var temp = 0 
	var array09 = [5]int{}
	for i:=0;i<5;i++{
		fmt.Scanln(&array09[i])
	}

	fmt.Println("保存后的数组:",array09)

	for j := 0; j < 5/2; j++ {
        temp = array09[j]
        array09[j]= array09[5-1-j]
        array09[5-1-j] = temp
	}
	fmt.Println("倒叙后的数组:",array09)
	// for j:=4;j>=0;j--{
	// 	fmt.Println(array09[j])
	// }

}

//6.试写出实现查找核心代码吗，比如已知数组arr[10]string;里面保存了十个元素，现在要查找“aa”是否存在，
	//打印提示，如果有多个结果要表明对应下标
//7.随机生成10个整数（1-100）之内，使用冒泡排序进行排序，然后使用二分查找法，查找是否有90这个数，并显示下标，
	//如果没有则提示“找不到改数”
//8.编写一个函数，可以接收一个数组，该数组有5个数，请找出最大的和最小的数，对应的数组下标是多少？
//9.定义一个数组，并且给出8个整数，求该数组中大于平均值的数的个数，和小于平局值的数的个数

