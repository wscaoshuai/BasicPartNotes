package main

import (
	"fmt"
	"time"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////
//并发与并行
//1.并发：多线程程序在单核上运行，就是并发..特点：1.多个任务作用在一个CPU上；2.微观上面看，在一个时间点上面，只有一个任务在执行
//2.并行：多线程程序在多核上运行，就是并行..特点：1.多个任务在CPU在同时运行；
//3.主线程是一个物理线程，直接作用在CPU，重量级，消耗资源大
//4.协程在主线程开启的，逻辑态，轻量级，消耗资源小
//5.大部分语言，并发基于线程来并发，开启过多线程会消耗大量资源，
//6.并行与并发会有安全问题，需要解决各个协程之间的通信问题

//1.进程就是程序在操作系统的一次执行过程，是系统资源的分配和调度
//2.线程是进程的一个执行实例，是程序执行的最小单位，比进程更小的独立运行的基本单元
//3.一个进程可以创建和销毁多个进程，同一个进程中的多个线程，可以并发执行
//4.一个程序至少有一个进程，一个进程至少有一个线程
//5.线程虽然好，但是也比较吃资源，提出一个协程goroutine的概念(轻量级的线程，是编译器做得优化)
//6.协程goroutine:
////6.1:golang内非常容易实现上万个
////6.2:go协程有独立的栈空间
////6.3:共享程序堆空间
////6.4:调度由用户调度
////6.5:协程是轻量级的线程
//7.如果主线程退出了，及时协程没有结束，也会退出
//8.协程也可以在主线程退出前，结束自己

//MPG模式基本介绍：
//M:操作系统的主要线程
//P:协程执行需要的上下文资源(可以根据实际情况来开启协程)
//G:协程

//channe的使用细节
//1.channel管道可以声明为只读，或者只写(默认情况下：可读可写)
//1.1[var chan01 chan int]:channel管道声明为可读可写
//1.2[var chan02 chan<- int]:channel管道声明为只写
//1.3[var chan03 <-chan int]:channel管道声明为只读
//2.使用select可以解决从管道取数据的阻塞问题
// 2.1语法：select{
// 				case v :=<-管道：
// 				语句
// 				...
// 				default:
// 				语句
// 				}
//3.goroutine中使用recover，可以解决panic问题
//3.1如果一个协程中，出现panic，没有捕获他就会造成崩溃，使用此方法保证主线程不受影响

////////////////////////////////////////////////////////////////////////////////////////////////////////////
// var ( //声明并make一个数组
// 	myMap = make(map[int]int, 10)
// 	//声明一个全局的互斥锁
// 	//lock是一个全局的互斥锁
// 	//sync是包：synchronized同步
// 	//Mutex：是互斥锁
// 	lock sync.Mutex
// )

// func test(n int) { //声明test函数---计算 n的阶乘
// 	res := 1
// 	for i := 1; i <= n; i++ {
// 		res *= i
// 	}

// 	lock.Lock() //全局变量加锁
// 	myMap[n] = res
// 	lock.Unlock() //解锁
// }

// type Cat struct { //声明一个猫的结构体
// 	Name string
// 	Age  int
// }
////////////////////////////////////////////////////////////////////////////////////////////////////////////
func putNum(intChan chan int) { //声明一个intChan放入1-100000个数据
	for i := 1; i <= 100000; i++ {
		intChan <- i
	}
	close(intChan) //关闭intChan管道
}

//intChan取数据，判断之后存入primeChan
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool
	// var num int
	// var ok bool
	for {
		num, ok := <-intChan
		if !ok { //判断是否能在intChan中取到数据
			break
		}

		flag = true                //假设该数为素数
		for i := 2; i < num; i++ { //执行判断
			if num%i == 0 { //说明改数不是素数
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num //将取到的数放入primeChan中
		}
	}
	fmt.Println("prime协程取不到数据，退出")

	exitChan <- true //此处还不能关闭primeChan管道,所以先向exit管道内发送一个管道结束提示(后面再调用exit)

}

////////////////////////////////////////////////////////////////////////////////////////////////////////////

func main() {
	/////////////////////////////////////////////////////////////////////////////////////////////////////////
	// //案例：统计1-1000000那些是素数
	// //传统方式：
	// bool := true
	// start := time.Now().Unix()
	// for i := 2; i < 100000; i++ {
	// 	bool = true
	// 	for j := 2; j < i; j++ {
	// 		if i%j == 0 {
	// 			bool = false
	// 			break
	// 		}
	// 	}

	// 	if bool {
	// 		fmt.Printf("以下%v是素数\n", i)
	// 	}
	// }
	// end := time.Now().Unix()
	// fmt.Println("程序统计所用耗时：", end-start)

	// //go方式（引用55-93）的方法
	intChan := make(chan int, 1000)   //管道
	primeChan := make(chan int, 4000) //结果管道
	exitChan := make(chan bool, 4)    //退出管道

	start := time.Now().Unix()

	go putNum(intChan) //开启协程向intChan存入100000个数

	for i := 0; i < 4; i++ { //此处开启4个协程，再将协程判断完正确的数放
		go primeNum(intChan, primeChan, exitChan)
	}

	go func() {
		for i := 0; i < 4; i++ { //处理主线程
			<-exitChan
		}

		end := time.Now().Unix()
		fmt.Println("程序统计所用耗时：", end-start)

		close(primeChan)
	}()

	for { //遍历primeChan，把结果取出来
		res, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Printf("素数=%d\n", res)
	}
	fmt.Println("线程退出")

	/////////////////////////////////////////////////////////////////////////////////////////////////////////

	// //这个就是MPG模式中的P
	// num := runtime.NumCPU() //本地机器的逻辑CPU个数
	// runtime.GOMAXPROCS(num) //可同时执行的最大CPU个数
	// fmt.Println("num=", num)

	// for i := 0; i <= 20; i++ { //此处调用开启多个协程完成test函数
	// 	go test(i) //fatal error: concurrent map writes
	// }

	// time.Sleep(5 * time.Second) //休眠5秒钟

	// lock.Lock()
	// for i, v := range myMap { //输出结果，遍历这个结果
	// 	fmt.Printf("map[%d]=%d\n", i, v)
	// }
	// lock.Unlock()
	/////////////////////////////////////////////////////////////////////////////////////////////////////////

	//不用管道channel会有什么问题
	//上述使用全局变量加锁，并不能完美解决goroutine通讯问题
	//1.主线程在等所有goroutine全部完成时间很难确定
	//2.主线程休眠时间长了，会增加等待时长，时间短了，可能还会有goroutine处于工作状态，从而销毁
	//3.通过全局变量加锁同步来实现通讯，也并不会利用多个线程对全局变量的读写操作

	//为什么需要管道
	//1.channel本质就是一个数据结构--列队
	//2.数据就是先进先出(FIFO)
	//3.多个Goroutine访问时无需加锁，channel本身即是线程安全的
	//4.channel是有类型的，一个string的channel只能存放string类型数据

	//channel的声明：var 变量名 chan数据类型
	//示例：var intChan chan int  (intChan用于存放int数据)
	//示例：var mapChan chan map[int]string  (mapChan用于存放map[int]string类型)
	//还可以存放结构体、结构体指针等。。。
	//1.channel是引用类型
	//2.必须初始化才能使用
	//3.管道是有类型的，且只能存放指定类型的数据(可以定义空接口，就可以实现任意数据类型)
	//4.channel存放满了之后就无法继续存放
	//5.channel取出数据后，还可以继续存放
	//6.如果没有使用channel的情况下，取完数据，继续取就会报错死锁
	/////////////////////////////////////////////////////////////////////////////////////////////////////////

	// //1.创建一个可以存放3个int类型的管道
	// var intChan chan int
	// intChan = make(chan int, 3) //给管道写入数据时，容量cap不能超过3，否则会报deadlock死锁
	// //2.查看管道是什么
	// fmt.Printf("intChan 的值是=%v intChan本身的地址=%p\n", intChan, &intChan)
	// //3.向管道内写入数据
	// intChan <- 10
	// num := 211
	// intChan <- num
	// intChan <- 50
	// //4.观察管道的长度以及容量
	// fmt.Printf("channel len=%v cap=%v \n", len(intChan), cap(intChan))
	// //5.从管道内取一个数据
	// var num2 int
	// num2 = <-intChan
	// fmt.Println("num2", num2)
	// fmt.Printf("intChan 的值是=%v intChan本身的地址=%p\n", intChan, &intChan)
	// num3 := <-intChan
	// num4 := <-intChan
	// //num5 := <-intChan //如果已经取出全部内容后，再继续取数据也会报deadlock死锁
	// fmt.Println("num3 =", num3, "num4", num4)
	// //6.实现管道内存放任意类型
	// allChan := make(chan interface{}, 3)
	// allChan <- 10
	// allChan <- "Tom"
	// cat := Cat{"猫", 3}
	// allChan <- cat
	// <-allChan                                           //此处是想直接取得第三个元素，所以要先将前两个数据，直接推出管道
	// <-allChan                                           //此处是想直接取得第三个元素，所以要先将前两个数据，直接推出管道
	// newCat := <-allChan                                 //从管道中取出Cat是
	// fmt.Printf("newCat=%T,newCat=%v\n", newCat, newCat) //此处是在运行层面
	// //fmt.Printf("newCat.Name=%v", newCat.Name)//此处是在编译层面,所以编译不通过，要通过类型断言
	// a := newCat.(Cat)
	// fmt.Printf("newCat.Name=%v", a.Name)
	/////////////////////////////////////////////////////////////////////////////////////////////////////////

	//channel管道的关闭
	// intChan02 := make(chan int, 3)
	// intChan02 <- 100
	// intChan02 <- 200
	// close(intChan02) //close关闭管道
	// //intChan02 <- 300 //管道关闭后不能写入数据否则panic
	// fmt.Println("okok", intChan02)
	// n1 := <-intChan02
	// fmt.Println("n1=", n1)
	/////////////////////////////////////////////////////////////////////////////////////////////////////////

	// //channel管道的遍历写入
	// intChan03 := make(chan int, 101)
	// for i := 0; i <= 100; i++ {
	// 	intChan03 <- i * 2 //存入100个数据到管道当中
	// }

	// //channel管道的遍历读出
	// close(intChan03)
	// for v := range intChan03 { //开始遍历,要及时关闭管道才能，避免出现deadlock错误
	// 	fmt.Println("v=", v)
	// }
}
