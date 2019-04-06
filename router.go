package main

import (
	"github.com/gin-gonic/gin"
)

func (s *Service) SetRouter(){
	r := gin.Default()

	r.POST("/admin/login", func(c *gin.Context){
		c.JSON(s.checkAdmin(c))
	})

	s.Router = r

}