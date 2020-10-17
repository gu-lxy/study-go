package network

import (
	"io/ioutil"
	"net/http"
)


func RequestHtmlPage(page_url string)(string,error){

	client := http.Client{}//相当于浏览器

	request,err := http.NewRequest("GET",page_url,nil)
	if err != nil {
		return "",err //该函数有返回值
	}


	request.Header.Add("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	//encode：编码；decode：解码
	//request.Header.Add("Accept-Encoding","gzip, deflate, br")
	request.Header.Add("Accept-Language","zh-CN,zh;q=0.9,en;q=0.8,vi;q=0.7")
	request.Header.Add("Cache-Control","no-cache")
	request.Header.Add("Connection","keep-alive")
	request.Header.Add("Host","movie.douban.com")
	request.Header.Add("Upgrade-Insecure-Requests","1")
	request.Header.Add("User-Agent","Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36")

	response, err := client.Do(request)
	if err != nil {
		return "",err //该函数有返回值
	}
	//从response中把网页数据读取出来 []byte
	htmlBytes,err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "",err //该函数有返回值
	}
	//把byte字节切片数据，转换成为字符串格式string
	return string(htmlBytes),nil
}
