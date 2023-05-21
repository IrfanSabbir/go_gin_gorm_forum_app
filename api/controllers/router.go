package controllers

import (
	middleware "github.com/IrfanSabbir/go_gin_gorm_forum_app/api/middlewares"
	"github.com/gin-gonic/gin"
)

func (s *Server) InitializeRoute() {
	v1 := s.Router.Group("/api/v1")
	{
		v1.GET("/ping", middleware.AuthMiddleware(), func(c *gin.Context) {
			user_id := c.GetString("user_id")
			c.JSON(200, gin.H{
				"message": "pong",
				"user_id": user_id,
			})
		})
		v1.POST("/register", s.CreateUser)
		v1.POST("/signin", s.SignIn)
	}
}
