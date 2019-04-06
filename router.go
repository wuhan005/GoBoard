package main

import (
	"github.com/gin-gonic/gin"
)

func (s *Service) SetRouter(){
	r := gin.Default()

	// 用户
	r.POST("/user/register", func(c *gin.Context){
		c.JSON(s.register(c))
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