package cal

import (
	"fmt"
	"testing"
)

func TestAddUpper(t *testing.T) {

	res := addUpper(10) //调用
	if res != 55 {
		t.Fatalf("AddUpper(10) 执行错误，期望值=%v 实际值=%v\n", 55, res)
	}
	t.Logf("Addupper(10) 执行正确...") //如果正确，则输出日志
}

func TestHello(t *testing.T) {
	fmt.Println("TestHello被调用...")
}
