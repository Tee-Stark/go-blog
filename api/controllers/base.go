package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/tee-stark/go-blog/api/models"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Init(DbDriver, DbUser, DbPass, DbPort, DbHost, DbName string) {
	var err error

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPass)
	server.DB, err = gorm.Open(DbDriver, DBURL)
	server.DB.LogMode(true)
	if err != nil {
		fmt.Printf("Unable to connect to %s database", DbDriver)
		log.Fatal("Error info: ", err)
	} else {
		fmt.Printf("Connected to %s database", DbDriver)
	}
	// run database migrations automatically
	server.DB.AutoMigrate(&models.User{}, &models.Post{})

	server.Router = mux.NewRouter()

	server.InitializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening on port ", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
