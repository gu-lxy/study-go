package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	//作业一：定义函数，接收两个参数，a和b。在函数中，
	//以a和b为矩形的长和宽，计算出矩形的周长，
	//并返回给调用处。附程序源码和运行调用效果截图。 
	JuXing(4,5)
	//作业二：定义函数，接收一个字符串，
	//在函数中统计出字符串中各个字符出现的次数，
	//并打印输出，在main函数中实现调用。附程序运行截图。   
	str("hbfvuhdjfohibuobnofbojbvfbojbobgbhbgkb")
	//作业三：定义一个函数，接收一个float类型数据，
	//作为圆形的半径，计算得到圆形的面积并返回，
	//在main中调用函数，附程序和运行效果截图。
	MianJi(3.3)
}
func JuXing(n int,m int)  {
	sum := 0
	sum = (n+m)*2
	fmt.Println(sum)
}
func str(s string)  {
	fmt.Println(strings.Count(s,"h"))
}
func MianJi(r float64)  {
	S := math.Pi*r*r
	fmt.Println(S)
}