package model

import (
	"fmt"
)

//存取款案例展示2

type Account struct {
	AccountNo string
	Pwd       string
	Balance   float64
}

func NewAccount(accountNo string, pwd string, balance float64) *Account { //判断密码是否符合
	//匹配密码
	if len(accountNo) < 6 || len(accountNo) > 10 {
		fmt.Println("账号错误(长度不符合)")
		return nil
	}
	if len(pwd) != 6 {
		fmt.Println("密码错误(长度不符合~)")
		return nil
	}
	if balance < 20 {
		fmt.Println("余额数字不对")
		return nil
	}

	return &Account{
		AccountNo: accountNo,
		Pwd:       pwd,
		Balance:   balance,
	}
}

func (account *Account) WithDraw(money float64, pwd string) { //取款
	//匹配密码
	if pwd != account.Pwd {
		fmt.Println("密码错误")
		return
	}
	if money <= 0 || money > account.Balance {
		fmt.Println("金额错误")
		return
	}
	account.Balance -= money
	fmt.Println("存款成功")
}

func (account *Account) Query(pwd string) { //查询余额
	//匹配密码
	if pwd != account.Pwd {
		fmt.Println("密码错误")
		return
	}
	fmt.Printf("你的账号：%v余额为%v", account.AccountNo, account.Balance)

}
