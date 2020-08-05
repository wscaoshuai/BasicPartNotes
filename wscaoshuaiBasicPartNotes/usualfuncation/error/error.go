package error

import (
	"errors"
	"fmt"
)

func Error() {
	//使用defer和recover来捕获异常处理
	defer func() {
		err := recover() //recover()是内置函数，来捕获异常的
		if err != nil {  //说明捕获到的错误(nil是err的一个0值)
			fmt.Println("发送邮件给1475001103@qq.com")
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

func ReadConf(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		return errors.New("读取文件错误...")
	}
}
