package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string  // `json:"monster_name"`     //指定结构体tag标签，如果未指定则显示输出原定义的
	Age      int     //`json:"monster_age"`      //此处指定的标签，是使用到了反射机制
	Birthday string  //`json:"monster_brithday"` //可以指定小写，以及各种名称，结构体内的变量名无法做到
	Sal      float64 //`json:"monster_sal"`
	Skill    string  //`json:"monster_skill"`
}

////////////////////////////////////////////////////////////////////////////////////////////////
//01将结构体进行序列化
// func testStruct() {
// 	monstr := Monster{
// 		Name:     "Tom",
// 		Age:      22,
// 		Birthday: "2000-02-20",
// 		Sal:      80000.0,
// 		Skill:    "Golang",
// 	}
// 	data, err := json.Marshal(&monstr) //此处将monster序列化
// 	if err != nil {
// 		fmt.Printf("序列号错误 err=%v\n", err)
// 	}
// 	fmt.Printf("monster序列化后=%v\n", string(data)) //输出序列化后的结果
// }

////////////////////////////////////////////////////////////////////////////////////////////////
//02将map进行序列化
// func testMap() {
// 	var a map[string]interface{}
// 	a = make(map[string]interface{})
// 	a["name"] = "jack"
// 	a["age"] = 20
// 	a["adress"] = "NewYork"

// 	data, err := json.Marshal(a)
// 	if err != nil {
// 		fmt.Printf("序列化错误 err=%v\n", err)
// 	}
// 	fmt.Printf("a[map]序列化后=%v\n", string(data)) //输出序列化后的结果
// }

////////////////////////////////////////////////////////////////////////////////////////////////
//03将切片进行序列化
// func testSlice() {
// 	var slice []map[string]interface{} //定义一个切片

// 	var m1 map[string]interface{}
// 	m1 = make(map[string]interface{}) //使用map前必须进行make
// 	m1["name"] = "mary"
// 	m1["age"] = 21
// 	m1["adress"] = "北京"
// 	slice = append(slice, m1)

// 	var m2 map[string]interface{}
// 	m2 = make(map[string]interface{}) //使用map前必须进行mak
// 	m2["name"] = "jhon"
// 	m2["age"] = 20
// 	m2["adress"] = "上海" //此处也可以换成数组(因为定义的是接口)：m2["adress"] = [2]string{"上海","深圳"}
// 	slice = append(slice, m2)

// 	data, err := json.Marshal(slice)
// 	if err != nil {
// 		fmt.Printf("序列化错误 err=%v\n", err)
// 	}
// 	fmt.Printf("切片序列化后=%v\n", string(data)) //输出序列化后的结果
// }

// ////////////////////////////////////////////////////////////////////////////////////////////////
// //04基本数据类型序列化(意义不大)
// func testFloat64() {
// 	var num1 float64 = 123.456
// 	data, err := json.Marshal(num1)
// 	if err != nil {
// 		fmt.Printf("序列化错误 err=%v\n", err)
// 	}
// 	fmt.Printf("num1序列化后=%v\n", string(data)) //输出序列化后的结果
// }

////////////////////////////////////////////////////////////////////////////////////////////////
//反序列化细节
//01反序列化之前后要保证数据类型的一致
//02通过其他方式获取的字符串，不需要做转义处理，只有自己写的才需要

//01结构体反序列化
// func UnmarshalStruct() {
// 	//↓str内容在实际开发中是通过网络传输，或者是读取文件获取到的
// 	str := "{\"name\":\"Tom\",\"Age\":500,\"Birthday\":\"1999-01-01\",\"Sal\":8000,\"Skill\":\"Golang\"}"
// 	var monster Monster //创建一个Monster实例
// 	err := json.Unmarshal([]byte(str), &monster)
// 	if err != nil {
// 		fmt.Printf("unmarshal err=%v\n", err)
// 	}
// 	fmt.Printf("反序列化后 monster=%v", monster)
// }
////////////////////////////////////////////////////////////////////////////////////////////////
//02map反序列化
// func UnmarshalMap() {
// 	//↓str内容在实际开发中是通过网络传输，或者是读取文件获取到的
// 	str := "{\"adress\":\"北京\",\"age\":22,\"name\":\"jack\"}"
// 	var a map[string]interface{} //在反序列化map时，所以不需要make就可以使用，已经被封装到unmarshall
// 	err := json.Unmarshal([]byte(str), &a)
// 	if err != nil {
// 		fmt.Printf("unmarshal err=%v\n", err)
// 	}
// 	fmt.Printf("反序列化后 a=%v", a)
// }

////////////////////////////////////////////////////////////////////////////////////////////////
func UnmarshalSlice() {
	str := "[{\"address\":\"北京\",\"age\":\"21\",\"name\":\"mary\"}," +
		"{\"adress\":[\"洛杉矶\",\"夏威夷\"],\"age\":\"20\",\"name\":\"jhon\"}]"
	var slice []map[string]interface{} //在反序列化map时，所以不需要make就可以使用，已经被封装到unmarshall
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后 a=%v", slice)
}

////////////////////////////////////////////////////////////////////////////////////////////////

func main() {
	// //结构体序列化
	// testStruct()
	// //map序列化
	// testMap()
	// //切片序列化
	// testSlice()
	// //基本数据类型序列化
	// testFloat64()
	// //结构体反序列化
	// UnmarshalStruct()
	// //map反序列化
	//UnmarshalMap()
	UnmarshalSlice()
}
