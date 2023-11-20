package main

import (
	"fmt"
	"math/rand"
	"sort"
	_ "sort"
)

//基本语法：
// type 接口名 interface{
// 	method1(参数列表)返回值列表
// 	method2(参数列表)返回值列表
// }
//↓实现接口的所有方法
// func (t 自定义类型) method1(参数列表) 返回值列表 {
// 	实现方法
// }
// func (t 自定义类型) method2(参数列表) 返回值列表 {
// 	实现方法
// }

//基本操作
//1.接口可以定义一组方法，但不需要实现，且不能包含任何变量，到某个自定义类型（某结构体）要用的时候再写方法
//2.结构里的所有方法都没有方法体，接口的方法都是没有实现的方法，golang接口体现了多态和高内聚低耦合的思想
//3.golang的接口，不需要显示的实现，只需要一个变量，就可以实现接口类型中的所有方法，这个变量就实现了这个接口
//4.golang的接口，基于方法实现，不需要指定接口名，所以可以同时实现 同方法的不同接口(但是方法名，方法数量必须相同)

//注意细节
//1.接口本身不能创建实例，但是可以指向一个实现了该接口的自定义类型的变量（实例）
//2.接口中所有的方法都没有方法体，既都是没有实现的方法
//3.在golang中，一个自定义类型需要将某个接口的所有方法都实现，这样才能说这个自定义类型实现了该接口
//4.一个自定义类型只有实现了某个接口，才能将该自定义类型的实例(变量)
//5.只要是自定义数据类型，就可以实现接口，不仅仅是结构体类型
//6.一个自定义类型可以实现多个接口
//7.golang接口中不能有任何变量
//8.一个接口可以继承多个别的接口，如果要实现此接口，必须也要将继承的接口的方法全部实现
//9.interface默认类型是一个指针(引用类型)，如果没有对interface初始化就使用，那么会nil
//10.空接口interface{}没有任何方法，所以所有类型都实现了空接口。既可以把任何一个变量赋值给空接口

//案例1
// type Usb interface {
// 	Strat()
// 	Stop()
// }

// type Phone struct {
// }

// //让Phone实现USB接口的方法
// func (p Phone) Strat() {
// 	fmt.Println("手机开始工作")
// }

// func (p Phone) Stop() {
// 	fmt.Println("手机停止工作")
// }

// type Camera struct {
// }

// func (c Camera) Strat() {
// 	fmt.Println("相机开始工作")
// }

// func (c Camera) Stop() {
// 	fmt.Println("相机停止工作")
// }

// type Computer struct {
// }

// func (c Computer) Working(usb Usb) { //编写一个方法Working，接收Usb接口类型：（此处也是多态的特点体现）
// 	usb.Strat() //通过USB接口变量来实现调用Start和Stop方法
// 	usb.Stop()
// }

// func main() {
// 	//先创建结构体变量
// 	computer := Computer{}
// 	phone := Phone{}
// 	camera := Camera{}

// 	computer.Working(phone)
// 	computer.Working(camera)
// }

//
//
//
//
// 案例2
// type Hero struct { //声明Hero结构体
// 	Name string
// 	Age  int
// }

// type HeroSlice []Hero //声明结构体切片

// func (hs HeroSlice) Len() int { //实现一个“interface”接口
// 	return len(hs)
// }

// func (hs HeroSlice) Less(i, j int) bool { //less方法就决定了你用什么标准排序，(此处按照hero的年龄大→小来排序)
// 	return hs[i].Age > hs[j].Age //按照年龄
// 	//return hs[i].Name < hs[j].Name //按照姓名

// }

// func (hs HeroSlice) Swap(i, j int) { //交换
// 	hs[i], hs[j] = hs[j], hs[i]
// }

// func main() {
// 	// var intSlice = []int{1, -1, 10, 2, 7} //定义一个数组/切片
// 	// //要求对 intSlice切片进行排序
// 	// //1.冒泡排序
// 	// //2.也可以使用系统提供的方法（sort.Ints）
// 	// sort.Ints(intSlice)
// 	// fmt.Println(intSlice)

// 	//编写 对结构体进行遍历
// 	//1.冒泡排序。。。
// 	//2.系统提供方法（sort）

// 	//测试：案例2
// 	var heroes HeroSlice
// 	for i := 0; i < 10; i++ {
// 		hero := Hero{
// 			Name: fmt.Sprintf("英雄~%d:", rand.Intn(100)),
// 			Age:  rand.Intn(100),
// 		}
// 		heroes = append(heroes, hero)
// 	}

// 	// for _, v := range heroes { //查看排序前情况
// 	// 	fmt.Println(v)
// 	// }

// 	sort.Sort(heroes)          //调用sort.Sort
// 	for _, v := range heroes { //查看排序后情况
// 		fmt.Println(v)
// 	}

// }

//
//
//
//
//案例3
type Student struct {
	Name  string
	Age   int
	Score float64
}

type StudentSlice []Student

func (student StudentSlice) Len() int { //实现一个“interface”接口
	return len(student)
}

func (student StudentSlice) Less(i, j int) bool { //less方法就决定了你用什么标准排序，(此处按照学生的成绩大→小来排序)
	return student[i].Score > student[j].Score //按照成绩
	//return student[i].Age > student[j].Age //按照年龄
	//return hs[i].Name < hs[j].Name //按照姓名
}

func (student StudentSlice) Swap(i, j int) { //交换
	student[i], student[j] = student[j], student[i]
}

func main() {
	var student StudentSlice
	for i := 0; i < 10; i++ {
		stu := Student{
			Name:  fmt.Sprintf("学生~%d:", rand.Intn(60)),
			Age:   rand.Intn(20),
			Score: float64(rand.Intn(100)),
		}
		student = append(student, stu)
	}
	// sort.Sort(student)          //调用sort.Sort
	// for _, v := range student { //查看排序后情况
	// fmt.Println(v)
	//}

	sort.Sort(student)          //调用sort.Sort
	for _, v := range student { //查看排序后情况
		fmt.Println(v)
	}
}
