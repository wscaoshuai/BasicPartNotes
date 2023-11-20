package main

import (
	"fmt"
)

//面向对象编程(OOP)应用实例
//1.声明(定义)结构体，确定结构体名称
//2.编写结构体的字段
//3.编写结构体方法

//1.学生案例
// 1.编写一个Student结构体，包含name,gender,age,id,score字段，分别为string，string，int，int，float64类型
// 2.结构体中声明一个say方法，返回string类型，返回信息汇总包含所有字段值。
// 3.在main方法中没创建Student结构体实例(变量)，并且访问say方法，并将调用结果打印输出

// type Student struct{
// 	name string
// 	gender string
// 	age int
// 	id int
// 	score float64

// }

// func (student *Student) say() string  {

// 	infoStr := fmt.Sprintf("student的信息 name=[%v] gender=[%v], age=[%v] id=[%v] score=[%v]",
// 			student.name, student.gender, student.age, student.id, student.score)
// 			return infoStr
// }

// func main()  {
// 	var stu = Student{
// 		name : "tom",
// 		gender : "male",
// 		age : 18,
// 		id : 1000,
// 		score : 99.99,
// 	}
// 	fmt.Println(stu.say())
// }

//2.长方体案例
//创建一个box结构体，长宽高键入，声明方法获取体积，创建box结构体变量，打印给定的立方体体积
// type Box struct{
// 	len float64
// 	width float64
// 	height float64
// }

// func (box *Box) getVolum()float64 {
// 	return box.len * box.width * box.height
// }

// func main()  {
// 	var box Box
// 	box.len = 1.1
// 	box.height = 2.2
// 	box.width = 1.1
// 	volumn := box.getVolum()
// 	fmt.Printf("体积为=%.2f",volumn)
// }

//3.景区案例
// 一个景区根据有人的年龄收取不同价格的门票，>18岁收取20元，
// 其他免费，请编写visitor结构体，根据年龄段决定能够买的门票价格并输出，
// type Visitor struct {
// 	Name string
// 	Age  int
// }

// func (visitor *Visitor) shouPrice() {
// 	if visitor.Age > 18 {
// 		fmt.Printf("游客的名字%v 游客的年龄%v 收费20元\n", visitor.Name, visitor.Age)
// 	} else {
// 		fmt.Printf("游客的名字%v 游客的年龄%v 免费\n", visitor.Name, visitor.Age)
// 	}
// }

// func main() {
// 	var visitor Visitor
// 	for {
// 		fmt.Println("请输入你的名字")
// 		fmt.Scanln(&visitor.Name)
// 		fmt.Println("请输入你的年龄")
// 		fmt.Scanln(&visitor.Age)
// 		visitor.shouPrice()
// 	}
// }

// 工厂模式:工厂模式是函数封装，实现工厂模式功能是创建方法
// golang 的结构体没有构造函数，通常可以使用工厂模式来解决
// 可以通过工厂模式来实现对某一个包内的结构体实例(变量)的访问与修改

// OOP三大特征：继承,封装,多态，go语言实现方法不一样，
//1.封装：就是把抽象的字段和对字段的操作封装在一起，数据被保护在内部，
//程序只能在其他包通过被授权的操作(方法)才能对字段进行操作(对电视机的操作就是典型的封装)
//1.1封装的理解与好处
// 可以隐藏显示细节
// 可以对数据进行验证，保证安全
// 1.2如何体现封装
// 对结构体中的属性进行封装
// 通过方法，包，实现

//2.继承：主要价值用：解决代码的复用性以及可维护性
//2.1接口：主要的价值：设计好各种规范(方法)，让其他之定义类型实现
//2.1.1接口比继承更加灵活
//2.1.2接口一定程度上实现了代码解耦

//3.多态：golang内的多态是以接口的形式实现的，可以按照统一的接口实现不同的功能，此时接口就可以呈现不同形态
//3.1多态数组：数组内不可以存放不同类型数据，但实现了该接口的任何结构体变量都可以存入该数组内
//3.2类型断言：由于接口是一般类型，不知道具体类型，若果要转成具体类型，就需要使用类型断言

//类型断言案例1：
// func main() {
// 	var a interface{}
// 	var point Point = Point{1, 2}
// 	a = point
// 	var b Pointb
// 	b = a.(Point)//此处就是断言类型，表示判断a是否指向point类型的变量，如果是就转成point类型，并赋值给b变量，否则报错
// 	fmt.Println(b)
// }
//类型断言（附带判断检测）
// func main() {
// 	var x interface{}
// 	var b2 float32 = 1.1
// 	x = b2

// 	y, ok := x.(float32)
// 	if ok {
// 		fmt.Println("convert success")
// 		fmt.Printf("y 的类型是%T 值是%v", y, y)

// 	} else {
// 		fmt.Println("convert fail")
// 	}
// 	fmt.Println("继续执行")
// }

