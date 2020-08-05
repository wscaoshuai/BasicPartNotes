package main
import(
	"fmt"
)

//39行二分查找函数
func BinaryFind(arr *[6]int,leftIndex int,rightIndex int,findval int){

	if leftIndex>rightIndex{
		fmt.Println("找不到")
		return
	}
	middle := (rightIndex + leftIndex)/2//此处找到中间下标
	if (*arr)[middle] > findval {
		BinaryFind(arr,leftIndex,middle-1,findval)
	}else if(*arr)[middle]<findval{
		BinaryFind(arr,middle+1,rightIndex,findval)
	}else{
		fmt.Printf("找到了，下标为：%v\n",middle)
	}

}

func main(){

	// //顺序查找查找方式1：
	// var names = [4]string{"Jack","Tom","Mary","Lisa"}
	// var heroname = ""
	// fmt.Scanln(&heroname)
	// for i:=0;i<len(names);i++{
	// 	if heroname == names[i]{
	// 		fmt.Printf("查找到了%v,下标%v",heroname,i)
	// 		break
	// 	}else if i==(len(names)-1){
	// 		fmt.Printf("没有找到%v",heroname)
	// 	}
	// }

	// //顺序查找查找方式2：（推荐使用）
	// var names = [4]string{"Jack","Tom","Mary","Lisa"}
	// var heroname = " "
	// index := -1
	// fmt.Scanln(&heroname)
	// for i:=0;i<len(names);i++{
	// 	if heroname == names[i]{
	// 		index = i
	// 		break
	// 	}
	// }
	// if index != -1{
	// 	fmt.Printf("查找到了%v,下标%v",heroname,index)
	// }else{
	// 	fmt.Printf("没有找到%v",heroname)
	// }

	//二分查找--从两端查找（效率高，有前提要有序）
	arr := [6]int{1,8,10,89,1000,1024}
	BinaryFind(&arr,0,len(arr)-1,1000)


}