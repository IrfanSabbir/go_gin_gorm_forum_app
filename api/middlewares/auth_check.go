package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	var errList = make(map[string]string)
	return func(c *gin.Context) {
		isTokenExist := c.Request.Header.Get("Authorization")
		if isTokenExist == "" {
			errList["Unauthorized"] = "Unauthorized"
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"error":  errList,
			})
			c.Abort()
			return
		}
		var splittedToken string
		if len(strings.Split(isTokenExist, " ")) == 2 {
			splittedToken = strings.Split(isTokenExist, " ")[1]
		}

		if splittedToken == "" {
			errList["Token_Invalid"] = "Token invalid"
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"error":  errList,
			})
			c.Abort()
			return
		}
		token, err := jwt.Parse(splittedToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Invalid authorization token %v", token.Header["alg"])
			}
			jwt_secret := []byte(os.Getenv("API_SECRET"))
			return jwt_secret, nil
		})

		if err != nil {
			errList["Token_Invalid"] = "Token invalid"
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"error":  errList,
			})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			user_id := fmt.Sprint(claims["user_id"])
			if user_id == "" {
				errList["Payload_missing"] = "Payload missing"
				c.JSON(http.StatusUnauthorized, gin.H{
					"status": http.StatusUnauthorized,
					"error":  errList,
				})
				c.Abort()
				return
			}
			c.Set("user_id", user_id)
		} else {
			errList["Token_Invalid"] = "Token invalid"
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"error":  errList,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
