package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pelletier/go-toml"
	"log"
)


type UserInfo struct {
	ID	 		int
	UserName 	string
	Password  	string
	Mail        string
	Auth		string
	Token     	string
	Avatar		string
}

func (s *Service) InitDB(){

	config, err := toml.LoadFile("config.toml")
	if err != nil {
		log.Fatal(err)
	}

	db_user := config.Get("mysql.user").(string)
	db_password := config.Get("mysql.password").(string)
	db_address := config.Get("mysql.addr").(string)
	db_name := config.Get("mysql.dbname").(string)

	DB, err := sql.Open("mysql", db_user + ":" + db_password +"@" + db_address + "/" + db_name + "?charset=utf8&parseTime=True&loc=Local")

	if err != nil{
		panic(err)
	}

	s.DB = DB
}

