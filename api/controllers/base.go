package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	// DBURL := fmt.Sprintf("host=%s,user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", DbHost, DbUser, DbPassword, DbName, DbPort)
	// server.DB, err = gorm.Open(postgres.New(postgres.Config{DSN: DBURL, PreferSimpleProtocol: true}), &gorm.Config{})
	server.DB, err = gorm.Open(postgres.Open(DBURL), &gorm.Config{})
	if err != nil {
		fmt.Printf("Cannot connect to %s database", Dbdriver)
		log.Fatal("error:", err)
	} else {
		fmt.Printf("Connected to %s", Dbdriver)
	}

	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 9920")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
