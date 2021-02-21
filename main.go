package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/rohitkharanghar/go-microservice/handlers"
)

func main() {

	logger := log.New(os.Stdout, "Go-Microservice", log.LstdFlags)

	hh := handlers.NewHello(logger)
	gh := handlers.NewGoodBye(logger)

	serverMutex := http.NewServeMux()
	serverMutex.Handle("/", hh)
	serverMutex.Handle("/goodBye", gh)
	server := http.Server{
		Addr:         ":9090",
		Handler:      serverMutex,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  10 * time.Second,
	}
	go func() {
		logger.Println("Starting Server at port 9090")

		err := server.ListenAndServe()
		if err != nil {
			logger.Printf("Error while starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt)
	signal.Notify(channel, os.Kill)

	sig := <-channel
	logger.Println("Got signal", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(ctx)
}
