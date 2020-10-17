package main

import (
	"fmt"
	"strings"
)

func main() {
   //http://192.168.10.100:8080/Day33_Servlet/aa.jpeg”，
	//要求通过对字符串的处理，获取并打印出文件的名称和文件的扩展名类型。
	//附程序源码和运行截图。
	a := "http://192.168.10.100:8080/Day33_Servlet/aa.jpeg"
	a1 :=a[27:]
	a2 :=strings.SplitAfterN(a1,".",1)
	a3 := 0
	n :=make(map[string]int)
	for _,v := range a2{
		a1Slice := strings.Split(v,".")
		fmt.Println(a1Slice)
		a3++
		n[a1Slice[1]]++
	}
	a4:="jpeg"
	fmt.Println("类型:",a4)
//	作业二：给定字符串内容“hello hello hello world”，
	//	通过对字符串的处理，统计llo出现的次数。附程序源码和运行截图。  
	 b := "hello hello hello world"
	 fmt.Println(strings.Count(b,"llo"))
//	作业三：定义一个函数，该函数接收一个整数类型，作为年份，在函数中，
	//	判断接收的整数类型，是否为闰年，并打印输出该年份否是闰年。附程序源码和运行截图。
	RunNian()
}
func RunNian()  {
	var c int
	fmt.Println("请输入：")
	fmt.Scanln(&c)
	if c%4 == 0 {
		fmt.Println("是闰年")
	}else {
		fmt.Println("不是闰年")
	}
}