package main

import (
	"fmt"
	"math/rand"
	"time"
)

//作业一：启动一个goroutine，生成100个随机数发送到容量为5的缓冲通道ch1中，
//随机数间隔20ms生成一个；启动另外一个goroutine，从ch1中读取数据，
//然后计算读取到的数值的3次方，放到通道ch2中；在main函数中，
//从通道ch2中读取数据，并打印输出。最后程序结束。
var ch1 chan int = make(chan int)
var ch2 chan int = make(chan int)
func main() {
	ch1 = make(chan int,5)
	ch2 = make(chan int,5)
	go fun11(ch1)//产生
	go fun12(ch2)//输出
	n := <- ch2
	fmt.Println(n)
	time.Sleep(1*time.Second)
}
func fun11(ch chan int)  {
	//设置随机数种子
	rand.Seed(time.Now().UnixNano())
	for i := 0; i <= 100; i++ {
		//产生0到100的整数
		num := rand.Intn(100)
		fmt.Println(num)
		ch1 <- num
		<-ch
	}
}
func fun12(ch chan int)  {
	for {
		var num int
		num = <-ch1
		ch2 <- num*num*num
		<-ch
	}
}