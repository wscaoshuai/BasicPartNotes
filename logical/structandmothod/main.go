package main

import(
	"fmt"
)

//方法练习题1：编写结构体（methodUtils),编程一个方法，方法不需要参数，在方法中打印一个10*8的矩形，在主函数中调用
// type MethodUtils struct{
// 	//无明确字段
// }

// func (methodUtils MethodUtils) Print() {
// 	for i:=0;i<11;i++{
// 		for j:=0;j<9;j++{
// 			fmt.Printf("*")
// 		}
// 		fmt.Println()
// 	}
// }

// func main()  {
// 	var methodUtils MethodUtils
// 	methodUtils.Print()
// }


//方法练习题2：编写一个方法，提供m和n两个参数，方法中打印一个m*n的矩形
// type MethodUtils struct{
// 	//无明确字段
// }

// func (methodUtils MethodUtils) Print(m int,n int) {
// 	for i:=0;i<m;i++{
// 		for j:=0;j<n;j++{
// 			fmt.Printf("*")
// 		}
// 		fmt.Println()
// 	}
// }

// func main()  {
// 	var methodUtils MethodUtils
// 	methodUtils.Print(10,2)
// }

//方法练习题3：编写一个方法，计算矩形的面积（可以接收长len,和宽width),将其作为方法返回值，在主函数中调用，并输出
// type MethodUtils struct{
// 	//无明确字段
// }

// func (methodUtils MethodUtils) area(len float64, width float64) (float64){
// 	return len * width
// }

// func main()  {
// 	var methodUtils MethodUtils
// 	areaRes := methodUtils.area(1.1,1.1)
// 	fmt.Println("面积为=",areaRes) 
// }

//方法练习题4：判断一个数是奇数还是偶数
// type MethodUtils struct{
// 	//无明确字段
// }
// func (methodUtils *MethodUtils)JudgeNum(num int){
// 	if num % 2 == 0{
// 		fmt.Println(num,"这个数是偶数")
// 	}else{
// 		fmt.Println(num,"这个数是奇数")
// 	}
// }

// func main()  {
// 	var methodUtils MethodUtils
// 	methodUtils.JudgeNum(20)
// }

//方法练习题5：根据行列，字符。打印对应行列的字符图形
// type MethodUtils struct{
// 	//无明确字段
// }
// func (methodUtils *MethodUtils)Print3(n int,m int,key string ){
// 	for i:=0;i<n;i++{
// 		for j:=0;j<m;j++{
// 			fmt.Print(key)
// 		}
// 		fmt.Println()
// 	}
// }

// func main()  {
// 	var methodUtils MethodUtils
// 	methodUtils.Print3(7,20,"J")
// }


//方法练习题6:实现计算器加减乘除功能(方式2)
// type Calcuator struct{
// 	Num1 float64
// 	Num2 float64
// }
// func (calcuator *Calcuator)getRes(operator byte )float64{
// 	res := 0.0
// 	switch operator {
// 	case '+':
// 		    res = calcuator.Num1 + calcuator.Num2
// 	case '-':
// 		    res = calcuator.Num1 - calcuator.Num2
// 	case '*':
// 		    res = calcuator.Num1 * calcuator.Num2
// 	case '/':
// 		    res = calcuator.Num1 / calcuator.Num2
// 	default:
// 		    fmt.Println("运算符输入有误。。。")
// 	} 
// 	return res 
// }
// func main()  {
// 	var calcuator Calcuator
// 	calcuator.Num1 = 1.2
// 	calcuator.Num2 = 2.2
// 	res := calcuator.getRes('+')
// 	fmt.Println("res=",res)
// }


//1.在MethodUtils结构体编写方法，接收整数，打印乘法表
// type MethodUtils struct{
// 	num int
// }
// func (methodUtils *MethodUtils)Print4(num int ){
// 	for i:=1;i<=num;i++{
// 		for j:=1;j<=i;j++{
// 			sum := i * j
// 			fmt.Printf("%d*%d=%d ",i,j,sum)
// 		}
// 		fmt.Println()
// 	}
// }

// func main()  {
// 	var methodUtils MethodUtils
// 	methodUtils.Print4(9)
// }

//2.编写方法，使给定的二维数组转置
type MethodUtils struct{
	
}
func (methodUtils *MethodUtils)Print5(){
	var Array [3][3]int = [3][3]int{{1,2,3},{4,5,6},{7,8,9}}
	for i:=0; i<len(Array);i++{
		for j:=0;j<i;j++{
		   Array[i][j],Array[j][i] = Array[j][i],Array[i][j]
		}
	  }
	  fmt.Println(Array)
}

func main()  {
	var methodUtils MethodUtils
	methodUtils.Print5()
}



