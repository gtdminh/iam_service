package main

import (
	"auth_service/controller"
	"auth_service/model"
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/julienschmidt/httprouter"
)

func main() {
	server := &model.Server{Router: httprouter.New()}
	server.Router.POST("/login", controller.LoginHandler)
	server.Router.POST("/consent", controller.ConsentHandler)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), server.Router))
}
