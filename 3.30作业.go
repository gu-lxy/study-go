package main

import (
	"fmt"
)

func main() {
	//一列数的规则如下：1、12、123、1234、12345、123456、….
	//	依次下去，求第n个数。要求：使用递归函数实现。附程序源码和截图。
	an := getShu(6)
	fmt.Println(an)
	//作业二：定义匿名函数，实现计算n的阶乘，
	//并打印计算的结果。调用匿名函数计算6的阶乘。
	s := f(6)
	fmt.Println(s)
	b := func (m int)int{
		if m ==1 {
			return 1
		}
		return (m-1)*m
	}(6)
	fmt.Println(b)
}
//1
func getShu(a int)int  {
	if a == 1 {
		return 1
	}
	return getShu(a-1)*10+a
}
//2
func f(n int)int{
	if n == 1 {
		return 1
	}
	return  f(n-1)*n
}