package defer1

import (
	"fmt"
)

func Sum(n1 int, n2 int) int {
	defer fmt.Println("ok1 n1=", n1)
	defer fmt.Println("ok2 n2=", n2)

	res := n1 + n2
	fmt.Println("ok3 res=", res)
	return res

}
