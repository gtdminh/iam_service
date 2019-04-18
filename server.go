package main

import (
	"auth_service/controller"
	"auth_service/model"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	server := &model.Server{Router: mux.NewRouter().PathPrefix("/user/").Subrouter()}
	server.Router.HandleFunc("/login", controller.LoginPage).Methods("GET")
	server.Router.HandleFunc("/login/post", controller.LoginHandler).Methods("POST")
	server.Router.HandleFunc("/consent", controller.ConsentHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), csrf.Protect([]byte("32-byte-long-auth-key"))(server.Router)))
}
