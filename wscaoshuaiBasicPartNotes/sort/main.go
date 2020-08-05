package main
import(
	"fmt"
)
//冒泡排序
func bubblesort(arr *[5]int){
	fmt.Println(*arr)
	for	i:=0;i<len(*arr)-1;i++{
		for j:=0;j<len(*arr)-1-i;j++{
			if (*arr)[j]>(*arr)[j+1]{
				(arr)[j],(*arr)[j+1] = (*arr)[j+1],(arr)[j]
			}
		}
	}
	fmt.Println("冒泡排序",(*arr))

func main(){
	var arr = [5]int{24,69,80,57,13}

	bubblesort(&arr)
	fmt.Println("此时主函数的数组已经改变（所以传的是地址）",arr)
	}
}

