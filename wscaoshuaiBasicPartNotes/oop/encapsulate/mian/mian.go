package main

import (
	"fmt"
	"progect1/oop/encapsulate/model"
)

func main() {
	//p := model.NewPerson("smith")
	//p.SetAge(18)
	//p.SetSal(5000)
	//fmt.Println(p)
	//fmt.Println("name = ",p.Name,"age = ",p.GetAge(),"sal = ",p.GetSal())
	var account = model.NewAccount("jz11111", "99999", 40)
	if account != nil {
		fmt.Println("创建成功", account)
	} else {
		fmt.Println("创建失败")
	}

}
