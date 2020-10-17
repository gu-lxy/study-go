package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//作业一：编写一个函数，该函数的作用是将数组进行翻转，
	//并返回数组。附程源码和效果运行截图。
	arr1 := fun1()
	fmt.Println("反转后的数组：",arr1)
	//	作业二：编写一个函数，该函数接收指针类型的数组，
	//	在函数中实现对数组的从小到大的排序，
	//	返回一个数组的指针。附程序源码和运行效果。
	arr2 := fun2()
	fmt.Println("排序后的数组：",arr2)
}
func fun1()*[8]int  {
	var a [8]int
	// 给数组随机赋值
	for i := 0; i < len(a); i++ {
		a[i] = rand.Intn(100)
	}
	fmt.Println("之前的数组：", a)
	length := len(a)
	for i := 0; i < length/2; i++ {
		temp := a[length-1-i]
		a[length-1-i] = a[i]
		a[i] = temp
	}
	return &a
}
func fun2()*[5]int {
	a := [5]int{23,15,8,32,9}
	fmt.Println("之前的数组；",a)
	for i := 1; i < len(a); i++{
		for index := 0; index < len(a) - i ; index++ {
			if a[index] > a[index+1]{
				a[index], a[index+1] = a[index+1], a[index]
			}
		}
	}
	return &a
}