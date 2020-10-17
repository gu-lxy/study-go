package reqhandle

import "net/http"

//处理请求的函数
func Index(writer http.ResponseWriter,request *http.Request){
	writer.Write([]byte("hello every one"))
}
