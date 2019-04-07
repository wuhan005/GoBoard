package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

func (s *Service)register(c *gin.Context)(int, interface{}){
	type registerForm struct{
		UserName 	string
		Password 	string
		Mail 	 	string
	}

	var input = new(registerForm)
	err := c.ShouldBind(&input)

	if err != nil {
		return s.errorMsg(403, "入参错误", http.StatusForbidden)
	}

	if input.UserName == "" || input.Password == "" || input.Mail == ""{
		return s.errorMsg(403, "入参错误", http.StatusForbidden)
	}

	// 判断用户名是否重复
	rows, err := s.DB.Query("SELECT * FROM `User` WHERE `UserName` = ?", input.UserName)

	defer rows.Close()

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

// 用户登录
func (s *Service)userLogin (c *gin.Context)(int, interface{}){
	type loginForm struct {
		UserName	string
		Password	string
	}

	var input = new(loginForm)
	err := c.ShouldBindJSON(input)

	if err != nil{
		return s.errorMsg(403, "入参错误", http.StatusForbidden)
	}

	if input.UserName == "" || input.Password == "" {
		return s.errorMsg(403, "入参错误", http.StatusForbidden)
	}

	rows, err := s.DB.Query("SELECT * FROM `User` WHERE `UserName` = ? AND `Password` = ?", input.UserName, input.Password)

	defer rows.Close()

	if err != nil{
		return s.errorMsg(500, "数据库错误", http.StatusBadGateway)
	}

	if !rows.Next() {
		return s.errorMsg(403, "用户名或密码错误", http.StatusBadGateway)
	}

	var userInfo = new(UserInfo)

	err = rows.Scan(&userInfo.ID, &userInfo.UserName, &userInfo.Password, &userInfo.Mail, &userInfo.Auth, &userInfo.Token)

	if err != nil {
		return s.errorMsg(500, "数据库错误", http.StatusBadGateway)
	}

	// 生成新 Token
	token := s.generateToken(userInfo.ID)

	return s.successMsg(200, "登录成功", map[string]interface{}{
		"UserName": userInfo.UserName,
		"Mail": userInfo.Mail,
		"Avatar": "https://cdn.v2ex.com/gravatar/" + md5Encode(userInfo.Mail) + ".png",
		"Token": token,
	})
}

// 生成 Token
func (s *Service) generateToken(ID int)(token string){
	randomToken := md5Encode(fmt.Sprint(rand.Intn(int(time.Now().UnixNano()))))

	_, err := s.DB.Exec("UPDATE `User` SET `Token` = ? WHERE `ID` = ?", randomToken, ID)

	if err == nil {
		return randomToken
	}else{
		return ""
	}

}


// 检查 Token
func (s *Service) checkUserToken(token string)(userID int, Auth string){
	if token == ""{
		return -1, "user"
	}

	rows, err := s.DB.Query("SELECT * FROM `User` WHERE `Token` = ?", token)

	defer rows.Close()

	if err != nil{
		return -1, "user"
	}

	var userInfo = new(UserInfo)

	if rows.Next(){
		err = rows.Scan(&userInfo.ID, &userInfo.UserName, &userInfo.Password, &userInfo.Mail, &userInfo.Auth, &userInfo.Token)

		if err != nil{
			return -1, "user"
		}
	}else{
		return -1, "user"
	}

	return userInfo.ID, userInfo.Auth
}