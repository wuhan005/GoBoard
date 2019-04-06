package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

type Service struct{
	Router	*gin.Engine
	DB 		*sql.DB
}

func main(){
	println("Servic Started")

	s := new(Service)
	s.SetRouter()
	s.InitDB()

	s.Router.Run(":12306")
}
