package cal

func addUpper(n int) int { //被测试函数01
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}

func GetSub(n1 int, n2 int) int { //被测试函数02
	return n1 - n2
}
