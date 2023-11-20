package mode03

import "fmt"

func Cal(n1 float64, n2 float64, operator byte) float64 { //实现+、-、*、/ 的自定义函数cal
	var res float64
	switch operator {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		fmt.Println("请输入正确的操作符")
	}

	return res
}
