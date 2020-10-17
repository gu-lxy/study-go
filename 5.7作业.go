package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {
	//江西软件职业技术大学区块链班5月7日作业 
	//作业一： 
	//①  编写一个web服务器程序，在自己电脑上的9001端口进行监听，
	//可以处理/userlogin请求,返回一句话“欢迎访问用户登录功能” 
	//②  编程实现一个客户端程序，请求步骤①中的服务器，
	//并访问/userlogin，在客户端程序中接收请求，
	//并打印出请求到的信息。 程序编写好后，
	//先运行服务器程序，然后运行客户端程序，
	//查看程序运行效果。 
	//nux := http.NewServeMux()
	//nux.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
	//	io.WriteString(writer,"用户登录")
	//})
	//nux.HandleFunc("/userlogin", func(writer http.ResponseWriter, request *http.Request) {
	//	io.WriteString(writer,"欢迎访问用户登录功能")
	//})
	//e := http.ListenAndServe("127.0.0.1:9001",nux)
	//if e != nil {
	//	fmt.Println(e.Error())
	//}

	client := http.Client{}
	urlStr := "http://www.baidu.com"
	params := &url.Values{
		"theCityName":{"欢迎访问用户登录功能"},
	}
	bufferString:= bytes.NewBufferString(params.Encode())
	response ,err :=client.Post(urlStr,"application/x-www-form-urlencoded",bufferString)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer response.Body.Close()
	fmt.Println(response.StatusCode)
}