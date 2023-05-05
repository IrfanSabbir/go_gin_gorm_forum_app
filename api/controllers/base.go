package controllers

import (
	"fmt"
	"log"
	"net/http"

	models "github.com/IrfanSabbir/go_gin_gorm_forum_app/api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
	server.DB, err = gorm.Open(mysql.Open(DBURL), &gorm.Config{})
	if err != nil {
		fmt.Printf("Cannot connect to %s database", Dbdriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database", Dbdriver)
	}
	server.DB.Debug().AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Comment{},
		&models.ResetPassword{},
		models.Like{},
	)

	server.Router = gin.Default()

	server.InitializeRoute()

}

func (server *Server) Run(apiPort string) {
	log.Fatal(http.ListenAndServe(apiPort, server.Router))
}
