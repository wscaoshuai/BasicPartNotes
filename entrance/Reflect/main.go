package main

func main() {
	//1.反射可以运行时动态获取变量的各种信息，比如变量的类型(type)，类别(kind)
	//2.结构体变量，还可以获取到结构体本身的信息，(包括结构体的字段、方法等)
	//3.通过反射，可以修改变量的值，可以调节关联的方法
	//4.使用反射，需要import("reflect")
	//5.函数
	//5.1 reflect.TypeOf(变量名)：获取变量的类型，返回reflect.Type类型
	//5.2 reflect.ValueOf(变量名)：通过获取变量的值，返回reflect.Value类型(该类型为结构体类型)

	//变量->接口->反射
	//1.接口->反射
	//rVal := reflect.ValueOf(b)//使用函数
	//2.反射->接口
	//iVal := reflect.Interface()//使用函数
	//3.接口->原变量类型
	//v := iVal.(Stu)//使用类型断言

	//反射的注意事项
	//1.reflect.Value.Kind,获取变量的类型，返回的是一个常量
	//1.1常量使用const修饰
	//1.2定义的时候必须初始化
	//1.3不能修改
	//1.4只能修饰bool，数值类型(int，float)，string类型
	//1.5语法：const identifier [type] = value
	//2.Type是类型，Kind是类别：
	//2.1:相同情况var num inr = 10 :num的Type是int而Kind也是int
	//2.2:不同情况var stu student stu 的Type是包名.Student而Kind是struct
	//3.Golang 中没有常量名必须字母大写的规定：TAX_RATE
	//4.仍然通过首字母的大小写来控制常量的访问范围
	//5.通过反射修改变量，注意使用SetXXX方法来设置需要对应的指针类型完成，以及resflect.Value.Elem()方法

}
