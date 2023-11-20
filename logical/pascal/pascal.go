package pascal

import (
	"fmt"
)

func Pascal(L int) {
	//多重循环控制练习2(金字塔)
	//规律：星号（层数*2—1） 空格（总层数——当前层数）
	for I := 1; I <= L; I++ { //打印层数
		for J := 1; J <= L-I; J++ {
			fmt.Print(" ")
		}
		for K := 1; K <= 2*I-1; K++ { //打印空格数
			if K == 1 || K == 2*I-1 || I == L { //将中间判断为空的条件
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
