package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 管理员登录
func (s *Service) checkAdmin(c *gin.Context) (int, interface{}){
	type userInput struct{
		UserName string
		Password string
	}

	u := new(userInput)

	err := c.ShouldBindJSON(u)

	if err != nil{
		return s.errorMsg(502, "入参错误", http.StatusBadGateway)
	}

	sqlStr := "SELECT * FROM `User` WHERE `UserName` = ? LIMIT 1"
	rows , err := s.DB.Query(sqlStr, u.UserName)

	if err != nil{
		panic(err)
	}

	defer rows.Close()


	arr := make([] *UserInfo, 0, 100)

	for rows.Next() {
		userInfo := new(UserInfo)

		err = rows.Scan(&userInfo.ID, &userInfo.UserName, &userInfo.Password, &userInfo.Auth, &userInfo.Token)

		if err != nil {
			panic(err)
		} else {
			arr = append(arr, userInfo)
		}
	}

	if len(arr) >= 1 {
		if u.Password == arr[0].Password{
			return s.successMsg(200, "登录成功", "")

		}else{
			// 密码错误
			return s.errorMsg(403, "用户名或密码错误", http.StatusForbidden)
		}

	}else{
		// 用户名错误
		return s.errorMsg(403, "用户名或密码错误", http.StatusForbidden)
	}

}
