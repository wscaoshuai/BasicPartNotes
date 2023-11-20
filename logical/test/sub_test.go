package cal

import (
	"testing"
)

func TestGetSub(t *testing.T) {

	res := GetSub(10, 3) //调用
	if res != 7 {
		t.Fatalf("GetSub(10) 执行错误，期望值=%v 实际值=%v\n", 7, res)
	}
	t.Logf("GetSub(10) 执行正确...") //如果正确，则输出日志
}
