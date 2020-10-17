package main

import (
	"NewMovieSpider/db_mysql"
	"NewMovieSpider/network"
	"NewMovieSpider/praseinfo"
	"NewMovieSpider/reqhandle"
	"fmt"
	"net/http"
)


func main() {


	movieNum,err := db_mysql.QueryMoviesNum()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if movieNum <=0 {//数据库中没有数据，先采集数据

		html,err :=network.RequestHtmlPage("https://movie.douban.com/chart")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		movies := praseinfo.ParseHtmlPage(html)

		//3、将采集到的页面数据保存到数据库当中
		for i :=0;i<len(movies);i++{
			_, err :=db_mysql.SaveMovies2Db(movies[i])
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}
	}

	fmt.Printf("数据库中已经存在数据%d条\n",movieNum)

	http.HandleFunc("/",reqhandle.Index)

	fmt.Println("当前网络已开启监听，可以通过http://127.0.0.1:9090进行访问")

	err =http.ListenAndServe(":9090",nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
