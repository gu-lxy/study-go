package main

import (
	"errors"
	"fmt"
)
func main() {
	//作业一：有结构体Person，包含name、age、address等属性。
	//定义一个函数接收结构体类型，判断传入的person的年龄是否已经成年，
	//如果已成年，返回true，如果未成年，返回false；如果年龄是0或者负数，
	//使用error返回年龄不合法等提示信息。 
	err := errors.New("")
	fmt.Print(err)
	p1 := Person{
		name:    "王权霸业",
		age:     20,
		address: "居无定所",
	}
	fmt.Println(p1)
	r1 := PersonInfo(p1)
	fmt.Println(r1)
	fmt.Println("=================")
	//作业二：定义长度为6的数组并进行赋初始值。
	//使用for循环访问0到10的数组下标上的元素。
	//处理程序可能遇到的异常，处理方式主动终止程序执行，
	//并给出异常的原因。
	a := [6]int{1,2,3,4,5,6}
	for in := 0;in <10 ;in++  {
		if in > len(a) {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(a[in])
	}
}
func PersonInfo(person Person) error {
	var err error
	if person.age >= 18 {
		err = errors.New(person.name + "已成年")
	}else if person.age <18 && person.age > 0{
		err = errors.New(person.name + "未成年")
	}else if person.age  <= 0 {
		err = errors.New(person.name + "年龄不合法")
	}
	return err
}
type Person struct {
	name string
	age int
	address string
}
