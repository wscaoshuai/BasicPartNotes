package main

//////////////////////////////////////////////////////////////////////////////////////////////////////////////
// //编写一个函数(每隔一秒种输出helloWorld)
// func test() {
// 	for i := 1; i <= 10; i++ {
// 		fmt.Println("test():Hello,World" + strconv.Itoa(i))
// 		time.Sleep(time.Second)
// 	}
// }
// func main() {

// 	//test()    //此处调用test函数
// 	go test() //此处就开启了一个协程

// 	for i := 1; i <= 10; i++ {
// 		fmt.Println("main():Hello,Golang" + strconv.Itoa(i))
// 		time.Sleep(time.Second)
// 	}
// }

//////////////////////////////////////////////////////////////////////////////////////////////////////////////

//channel管道的基本练习01
//创建一个Person结构体，使用rand随机生成10个Person实例，放入channel，遍历channel，逐个打印Person实例的信息在终端
type Person struct {
	Name   string
	Age    int
	Adress string
}

func main() {

}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////
//channel管道与goroutine基本练习01
//开启一个writeData协程，向管道intChan中读取50个整数
//开启一个readData协程，从管道intChan中读取writeData写入数据
//writeData和readData操作的是同一个管道，主线程需要等待前两个协程完成，才能退出(管道解决)

// func writeData(intChan03 chan int) {
// 	for i := 1; i <= 50; i++ {
// 		intChan03 <- i //存入50个数据到管道当中
// 		fmt.Printf("writeData 写入数据=%v\n", i)
// 		//time.Sleep(time.Second * 1)//可以关掉
// 	}
// 	close(intChan03) //及时关闭才不影响后面读操作
// }

// func readData(intChan03 chan int, exitChan04 chan bool) {
// 	for {
// 		v, ok := <-intChan03
// 		if !ok {
// 			break
// 		}
// 		fmt.Printf("readData 读到数据=%v\n", v)
// 	}
// 	exitChan04 <- true
// 	close(exitChan04)
// }
// func main() {
// 	intChan03 := make(chan int, 50)
// 	exitChan04 := make(chan bool, 1)
// 	go writeData(intChan03)
// 	go readData(intChan03, exitChan04)

// 	//time.Sleep(time.Second*10)//用↓方法解决不用休眠也可使用
// 	//在主函数内，下述for循环不需要，只留一个<-exitChan就可以起到阻塞效果
// 	for {
// 		_, ok := <-exitChan04
// 		if !ok {
// 			break
// 		}
// 	}
// }

//////////////////////////////////////////////////////////////////////////////////////////////////////////////
//channel管道与goroutine基本练习02
//启动一个协程，将1-2000的数放入到一个channel中(例如numChan)
//启动8个协程，从numChan取出数(比如n)，计算1+n...n的值，并存放到ResChan
//8个协程完成工作后，再遍历resChan，显示结果:res[1]=1...res[10]=55

//////////////////////////////////////////////////////////////////////////////////////////////////////////////
//channel管道与goroutine基本练习03
//开启一个协程WriteDataToFile，随机生成1000个数据，存放到文件中
//当WriteDataToFile完成写入数据后，就让sort协程从文件读取1000个文件，并完成排序，重新写入到另一个协程内
//扩展：10个协程WriteDataToFile，每个协程随机生成1000个数据，存放到10个文件内，
//扩展：当10个文件都生成以后，让10个sort协程从10个文件中读取1000个文件，排序后重新写入
