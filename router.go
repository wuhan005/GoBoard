package main

import (
	"github.com/gin-gonic/gin"
)

func (s *Service) SetRouter(){
	r := gin.Default()

	// 留言
	r.POST("/message", func(c *gin.Context){
		c.JSON(s.newMessage(c))
	})
	
	r.GET("/messages", func(c *gin.Context) {
		c.JSON(s.listMessage(c))
	})

	// 用户
	r.POST("/user/register", func(c *gin.Context){
		c.JSON(s.register(c))
	})

	r.POST("/user/login", func(c *gin.Context){
		c.JSON(s.userLogin(c))
	})

	// 管理员
	r.POST("/admin/login", func(c *gin.Context){
		c.JSON(s.checkAdmin(c))
	})

	r.GET("/admin/logout", func(c *gin.Context){
		c.JSON(s.adminLogout(c))
	})

	s.Router = r

}