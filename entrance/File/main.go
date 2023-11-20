package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// //如何读取和关闭一个文件
	// //打开一个文件(file称呼：1.file文件对象；2.file指针；3.file文件句柄)
	// file, err := os.Open("c:/test.txt")
	// if err != nil {
	// 	fmt.Println("open file err=", err)
	// }
	// //输出文件
	// fmt.Printf("file=%v", file)
	// //关闭文件
	// err = file.Close()
	// if err != nil {
	// 	fmt.Println("close file err =", err)
	// }
	//////////////////////////////////////////////////////////////////////////////////////////////////////
	// //文件读取示例01
	// file, err := os.Open("c:/test.txt")
	// if err != nil {
	// 	fmt.Println("open file err=", err)
	// }
	// defer file.Close() //凡是结束读取文件函数时，要及时关闭file避免内存泄露
	// /*
	// 		创建一个*Reader，是带缓冲的(适用于文件比较大的时候)
	// 		 const(
	// 	 		defaultBufSize = 4096//默认的缓冲区是4096
	// 	 	)
	// */
	// reader := bufio.NewReader(file)
	// //循环读取文件内容
	// for {
	// 	str, err := reader.ReadString('\n') //读到一个换行就结束读取(str是读取的内容，err是错误信息)
	// 	if err == io.EOF {                  //io.EOF表示文件的末尾
	// 		break
	// 	}
	// 	//输出内容
	// 	fmt.Print(str)
	// }
	// fmt.Println("文件读取结束。。。")
	//////////////////////////////////////////////////////////////////////////////////////////////////////
	// //文件读取示例02(写入操作)
	// file := ("c:/test.txt")
	// content, err := ioutil.ReadFile(file) //此处调用ReadFile,只适合文件较小的情况！！！！！
	// if err != nil {
	// 	fmt.Printf("read file err=%v", err)
	// }
	// //fmt.Printf("%v", content) //把读取的内容(此处读到的并非原始内容)显示到终端([]byte)
	// fmt.Printf("%v", string(content)) //这里将content类型转化成string即可显示原始内容
	// //因为此处没有Open文件，所以也不需要显示的Close文件
	// //因为文件Open和Close已经被封装到ReadFile函数内部了
	//////////////////////////////////////////////////////////////////////////////////////////////////////
	// //文件写入示例01
	// //创建一个新文件，并写入helloworld(先打开文件)
	// filePath := "c:/test/test01.txt"
	// file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666) //以写的方式打开，如果没有则创建↓
	// if err != nil {
	// 	fmt.Printf("open file err=%v\n", err)
	// 	return
	// }
	// str := "helloworld\n"
	// writer := bufio.NewWriter(file)
	// for i := 0; i < 5; i++ {
	// 	writer.WriteString(str)
	// }
	// defer file.Close() //及时关闭file句柄防止内存泄露
	// //因为writer是带缓存的，因此在调用writerstring方法时，因为内容是先写入到缓存的
	// //因此需要调用Flush方法，将缓冲的数据，真正的写入到文件中去，否则文件会无数据
	// writer.Flush()
	//////////////////////////////////////////////////////////////////////////////////////////////////////
	// //文件写入示例02(覆盖操作)
	// filePath2 := "c:/test/test01.txt"
	// file2, err2 := os.OpenFile(filePath2, os.O_WRONLY|os.O_TRUNC, 0666) //以写的方式打开，不论有无内容清空↑
	// if err2 != nil {
	// 	fmt.Printf("open file err=%v\n", err2)
	// 	return
	// }
	// str2 := "你好！北京\n"
	// writer2 := bufio.NewWriter(file2)
	// for i := 0; i < 10; i++ {
	// 	writer2.WriteString(str2)
	// }
	// defer file2.Close() //及时关闭file句柄防止内存泄露
	// //因为writer是带缓存的，因此在调用writerstring方法时，因为内容是先写入到缓存的
	// //因此需要调用Flush方法，将缓冲的数据，真正的写入到文件中去，否则文件会无数据
	// writer2.Flush()
	//////////////////////////////////////////////////////////////////////////////////////////////////////
	// //文件写入示例03(追加操作)
	// //创建一个新文件，并写入helloworld(先打开文件)
	// filePath3 := "c:/test/test01.txt"
	// file3, err3 := os.OpenFile(filePath3, os.O_WRONLY|os.O_APPEND, 0666) //以写的方式打开，在文本末尾追加
	// if err3 != nil {
	// 	fmt.Printf("open file err=%v\n", err3)
	// 	return
	// }
	// str3 := "这里是追加的内容\n"
	// writer3 := bufio.NewWriter(file3)
	// for i := 0; i < 20; i++ {
	// 	writer3.WriteString(str3)
	// }
	// defer file3.Close() //及时关闭file句柄防止内存泄露
	// //因为writer是带缓存的，因此在调用writerstring方法时，因为内容是先写入到缓存的
	// //因此需要调用Flush方法，将缓冲的数据，真正的写入到文件中去，否则文件会无数据
	// writer3.Flush()
	//////////////////////////////////////////////////////////////////////////////////////////////////////
	//文件写入示例04
	//创建一个新文件，并写入helloworld(先打开文件)
	// filePath4 := "c:/test/test01.txt"
	// file4, err4 := os.OpenFile(filePath4, os.O_RDWR|os.O_CREATE, 0666) //以读写的方式打开，且在末尾追加
	// if err4 != nil {
	// 	fmt.Printf("open file err=%v\n", err4)
	// 	return
	// }

	// reader4 := bufio.NewReader(file4) //先读取文件内容
	// for {
	// 	str, err := reader4.ReadString('\n')
	// 	if err == io.EOF { //判断是否读取到了末尾
	// 		break
	// 	}
	// 	fmt.Print(str) //显示到终端
	// }

	// str4 := "这里是显示在终端且追加内容\n"
	// writer4 := bufio.NewWriter(file4)
	// for i := 0; i < 5; i++ {
	// 	writer4.WriteString(str4)
	// }
	// defer file4.Close() //及时关闭file句柄防止内存泄露
	// //因为writer是带缓存的，因此在调用writerstring方法时，因为内容是先写入到缓存的
	// //因此需要调用Flush方法，将缓冲的数据，真正的写入到文件中去，否则文件会无数据
	// writer4.Flush()

	//////////////////////////////////////////////////////////////////////////////////////////////////////
	//文件复制示例01
	file1Path := "c:/test/test.txt" //将test文件内容导入到test01内
	file2Path := "c:/test/test01.txt"
	data, err := ioutil.ReadFile(file1Path) //先将test读到内存中
	if err != nil {
		fmt.Printf("read fil err=%v", err)
		return
	}
	err = ioutil.WriteFile(file2Path, data, 0666) //再将读到的内容写入test01中
	if err != nil {
		fmt.Printf("write file error =%v", err)
	}
	//////////////////////////////////////////////////////////////////////////////////////////////////////
	// //解析命令行参数示例01
	// fmt.Println("命令行的参数有", len(os.Args))//调用os.Args不能任意指定，不能调换顺序，不是很方便
	// for i, v := range os.Args {
	// 	fmt.Printf("args[%v]=%v\n", i, v)
	// }
	//////////////////////////////////////////////////////////////////////////////////////////////////////
	//解析命令行参数示例01
	// var user string //定义一个变量，用于接收输入的参数
	// var pwd string
	// var host string
	// var port int
	// //这里就是用来接收用户输入的-u后面的参数值; -u，就是-u的指定参数; “ ”，默认值
	// flag.StringVar(&user, "u", "", "用户名，默认为空")
	// flag.StringVar(&pwd, "pwd", "", "密码，默认为空")
	// flag.StringVar(&host, "h", "localhost", "主机名，默认为localhost")
	// flag.IntVar(&port, "port", 3306, "端口号，默认为3306")
	// flag.Parse() //转换必须调用该方法
	// fmt.Printf("use=%v pwd=%v host=%v port=%v", user, pwd, host, port)

}

// //编写一个方法判断一个文件或者文件夹是否存在
// func PathExists(path string) (bool, error) {
// 	_, err := os.Stat(path)
// 	if err == nil {
// 		return true, nil
// 	}
// 	if os.IsNotExist(err) {
// 		return false, nil
// 	}
// 	return false, err
// }
