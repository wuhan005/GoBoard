package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Message struct{
	ID			int
	Content 	string
	Create_Date int
	UserID 		int
	UserName   	string
	Mail        string
	Avatar 		string
}

func (s *Service) newMessage(c *gin.Context)(int, interface{}){
	type input struct{
		Content string
	}

	message := new(input)

	token := c.GetHeader("Token")
	userID := s.checkUserToken(token)

	err := c.ShouldBindJSON(message)

	if err != nil || message.Content == ""{
		return s.errorMsg(502, "入参错误", http.StatusBadGateway)
	}

	// 用户验权错误
	if userID == -1{
		return s.errorMsg(403, "禁止访问", http.StatusForbidden)
	}


	_, err = s.DB.Exec("INSERT INTO `Message` (`ID`, `Content`, `Create_Date`, `UserID`) VALUES (NULL, ?, ?, ?)", message.Content, time.Now().Unix(), userID)

	if err != nil{
		return s.errorMsg(500, "数据库错误", http.StatusBadGateway)
	}

	return s.successMsg(200, "留言成功", "")
}

func (s *Service) listMessage(c *gin.Context)(int, interface{}){
	rows, err := s.DB.Query("SELECT `Message`.*, `User`.`Mail`, `User`.`UserName` FROM `Message` INNER JOIN `User` WHERE `Message`.`UserID` = `User`.`ID`")
	if err != nil{
		return s.errorMsg(500, "数据库错误", http.StatusBadGateway)
	}

	defer rows.Close()

	messageArray := make([] *Message, 0, 100)

	for rows.Next() {
		data := new(Message)

		err = rows.Scan(&data.ID, &data.Content, &data.Create_Date, &data.UserID, &data.Mail, &data.UserName)

		if err != nil {
			panic(err)
		} else {
			data.Avatar = "https://cdn.v2ex.com/gravatar/" + md5Encode(data.Mail) + ".png"
			messageArray = append(messageArray, data)
		}
	}

	return s.successMsg(200, "成功", messageArray)
}