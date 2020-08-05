package main

import (
	"fmt"

//存取款案例展示1

type Account struct{
	AccountNo string
	Pwd string
	Balance float64
}

func (account *Account) Deposite (money float64,pwd string) {//存款
	//匹配密码
	if pwd != account.Pwd{
		fmt.Println("密码错误")
		return
	}
	if money <= 0 {
		fmt.Println("金额错误")
		return
	}
	account.Balance += money
	fmt.Println("存款成功")
}

func (account *Account) WithDraw (money float64,pwd string) {//取款
	//匹配密码
	if pwd != account.Pwd{
		fmt.Println("密码错误")
		return
	}
	if money <= 0 || money > account.Balance{
		fmt.Println("金额错误")
		return
	}
	account.Balance -= money
	fmt.Println("存款成功")
}

func (account *Account) Query (pwd string) {//查询余额
	//匹配密码
	if pwd != account.Pwd{
		fmt.Println("密码错误")
		return
	}
	fmt.Printf("你的账号：%v余额为%v",account.AccountNo,account.Balance)

}

func main()  {
	account := Account{
		AccountNo : "gs111111",
		Pwd : "666666" ,
		Balance : 100.0,
	}
	account.Query("666666")

}
