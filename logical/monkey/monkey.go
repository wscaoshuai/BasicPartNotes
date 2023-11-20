package monkey

import "fmt"

func Monkey(n int64) int64 {
	if n > 10 || n < 1 {
		fmt.Println("输入的天数有误")
		return 0 //表示输入的天数有误
	}
	if n == 10 {
		return 1
	} else {
		return (Monkey(n+1) + 1) * 2
	}
}
