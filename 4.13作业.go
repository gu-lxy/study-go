package main

import (
	"fmt"
	"math"
)

func main() {
	R := r{0.3}
	fmt.Println("圆形的周长：",Circulir(&R).C())
	fmt.Println("圆形面积：",Circulir(&R).S())

	yd := yd{"王二狗","男","田径",3}
	fmt.Println(yd)

	c := cpp{"奔驰","宝马"}
	fmt.Println(c)
	cl := clx{cpp{"奔驰","宝马"},"电动","三轮"}
	fmt.Println(cl)
}
// 作业一：定义圆形结构体，
//包含有半径、周长、面积等属性；
//定义圆形结构体所拥有的两个方法：
//计算面积、计算周长。在main函数中实现调用。
type Circulir interface {
	C() float64 //周长
	S() float64 //面积
}
func GetC(circulir Circulir) float64  {
	return circulir.C()
}
func GetS(circulir Circulir) float64 {
	return circulir.S()
}
type r struct {
	r float64
}
func (c *r) C() float64 {
	return 2*c.r*math.Pi
}
func (s *r) S() float64 {
	return s.r*s.r*math.Pi
}
//作业二：定义运动员结构体，运动员有姓名、性别、
//运动项目、成绩等属性；运动员有一个run方法，
//在run方法中打印运动员所属的运动项目
//。将运动员定义在包player中，在main包中实现调用。
type yd struct {
	nm string
	x string
	xm string
	cj int
}
//作业三：使用接口定义汽车的标准：
//① 汽车品牌  ② 汽车能够驱动 并定义结构体，
//实现卡车，电动汽车、三轮车三种定义，
//实现汽车接口标准。分别实例化解放牌卡车，
//特斯拉电动车，时风三轮车，并调用对应的驱动方法。
type cpp struct {
	bc string
	bm string
}
type clx struct {
	cpp
	dd string
	sl string
}