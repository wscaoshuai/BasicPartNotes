package main
import (
	"fmt"
	"sort"
)

func main(){
	//基本语法：
	// 1.var map变量名 map[keytype]valuetype
	//2.key的类型：key不可重复。bool ，数字，string，指针，channel，（很少有：接口，结构体，数组，）
	//			（通常为int，string）（不可用：slice，map，function）
	//3.valuetype的类型：与key的类型基本相同：通常为string，map，struct
	//4.map 不会分配内存，初始化需要make，分配内存后才能赋值和使用

	//map的使用细节
	// 1.map是引用类型，遵守引用类型传递的机制，在一个函数接收map，修改后，会直接修改原来的map
	// 2.map的容量达到后，在想增加map元素，会自动扩容，不会panic，map可以动态增长键值（key-value)
	// 3.map的value也经常使用struct类型，更适合管理复杂的数据（比前面的value是一个map更好）

	//案例演示：
	//定义一个map
	// var a map[string]string
	// a = make(map[string]string,10)//5.使用map前需要先make，make就是给map分配数据空间
	// a["no1"] = "宋江"
	// a["no2"] = "武松"
	// a["no3"] = "武松"//6.value的值可以相同
	// a["no1"] = "吴用"//7.key无序，且不能相同
	// fmt.Println(a)

	//map使用方式1
		//先声明再make
		//上面↑
	//map使用方式2
		//声明后直接make
		// cities := make(map[string]string)
		// cities["no1"] = "北京"
		// cities["no2"] = "上海"
		// cities["no3"] = "广州"
		// cities["no4"] = "深圳"
		// fmt.Println(cities)
	//map使用方式3
		//声明的同时直接赋值
		// var names map[string]string = map[string]string{
		// 	"name1" : "tom",
		// 	"name2" : "jack",
		// 	"name3" : "john",
		// }
		// fmt.Println("names:",names)

		//课堂练习：存放三个学生的信息（name,age,sex)
		// var studentMap = make(map[string]map[string]string)
		// studentMap["stu01"] = make(map[string]string,3)
		// studentMap["stu01"]["name"] = "tom"
		// studentMap["stu01"]["age"] = "20"
		// studentMap["stu01"]["address"] = "北京长安街"

		// studentMap["stu02"] = make(map[string]string,3)//此处不能少，否则panic
		// studentMap["stu02"]["name"] = "Jack"
		// studentMap["stu02"]["age"] = "21"
		// studentMap["stu02"]["address"] = "上海外滩"
		// fmt.Println(studentMap)

		//此处为遍↑历map示例
		//2.
		// for k1 , v1 := range studentMap{
		// 	fmt.Println(k1,v1)
		// 	for k2 , v2 := range v1{
		// 		fmt.Printf("\t k2=%v v2=%v\n",k2,v2)
		// 	}
		// }


		//map的增删查改
		// cities := make(map[string]string)
		// cities["no1"] = "北京"
		// cities["no2"] = "上海"
		// cities["no3"] = "广州"
		// cities["no4"] = "深圳"
		// fmt.Println(cities)
		//1.map的增加与更新
		// map["key"] = value//如果key还没有，就是增加，如果key存在就是修改
		// cities["no4"] = "杭州"//此处已经存在key:no4;所以当前操作为修改
		// fmt.Println(cities)

		//2.map的删除：
		//delete(map,"key"),delete是一个内置函数，如果key存在，则删除key-value,如果key不存在，操作不当，但是也不会报错
		// delete(cities,"no1")
		// fmt.Println(cities)
		//如何一次性删除所有的key
			//1.遍历所有的key，逐个删除
			//2.直接make一个新空间(推荐使用)
			// cities = make(map[string]string)
			// fmt.Println(cities)

		//3.map的查找
		// val , ok := cities["no2"]
		// if ok {
		// 	fmt.Printf("有no2,值为：%v\n",val)
		// }else{
		// 	fmt.Println("没有no2 key\n")
		// }

		//map的遍历（使用for-range,不能使用for循环）
		//示例1.
		// for k , v := range cities{
		// 	fmt.Printf("k=%v v=%v\n",k,v)
		// }

		//map的长度（可以用len）
		// fmt.Println("cities 有",len(cities),"对 key-value")

		//map的切片
		var monster  []map[string]string
		monster = make([]map[string]string,2)//切片的格式make([]类型，长度)
		if monster[0] == nil{
			monster[0] = make(map[string]string,2)
			monster[0]["name"] = "牛魔王"
			monster[0]["age"] = "500"
		}

		if monster[1] == nil{
			monster[1] = make(map[string]string,2)
			monster[1]["name"] = "玉兔精"
			monster[1]["age"] = "400"
		}
		//map的动态增加
		//可以使用append函数使map动态增加，而不是更改切片长度
		newMoster  := map[string]string{
			"name" : "新的妖怪~火云邪神",
			"age" : "200",
		}
		monster = append(monster,newMoster)
		fmt.Println(monster)

		//map的排序
		// 1.先将map的key放入进切片内
		// 2.对切片排序
		// 3.遍历切片，然后按照key来输出map的值
		var key []int
		for k,_ := range monster {
			key = append(key,k)
		}
		sort.Ints(key)
		fmt.Println(key)

		for _,k := range key{
			fmt.Printf("map[%v]=%v \n",k,monster[k])
		}

		
}

