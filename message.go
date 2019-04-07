package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (s *Service) newMessage(c *gin.Context)(int, interface{}){
	type input struct{
		Content string
	}

	message := new(input)

	token := c.GetHeader("Token")
	userID := s.checkUserToken(token)

	// 用户验权错误
	if userID == -1{
		return s.errorMsg(403, "禁止访问", http.StatusForbidden)
	}

	c.ShouldBindJSON(message)

	_, err := s.DB.Exec("INSERT INTO `Message` (`ID`, `Content`, `Create_Date`, `UserID`) VALUES (NULL, ?, ?, ?)", message.Content, time.Now().Unix(), userID)

	if err != nil{
		return s.errorMsg(500, "数据库错误", http.StatusBadGateway)
	}

	return s.successMsg(200, "留言成功", "")
}