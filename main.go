package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Prokop6/simple_http_server/http_handlers"
	"github.com/Prokop6/simple_http_server/logger"
)

func gracefulShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	s := <-quit
	logger.Log_info.Print("Closing server...", s)

	//server.Close()
	//server.Shutdown(context.TODO())

	os.Exit(0)
}

func main() {

	go gracefulShutdown()

	mux := http.NewServeMux()
	mux.HandleFunc("/", http_handlers.RootHandler)
	mux.HandleFunc("/echo", http_handlers.EchoHandler)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	logger.Log_info.Print("Starting server")

	go log.Fatal(server.ListenAndServe())

}
