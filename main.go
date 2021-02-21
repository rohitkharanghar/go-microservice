package main

import (
	"log"
	"net/http"
	"os"

	"github.com/rohitkharanghar/go-microservice/handlers"
)

func main() {

	logger := log.New(os.Stdout, "Go-Microservice", log.LstdFlags)

	hh := handlers.NewHello(logger)
	gh := handlers.NewGoodBye(logger)

	serverMutex := http.NewServeMux()
	serverMutex.Handle("/", hh)
	serverMutex.Handle("/goodBye", gh)
	http.ListenAndServe(":9090", serverMutex)
}
