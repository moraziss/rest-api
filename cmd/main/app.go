package main

import (
	"github/morazss/rest-api/internal/user"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	log.Println("create router")
	router := httprouter.New()

	log.Println("register user routes")
	handler := user.NewHandler()
	handler.Register(router)

	start(router)
}

func start(router *httprouter.Router) {
	log.Println("start application")

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("server is running on port 8080")
	log.Fatal(server.Serve(listener))
}
