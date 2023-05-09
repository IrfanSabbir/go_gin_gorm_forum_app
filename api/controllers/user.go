package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	models "github.com/IrfanSabbir/go_gin_gorm_forum_app/api/models"
	"github.com/IrfanSabbir/go_gin_gorm_forum_app/api/utils/formaterror"
	"github.com/gin-gonic/gin"
)

func (server *Server) CreateUser(c *gin.Context) {
	errList = map[string]string{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errList["Invalid_body"] = "Unable to read body"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}
	user := models.User{}

	err = json.Unmarshal(body, &user)
	if err != nil {
		errList["Invalid_body"] = "Cannot unmarshal body"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}
	user.Prepare()

	errList = user.Validate("Create")
	if len(errList) > 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return
	}

	err = user.BeforeSave()
	if err != nil {
		errList["password_error"] = err.Error()
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
	}

	userCreated, err := user.SaveUser(server.DB)
	if err != nil {
		formatedError := formaterror.FormatError(err.Error())
		errList = formatedError
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
	}
	fmt.Println(user)
	c.JSON(http.StatusCreated, gin.H{
		"status":   http.StatusCreated,
		"response": userCreated,
	})
}
