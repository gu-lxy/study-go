package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//江西软件职业技术大学区块链班5月9日作业 
	//作业一： 
	//①  编写user_info.html，该页面用于实现用户的实名信息认证。该页面上包含的信息有：
	//真实姓名，性别（单选框），出生年月，家庭住址，身份证号，学历（下拉选择框），
	//发证机关（xxx公安局），该页面有两提交按钮，点击提交按钮，可以提交数据到服务器。
	// ② 使用Go语言程序编写一个服务器程序，在本地机上运行，
	//接收user_info页面提交的数据。在服务器端解析前端提交的数据，
	//并对数据做检查，最后根据检查返回提示信息。  
	//数据规范： 
	http.HandleFunc("/user_reg",reg)
	e := http.ListenAndServe("127.0.0.1:8899",nil)
	if e != nil {
		log.Fatal(e)
	}
}
func reg(w http.ResponseWriter,request *http.Request)  {
	//姓名：最长不能超过12个字，名字不能为空，否则提示姓名不符合规范。 
	//性别：有男、女两个选项，两个选项必须选择一个，否则提示：请选择性别信息。 
	//出生年月：出生年月日不能超过2020年5月9日，否则提示：太着
	//急了，还没到日子呢。 
	//身份证号：长度必须为18位，否则提示身份证长度不符合规范。
	request.ParseForm()
	xingming := request.PostForm.Get("xingming")
	fmt.Println(xingming)
	if len(xingming) == 0 {
		fmt.Println("请重新输入")
		return
	}else if len(xingming) >= 12 {
		fmt.Println("请重新输入")
		return
	}
	xinbie  := request.PostForm.Get("xinbie")
	fmt.Println(xinbie)
	if len(xinbie) == 0 {
		fmt.Println("请重新输入")
		return
	}
	zhuzhi  := request.PostForm.Get("huzhi")
	fmt.Println(zhuzhi)
	nianyue  := request.PostForm.Get("nianyue")
	fmt.Println(nianyue)
	if nianyue >= "20200509" {
		fmt.Println("请重新输入")
		return
	}
	sfz  := request.PostForm.Get("sfz")
	fmt.Println(sfz)
	if len(sfz) < 18 && len(sfz)<18 {
		fmt.Println("请重新输入")
		return
	}
	xueli  := request.PostForm.Get("xueli")
	fmt.Println(xueli)
}