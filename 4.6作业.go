package main

import "fmt"

func main() {
	p1 := XinXi{9527,
		"王权霸业",
		"乞丐",
		400,
		 0,
	}
	fmt.Println(p1)
	a1 := fun(p1.salary)
	fmt.Println(a1)
	p2 := XinXi{
		number:     9528,
		name:       "王权富贵",
		department: "乞丐",
		salary:     500,
		teaxes:     0,
	}
	fmt.Println(p2)
	fun(p2.salary)
	a2 := fun(p2.salary)
	fmt.Println(a2)
}
//作业一：公司里的职员，可以描述职员的信息包括：员工编号，
//姓名，部门，薪水，交税数额。要求：编程创建结构体。
//并创建几个具体的职员信息，交税数额均设置为默认值0，然后打印输出这些职员的信息。   
type XinXi struct {
	number int
	name string
	department string
	salary int
	teaxes int
}
//作业二：在作业一的基础上，创建一个函数，该函数接收职员类型，
//在函数中判断职员的薪水是否大于5000，薪水小于等于5000元，
//交税数额为0； 如果大于5000，则计算员工要交的税额，并赋值到相应的字段。
//最后在main函数中调用函数，并接受返回值，打印员工要交的税额。  
//注：员工应交税税额 = （薪水 – 5000） * 20 %
func fun(salary int) int {
	var sum  int
	if salary > 5000 {
		sum = (salary-5000)*(20/100)
	}
	return sum
}