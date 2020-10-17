package main

import (
	"fmt"
	"log"
	"net/http"
)

//作业一：编程实现一个web服务器，可以处理名为/login的请求，
//模拟登陆请求，携带用户名和密码两个数据。在服务器端接收登陆的数据，
//并对请求数据进行判断。如果用户名是hello，密码是123456，
//表示用户存在，返回“恭喜登录成功“结果字样，否则返回”对不起，
//您的账号或者密码不正确，请重新尝试“的提示语。
func main() {
	fmt.Println("服务启动中")
	http.HandleFunc("/hello_luxiaoyang", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("请求类型(方法):",request.Method)
		na := request.FormValue("na")
		fmt.Println(na)
		mm := request.Form.Get("mm")
		fmt.Println(mm)
		if na == "hello" && mm == "123456" {
			fmt.Println("恭喜登录成功")
		}else {
			fmt.Println("对不起,您的账号或者密码不正确，请重新尝试")
		}
	})
	err := http.ListenAndServe("127.0.0.1:8080",nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}