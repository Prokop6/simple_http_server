package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/Prokop6/simple_http_server/http_handlers"
	"github.com/Prokop6/simple_http_server/logger"
)

var new_logger = logger.GetLogger()

//go:embed http_templates/*.html
var templates embed.FS


func main() {


	mux := http.NewServeMux()
	mux.HandleFunc("/", http_handlers.RootHandler)
	mux.HandleFunc("/echo", http_handlers.EchoHandler)

	new_logger.Println("Starting server")

	log.Fatal(http.ListenAndServe(":8080", mux))

}
