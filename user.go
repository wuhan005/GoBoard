package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Service)register(c *gin.Context)(int, interface{}){
	type registerForm struct{
		UserName 	string
		Password 	string
		Mail 	 	string
	}

	var input = new(registerForm)
	err := c.ShouldBindJSON(&input)

	if err != nil {
		return s.errorMsg(403, "入参错误", http.StatusForbidden)
	}

	// 判断用户名是否重复
	rows, err := s.DB.Query("SELECT * FROM `User` WHERE `UserName` = ?", input.UserName)

	if err != nil{
		return s.errorMsg(500, "数据库错误", http.StatusBadGateway)
	}

	if rows.Next(){
		return s.errorMsg(403, "用户名重复，换一个吧~", http.StatusBadGateway)
	}

	// 添加用户
	_, err = s.DB.Exec("INSERT INTO `User` (`ID`, `UserName`, `Password`, `Mail`, `Auth`, `Token`) VALUES (NULL, ?, ?, ?, 'user', '')",
		input.UserName, input.Password, input.Mail)

	if err != nil{
		return s.errorMsg(500, "数据库错误", http.StatusBadGateway)
	}

	return s.successMsg(200, "注册成功", "")

}
