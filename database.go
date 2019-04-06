package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


type UserInfo struct {
	ID	 		int
	UserName 	string
	Password  	string
	Auth		string
	Token     	string
}

func (s *Service) InitDB(){
	DB, err := sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/GoBoard?charset=utf8&parseTime=True&loc=Local")

	if err != nil{
		panic(err)
	}

	s.DB = DB
}

