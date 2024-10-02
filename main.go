package main

import (
	"net/http"

	"github.com/Prokop6/simple_http_server/http_handlers"
	"github.com/Prokop6/simple_http_server/logger"
)

var new_logger = logger.GetLogger()


func main() {


	mux := http.NewServeMux()
	mux.HandleFunc("/", http_handlers.RootHandler)
	mux.HandleFunc("/echo", http_handlers.EchoHandler)

	new_logger.Println("Starting server")

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		new_logger.Fatalln(err.Error())

	}

}
