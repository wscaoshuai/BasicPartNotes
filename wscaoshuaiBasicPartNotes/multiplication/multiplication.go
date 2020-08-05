package multiplication

import (
	"fmt"
)

func Multiplication1(S int) {
	//多重循环控制练习3(九九乘法表)
	// var A int64   //行
	// var B int64   //列
	//var sum int64 //和
	for A := 1; A <= S; A++ {
		for B := 1; B <= A; B++ {
			sum := A * B
			fmt.Printf("%v*%v=%v \t", A, B, sum)
		}
		fmt.Printf("\n")
	}

}
