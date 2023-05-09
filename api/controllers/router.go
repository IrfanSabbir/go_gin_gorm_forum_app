package controllers

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) InitializeRoute() {
	v1 := s.Router.Group("/api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		v1.POST("/register", s.CreateUser)
	}
}
