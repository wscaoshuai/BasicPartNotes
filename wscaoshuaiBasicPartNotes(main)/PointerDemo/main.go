package main

import (
	"fmt"
	"go_code/progect1/mode02" //此处引用了该地址的main.go文件，将此文件内定义的数字类型，字符型的内容开头大写则可以被引用（public),小写则不可以
)

func main() {
	fmt.Println(mode02.HeroName)
}
