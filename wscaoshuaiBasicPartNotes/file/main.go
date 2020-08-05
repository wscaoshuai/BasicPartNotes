package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//定义一个结构体，用于保存统计结果
type CharCount struct {
	ChCount    int //记录英文个数
	NumCount   int //记录数字个数
	SpaceCount int //记录空格个数
	OtherCount int //记录其他字符个数
}

//统计文件字符数量
func main() {
	//打开一个文件，创建一个Reader(缓存)
	//每读一行，就统计各个字符有多少
	//然后返回结果保存到一个结构体中
	fileName := "c:/test/test02.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("open file err=%v", err)
		return
	}
	defer file.Close() //关闭句柄防止内存泄露

	var count CharCount             //定义一个结构体实例
	reader := bufio.NewReader(file) //创建一个Reader
	for {                           //循环读取
		str, err := reader.ReadString('\n')
		if err == io.EOF { //末尾终止
			break
		}
		for _, v := range str {
			switch { //此处switch用法为switch分支结构，不用判断v
			case v >= 'a' && v <= 'z':
				fallthrough //此处为穿透
			case v >= 'A' && v <= 'Z':
				count.ChCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			default:
				count.OtherCount++

			}
		}
	}

	//输出统计结果，查看是否正确
	fmt.Printf("字母，数字，空格，其他；的数量分别是：%v,%v,%v,%v",
		count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)
}
