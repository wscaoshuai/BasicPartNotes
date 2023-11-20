package main

import (
	"fmt"
	"reflect"
)

//反射最佳实例
type Monster struct {
	Name   string  `json:"name"`
	Age    int     `json:"monster_age"`
	Score  float64 `json:"成绩"`
	Sex    string  `json:"性别"`
	Adress string  `json:"地址"`
}

func (s Monster) Print() { //打印方法，显示s的值
	fmt.Println("---start---")
	fmt.Println(s)
	fmt.Println("---end---")
}

func (s Monster) GetSum(n1, n2 int) int { //求和方法，返回两个数的和
	return n1 + n2
}

func (s Monster) Set(name string, age int, score float64, sex string) { //接收4个值给monster赋值
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

func TestStruct(a interface{}) {
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()
	if kd != reflect.Struct {
		fmt.Println("expect strct")
		return
	}
	num := val.NumField()
	fmt.Printf("struct has %d fields\n", num)
	for i := 0; i < num; i++ { //变量结构体的所有字段
		fmt.Printf("Field %d: 值为=%v\n", i, val.Field(i))
		tagVal := typ.Field(i).Tag.Get("json") //获取struct标签
		if tagVal != "" {
			fmt.Printf("Field %d: tag为=%v\n", i, tagVal)
		}
	}
	numOfMethod := val.NumMethod()
	fmt.Printf("struct has %d method\n", numOfMethod)

	val.Method(1).Call(nil) //获取到第二个方法然后调用它

	var params []reflect.Value //调用结构体的第一个方法Method(0)
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := val.Method(0).Call(params)   //传入的参数是：[]reflect.Value
	fmt.Println("res = ", res[0].Int()) //返回结果是：[]reflect.Value
}

func main() {
	var a Monster = Monster{
		Name:   "Tom",
		Age:    22,
		Score:  99.9,
		Sex:    "男",
		Adress: "上海",
	}
	TestStruct(a)
}
