package api

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/tee-stark/go-blog/api/controllers"
)

var server = controllers.Server{}

func Run() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env: %v", err)
	} else {
		fmt.Println("Reading all environment variables...")
	}

	server.Init(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	addr := fmt.Sprintf(": %s", string(os.Getenv("PORT")))
	addr = strings.ReplaceAll(addr, " ", "")
	server.Run(addr)
}
