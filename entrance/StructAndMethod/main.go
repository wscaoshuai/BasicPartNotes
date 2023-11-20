package main

import(
	"fmt"
	_"encoding/json"
)

//结构体与结构体变量（实例）的区别与联系
//1.结构体是自定义的数据类型，代表一类事物
//2.结构体变量（实例）是具体的，代表一个具体变量



//1.定义一个Cat结构体
//  type Cat struct{
// 	 Name string//区分大小写
// 	 Age int
// 	 Color string 
//  }

// func main()  {
// 	var cat1 Cat  
// 	cat1.Name = "小白"
// 	cat1.Age = 3
// 	cat1.Color = "白色"
// 	fmt.Println("Cat1=",cat1)
// }



//2.定义一个Person结构体
// type Person struct{
// 	Name string
// 	Age int 

// }

// func main()  {
// 	//方式1：
// 	var person1 Person
// 	person1.Name = "jack"
// 	fmt.Println(person1)
	
// 	//方式2：
// 	person2 := Person{}//也可以直接赋值
// 	person2.Name = "tom"
// 	fmt.Println(person2)

// 	//方式3：
// 	var p3 *Person = new(Person)
// 	(*p3).Name = "mary"
// 	p3.Name = "abc"//等价于↑
// 	fmt.Println(*p3)

// 	//方式4：
// 	var person *Person = &Person{}//也可以直接赋值
// 	(*person).Name = "scort"
// 	person.Name = "scort~"//等价于↑
// 	fmt.Println(*person)
// }


//结构体细节1：
// type Point struct{
// 	x int 
// 	y int 
// }

// type Rect struct{
// 	leftUp,rightDown Point
// }

// func main()  {
// 	r1 := Rect{Point{1,2},Point{3,4}}
// 	//r1 有4个int，在内存中式连续分布的。
// 	//打印地址
// 	fmt.Printf("r1.leftUp.x 地址=%p r1.leftup.y 地址%p r.rightDown.x 地址=%p r1.rightDown.y 地址=%p",&r1.leftUp.x,
// 			&r1.leftUp.y,&r1.rightDown.x,&r1.rightDown.y)
// }



//结构体细节2：结构体使用户单独定义的类型，和其他类型进行转换时需要有完全相同的字段（名字，个数，类型）
// type A struct{
// 	Num int 
// }
// type B struct{
// 	Num int 
// }
// func main()  {
// 	var a A
// 	var b B
// 	//a = b  //此处类型不同，不能直接转换 
// 	fmt.Println(a,b)
// }



//结构体细节3：struct的每个字段上，可以写一个tag，该tag可以通过反射机制获取，常见场景：序列化与反序列化，将struct
			//变量进行json处理

// type Monster struct{
// 	Name string `json:"name"`//结构体的标签
// 	Age int		`json:"age"`
// 	Skill string `json:"skill"`
// }
// func main()  {
// 	//创建一个monster 变量
// 	monster := Monster{"牛魔王",500,"芭蕉扇"}
// 	//将monster变量序列化为json格式字符串
// 	jsonStr,err :=  json.Marshal(monster)//此处用到了反射  将大写的Name装换成小写name
// 	if err != nil {
// 		fmt.Println("json 错误处理",err )
// 	}
// 	fmt.Println("jsonStr",string(jsonStr)) //此处需转换

// }



//方法：
// func(recevier type)methodName(参数列表) (返回值列表){
// 	方法体
// 	return 返回值
// }

// 1.参数列表：表示方法输入
// 2.recover type:表示这个方法和type这个类型进行绑定，或者说该方法作用余type类型
// 3.recover type：type可以是结构体，也可以其他的自定义类型
// 4.recover：就是type类型的一个变量(实例),比如：Person结构体的一个变量（实例）
// 5.参数列表：表示方法输入
// 6.返回值列表：表示返回的值，可以多个
// 7.方法主体：表示为了实现某个功能代码块
// 8.return语句：不是必须的

//方法Method是作用在一个指定的数据类型上面，struct包括自定义类型都可以有方法
//方法的声明与调用
// 1.结构体类型是值类型，在方法调用中，遵守值类型的传递机制，是值拷贝传递方式
// 2.如果希望修改结构体变量的值，可以通过结构体指针的方式处理
// 3.Golang中的方法作用在指定的数据类型上(即：和指定的数据类型绑定)，因此自定义类型，都可以有方法，不仅仅是struct
// 4.方法的访问控制的规则，和函数一样，方法名首字母小写只能在本包内访问，大写则相反。
// 5.如果一个变量实现了String()这个方法，那么fmt.Println默认会调用这个变量的String()进行输出
//变量调用方法时，相同：其调用机制和函数一样；
//不同：该变量本身也会作为一种参数传递到方法：如果变量是值类型，那么则进行值拷贝，如果变量时引用类型，则是地址拷贝

// type Person struct {//定义一个结构体Person
// 	Name string
// }
// func (person Person)test(){  //Person结构体的方法test, (person Person):test方法和Person类型绑定
// 	fmt.Println("test()",person.Name)
// }

// //1.给Person 结构体添加speak方法
// func (p Person) speak() {
// 	fmt.Println(p.Name,"是一个好人")
// }

// //2.给Person添加一个计算的方法
// func (p Person) jisuan() {
// 	res := 0
// 	for i:=1;i<=1000;i++{
// 		res += i
// 	}
// 	fmt.Println(p.Name,"计算结果：",res)
// }

// //3.给Person添加一个计算2的方法
// func (p Person) jisuan2(n int) {
// 	res := 0
// 	for i:=1;i<=n;i++{
// 		res += i
// 	}
// 	fmt.Println(p.Name,"计算结果：",res)
// }

// //4.给Person添加一个求和的方法（有返回值）多个返回值用括号括起来
// func (p Person) getsum(x int ,y int ) int  {
// 	return x + y
// }
// func main()  {
// 	var p Person
// 	p.Name = "tom"
// 	p.test()//调用方法
// 	fmt.Println("main() p.Name=",p.Name)
// 	p.speak()//调用方法
// 	p.jisuan()//调用方法
// 	p.jisuan2(10)//调用方法
// 	res := p.getsum(10,20)
// 	fmt.Println("res=",res)
// }

// type Circle struct{
// 	radius float64
// }

// func (circle Circle) area()float64{
// 	return 3.14 * circle.radius * circle.radius
// }

// func main()  {
// 	var circle Circle
// 	circle.radius = 4.0
// 	res := circle.area()
// 	fmt.Println("面积是：",res)

// }

type Student struct{
	Name string
	Age int
}
func (stu *Student) String()string  {
	str:= fmt.Sprintf("Name=[%v] Age=[%v]",stu.Name,stu.Age)
	return str
}
func main()  {
	stu := Student{
		Name : "tom",
		Age : 20,
	}
	fmt.Println(&stu)
}


//方法和函数的区别
// 1.调用方式不一样
// 函数调用方式：  函数名（实参列表）
// 方法调用方式：  变量.方法名（实参列表）
// 2.对于普通函数，接受者为值类型时，不能将指针类型的数据直接传递，反之
// 3.对于方法（如struct的方法），接受者为值类型时，可以直接用指针类型的变量调用方法，反之也可以
//不管调用形式如何，真正决定是值拷贝还是地址拷贝，是看这个方法和哪个类型绑定
//如果是和值类型，比如(p Person),则是值拷贝，如果和指针类型，比如(p *Person)则是地址拷贝
