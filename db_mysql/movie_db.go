package db_mysql

import (
	"NewMovieSpider/entity"
	"database/sql"
)
//导入驱动
import _ "github.com/go-sql-driver/mysql"


func OpenDatabase()(*sql.DB,error){
	database,err :=sql.Open("mysql","root:yu123456@tcp(127.0.0.1:3306)/newmovie?charset=utf8")
	if err != nil {
		return nil,err
	}
	return database,nil
}


func QueryMoviesNum() (int,error){
	//1、连接并打开数据库
	database,err := OpenDatabase()
	if err != nil {
		return 0,err
	}
	rows :=database.QueryRow("select count(movie_id) recordnum from movie")
	var recordNum int
	err =rows.Scan(&recordNum)
	if err != nil {
		return 0,err
	}
	return recordNum,nil
}


func SaveMovies2Db(movie entity.Movie)(int64,error){

	database,err := OpenDatabase()
	if err != nil {
		return 0,err
	}

	result, err := database.Exec("insert into " +
		"movie(" +
		"movie_id, movie_name, movie_desc, movie_rate, movie_vote, movie_cover_url)" +
		"values(" +
		"?, ?, ?, ?, ?, ?)",
		movie.Id, movie.Name, movie.Desc,movie.RateNum,movie.VoteNum,movie.CoverUrl)
	if err != nil {
		return 0,err
	}
	rowId,err :=result.RowsAffected()
	if err != nil {
		return 0,err
	}
	return rowId,nil
}