//断言类型实例2
//编写一个函数，实现判断输入的参数是什么类型
func TypeJudge(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("第%v个参数是 bool 值是%v\n ", i+1, x)
		case float32:
			fmt.Printf("第%v个参数是 float64 值是%v\n", i+1, x)
		case float64:
			fmt.Printf("第%v个参数是 float64 值是%v\n", i+1, x)
		case int, int32, int64:
			fmt.Printf("第%v个参数是 int 值是%v\n", i+1, x)
		case string:
			fmt.Printf("第%v个参数是 string 值是%v\n", i+1, x)
		case Student:
			fmt.Printf("第%v个参数是 string 值是%v\n", i+1, x)
		case *Student:
			fmt.Printf("第%v个参数是 string 值是%v\n", i+1, x)
		default:
			fmt.Printf("第%v个参数是 unknown, 值是%v\n", i+1, x)
		}
	}
}

type Student struct {
}

func main() {
	var n1 float32 = 1.1
	var n2 float64 = 2.3
	var n3 int32 = 30
	var n4 string = "tom"
	address := "北京"
	n5 := 300
	Stu1 := Student{}
	Stu2 := &Student{}
	TypeJudge(n1, n2, n3, n4, address, n5, Stu1, Stu2)
}

//在代码之前的基础上，增加判断Student类型，和*Student类型

//案例：实现Person.go 不能随便查看人的年龄工资等隐私，并对输入的年龄进行合理验证
//OOP三大特征01:继承
//问题：学生成绩管理系统：引出继承的必要性
//不同年级同学正在考试，如何处理,要不断建立不同结构体，不同方法，出现代码冗余，不利于代码维护与扩展，所以通过继承的方式来解决
//golang解决继承的方式：在一个结构体内嵌入一个匿名结构体，就实现了结构体对匿名结构体的继承
//细节1：结构体可以使用嵌套匿名结构体所有的字段和方法，既首字母大写或者小写的字段，方法，都可以使用(仅限同包内)
//细节2.匿名结构体字段的访问可以简化(本结构体-->匿名结构体-->上级匿名结构体)
//细节3.结构体和匿名结构体同名(字段或者方法)，编译器采用就近访问原则，此时可以采用匿名结构体名访问
//细节4.结构体内有多个非同名(但有相同字段与方法)的匿名结构体时(匿名结构体内没有与之前相同的字段与方法)，
//     访问他必须要指定匿名结构体的名字
//细节5.如果一个struct嵌套了一个有名结构体(嵌套时给结构体前面，起个名字)，这种模式就是组合，如果是组合关系，
//     访问组合的结构体的字段或方法时，就必须带上结构体的名字
//细节6.嵌套匿名结构体后，也可以在创建结构体变量(实例)时，直接指定各个匿名结构体字段的值
// type Student struct {
// 	Name  string
// 	Age   int
// 	Score int
// }

// //Graduate,pupil共有的方法绑定到*student
// func (stu *Student) ShowInfo() {
// 	fmt.Printf("学生名=%v 年龄=%v 成绩=%v ", stu.Name, stu.Age, stu.Score)
// }
// func (stu *Student) SetScore(score int) {
// 	stu.Score = score
// }

// type Pupil struct { //小学生
// 	Student //此处为嵌入student匿名结构体
// 	//此处也可以直接嵌入数据类型(int string...)
// }
// type Graduate struct { //大学生
// 	Student //此处为嵌入student匿名结构体
// 	//此处也可以直接嵌入数据类型(int string...)
// }

// //小学生↓
// // func (p *Pupil) ShowInfo() {//此方法与Graduate相同故放置共有方法(父类)中↑
// // 	fmt.Printf("学生名=%v 年龄=%v 成绩=%v ", p.Name, p.Age, p.Score)
// // }

// // func (p *Pupil) SetScore(score int) {//此方法与Graduate相同故放置共有方法(父类)中↑
// // 	p.Score = score
// // }

// func (p *Pupil) testing() { //此方法为特有方法，应保留(也可以增加)
// 	fmt.Println("小学生正在考试中。。。。")
// }

// //大学生↓
// // func (p *Graduate) ShowInfo() {//此方法与Graduate相同故放置共有方法(父类)中↑
// // 	fmt.Printf("学生名=%v 年龄=%v 成绩=%v ", p.Name, p.Age, p.Score)
// // }

// // func (p *Graduate) SetScore(score int) {//此方法与Graduate相同故放置共有方法(父类)中↑
// // 	p.Score = score
// // }

// func (p *Graduate) testing() { //此方法为特有方法，应保留(也可以增加)
// 	fmt.Println("大学生正在考试中。。。。")
// }

// // func main() {//未使用匿名结构体之前写法
// // 	var pupil = &Pupil{
// // 		Name: "tom",
// // 		Age:  10,
// // 	}
// // 	pupil.testing()
// // 	pupil.SetScore(90)
// // 	pupil.ShowInfo()
// // }

// func main() {
// 	pupil := &Pupil{}
// 	pupil.Student.Name = "tom"
// 	pupil.Student.Age = 8
// 	pupil.testing()
// 	pupil.Student.SetScore(90)
// 	pupil.Student.ShowInfo()
// 	fmt.Println()

// 	Graduate := &Graduate{}
// 	Graduate.Student.Name = "jack"
// 	Graduate.Student.Age = 18
// 	Graduate.testing()
// 	Graduate.Student.SetScore(99)
// 	Graduate.Student.ShowInfo()
// }
