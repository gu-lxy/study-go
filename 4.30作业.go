package main

import (
	"fmt"
	"net/http"
)

//type mymux struct {
//
//}
////
////func (mux *mymux) ServeHTTP(writer http.ResponseWriter,request *http.Request)  {
////	p := request.URL.Path
////	fmt.Println(p)
////	if p == "/userinfo"{
////		log(writer,request)
////		return
////	}else if p == "/student" {
////		reg(writer,request)
////		return
////	}
////}
func main() {
//	//作业一:自定义一个路由处理器，并使用该路由复用器处理http请求，
//	//使用自定义的路由处理器处理请求，要求web监听的端口是8090。 
//	//①  请求路径：/userinfo   返回内容：查询用户信息 ② 
//	// 请求路径：/student   返回内容：查询学生信息  
//	mux := &mymux{}
//	err := http.ListenAndServe("127.0.0.1:8090",mux)
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//	//作业二：编写一个web服务，在浏览器中访问该服务时，
//	//显示自己硬盘上某个存在的文目录。服务的端口是9000。 
	e := http.ListenAndServe("127.0.0.1:9000",http.FileServer(http.Dir("C:/Users/86187/Music/ms")))
	if e != nil {
		fmt.Println(e.Error())
	}
}
//func log(writer http.ResponseWriter,request *http.Request)  {
//	writer.Write([]byte("查询用户信息"))
//}
//func reg(writer http.ResponseWriter,request *http.Request)  {
//	writer.Write([]byte("查询学生信息"))
//}