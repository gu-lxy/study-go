package main

import (
	"fmt"
	"strings"
)

func main() {
	//定义字符串，内容为 “我爱你中国，I LOVE CHINA”,
	//要求输出该字符串的长度，并遍历该字符串，
	//依次输出该字符串当中的每个汉字、符号或者字母。附程序源码和运行截图。 
	a1 := "我爱你中国，I LOVE CHINA"
	fmt.Println(len(a1))
	for i:=0;i<len(a1) ;i++ {
		fmt.Println(i,"-->",a1[i])
	}
	a2 := "我爱你中国"
	a3 := []rune(a2)
	for i1:=0;i1<len(a3) ;i1++  {
		fmt.Printf("%c",a3[i1])
	}

	//作业二：定义字节切片，其元素值为97，98，99，100；
	//要求遍历该切片，并打印出切片元素值所对应的ASCII编码表中的字符。附程序源码和程序运行结果。
	fmt.Println()
    b1 :=[]rune{97,98,99,100}
	for i2:=0;i2<len(b1) ;i2++  {
		fmt.Printf("%c",b1[i2])
	}
	//作业三：定义字符串变量str，内容值为“WUHANJIAYOUZHONGGUOJIAYOU“，
	//要求判断该字符串中是否存在”ZHONGGUO“字符串，
	//如果存在，请打印输出该”ZHONGGUO“在字符串中出现的位置。附程序源码和程序运行结果。
	fmt.Println()
	c := "WUHANJIAYOUZHONGGUOJIAYOU"
	if strings.ContainsAny(c,"ZHONGGUO"){
		fmt.Println("ZHONGGUO")
	}
}
