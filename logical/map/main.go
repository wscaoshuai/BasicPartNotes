package main
import(
	"fmt"
)

//map课堂练习01
		// 使用map[string]map[string]string的map类型
		// key表示用户名，唯一，不可重复
		// 如果某用户存在，则修改其密码为“88888”，如果不存在则，增加这个用户（昵称nickname，密码pwd)
		// 编写一个函数

func modifyUser(user map[string]map[string]string,name string){
	if user[name] != nil{
		user[name]["pwd"] = "88888"
	}else{
		user[name] = make(map[string]string,2)
		user[name]["pwd"] = "88888"
		user[name]["nickname"] = "昵称" 
	}
	 
}


func main(){
	user := make(map[string]map[string]string,2)
	user["jack"] = make(map[string]string,2)
	user["jack"]["pwd"] = "99999"
	user["jack"]["nickname"] = "jack~"
	modifyUser(user,"tom")
	modifyUser(user,"mary")
	modifyUser(user,"jack")
	
	fmt.Println(user)
}