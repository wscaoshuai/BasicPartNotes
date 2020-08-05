package closure

func AddUpper() func(int) int {
	var n = 10
	return func(x int) int { //此处返回一个匿名函数，但是匿名函数引用了函数外的变量n，构成一个整体，就称为闭包
		n = n + x
		return n
	}

}
