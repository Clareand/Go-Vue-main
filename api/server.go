package api

import (
	"fmt"
	"log"
	"os"

	"github.com/Clareand/rest-api/api/controllers"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}

func Run() {
	var err error
	err = godotenv.Load()

	if err != nil {
		log.Fatalf("Faild get env, %v", err)
	} else {
		fmt.Println("env readed")
	}
	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	server.Run(":8090")
}
