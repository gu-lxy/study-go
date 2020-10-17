package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	go funa1()
	go funa2()
	go funa3()
	go funa4()
	go funa5()

	time.Sleep(30*time.Second)
}
//在一个子协程中打印输出0-1000的随机整数，
////每次打印输出前，首先睡眠200ms，
////打印99个数字务执行完毕。；在另一个子协程中，
//////////每间隔100ms输出一个一个大写字母
//////////（范围A-Z），打印199个字幕。
//////////在main中等待两个子协程任
func funa1() {
for a := rand.Intn(1000); a <= 99; a++ {
fmt.Println(a)
time.Sleep(20*time.Millisecond)
}
}
func funa2() {
	a2 := [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "M", "n", "o", "p", "q", "r", "s", "y", "z"}
	for index := 0; index <= 199; index++ {
		fmt.Println(a2[index%6])
		time.Sleep(10 * time.Millisecond)
	}
}
//在main中启动3个子协程，执行一些耗时操作。使用sync.WaitGroup等待子协程任务完全执行结束后，整个程序执行结束。
func funa3()  {
	for a := 0; a <= 10000; a++ {//0~10000数字
		fmt.Println(a)
		time.Sleep(10*time.Millisecond)
	}
}
func funa4()  {
	for a := 0.1; a <= 999.9; a++ {//float64 0.1~999.9
		fmt.Println(a)
		time.Sleep(20*time.Millisecond)
	}
}
func funa5() {
	a2 := [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "M", "n", "o", "p", "q", "r", "s", "y", "z"}
for index := 0; index <= 199; index++ {
fmt.Println(a2[index%6])
time.Sleep(10 * time.Millisecond)
}
}