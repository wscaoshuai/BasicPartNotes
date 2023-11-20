package suffix

import "strings"

func Makesuffix(suffix string) func(string) string {

	return func(name string) string {
		if !strings.HasSuffix(name, suffix) { //如果没有指定的后缀名，则加上，否则就返回原来的名字
			return name + suffix
		}
		return name
	}
}
