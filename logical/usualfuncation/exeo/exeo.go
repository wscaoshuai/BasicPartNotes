package exeo

import (
	"fmt"
	// "math/rand"
	// "time"
)

//循环打印输入的月份的天数[continue实现]
// func Exeo01() {
// 	var month int
// 	var year int64
// label1:
// 	for i := 0; i >= 0; i++ {
// 		fmt.Println("请输入要查询的月份")
// 		fmt.Scanln(&month)
// 		if month > 12 || month < 1 {
// 			fmt.Println("请输入正确的月份")
// 			break label1
// 		} else {
// 			switch month {
// 			case 1, 3, 5, 7, 8, 10, 12:
// 				fmt.Println("但前输入的月份一共有31天！")
// 			case 4, 6, 9, 11:
// 				fmt.Println("当前输入的月份一共有30天！")
// 			case 2:
// 				fmt.Println("请输入年份：")
// 				fmt.Scanln(&year)
// 				if year%400 == 0 || year/100 == 0 && year/4 == 0 {
// 					fmt.Println("当前输入的月份一共有28天！")
// 					break label1
// 				} else {
// 					fmt.Println("当前输入的月份一共有29天！")
// 					break label1
// 				}
// 			}
// 		}
// 	}
// }

// //函数循环练习2
// //生成一个1-100的随机数，反复判断是否猜中
// func Exeo02() {
// 	var k = 20
// 	for j := 1; j <= 100; j++ {
// 		rand.Seed(time.Now().UnixNano())
// 		j := rand.Intn(100)
// 		fmt.Println(j)
// 		if k == j {
// 			fmt.Println("生成的随机数与预设数值相等")
// 		} else {
// 			fmt.Println("生成的随机数与预设数值不相等")
// 		}
// 	}
// }

//函数循环练习3
//输入并判断是否是100以内的素数并且相加，且每行显示5个
// func Exeo03() {
// 	var l = 0
// 	var n = 0
// 	for m := 1; m <= 100; m++ {
// 		rand.Seed(time.Now().UnixNano())
// 		m := rand.Intn(100 + 1)
// 		if m%m == 0 || m%1 == m {
// 			l += m
// 		}
// 		fmt.Printf("%5d", m)
// 		n++
// 		if n%5 == 0 {
// 			fmt.Printf("\n")
// 		}
// 	}
// 	fmt.Println("100以内随机的素数总和是:", l)
// }

//函数循环练习4
//从1990.1.1起判断其中某一天是打渔还是晒网
func Exeo04() {
	var year int64
	var month int64
	var day int64
	var yearsum int64
	var allsum int64
	// var dayu int64
	// var shaiwang int64

	fmt.Println("请输入所需要判断的年份")
	fmt.Scanln(&year)
	fmt.Println("请输入所需要的判断的月份")
	fmt.Scanln(&month)
	fmt.Println("请输入所需要的判断的日期")
	fmt.Scanln(&day)
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		yearsum = (year - 1990) * 366
		switch month {
		case 1:
			day = 31
		case 2:
			day = 60
		case 3:
			day = 91
		case 4:
			day = 121
		case 5:
			day = 152
		case 6:
			day = 182
		case 7:
			day = 213
		case 8:
			day = 244
		case 9:
			day = 274
		case 10:
			day = 305
		case 11:
			day = 335
		case 12:
			day = 366
			allsum = yearsum + month + day
			if allsum%3 == 0 {
				fmt.Println("今天打渔")
			} else if allsum%2 == 0 {
				fmt.Println("今天晒网")
			}

		}
	} else {
		yearsum = (year - 1990) * 365
		switch month {
		case 1:
			day = 31
		case 2:
			day = 59
		case 3:
			day = 90
		case 4:
			day = 120
		case 5:
			day = 151
		case 6:
			day = 181
		case 7:
			day = 212
		case 8:
			day = 243
		case 9:
			day = 273
		case 10:
			day = 304
		case 11:
			day = 334
		case 12:
			day = 365
			allsum = yearsum + month
			if allsum%3 == 0 {
				fmt.Println("今天打渔")
			} else if allsum%2 == 0 {
				fmt.Println("今天晒网")
			}
		}
	}
}
