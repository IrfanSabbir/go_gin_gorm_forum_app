package api

import (
	"fmt"
	"os"

	"github.com/IrfanSabbir/go_gin_gorm_forum_app/api/controllers"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
}

func Run() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Unabe to load env")
	} else {
		fmt.Println("Env loaded")
	}
	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	apiPort := fmt.Sprintf(":%s", os.Getenv("API_PORT"))
	fmt.Printf("Listening to port %s", apiPort)
	server.Run(apiPort)

}
