package main

import (
	"fmt"
)

func main() {
	var name string
	var age int
	var sla float32
	var test bool
	fmt.Println("姓名：")
	fmt.Scanln(&name)
	fmt.Println("年龄：")
	fmt.Scanln(&age)
	fmt.Println("工资:")
	fmt.Scanln(&sla)
	fmt.Println("考试:")
	fmt.Scanln(&test)
	fmt.Printf("姓名%v 年龄%v  工资%v 考试%v", name, age, sla, test)
}
