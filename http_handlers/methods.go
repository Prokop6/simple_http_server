package http_handlers

import (
	"io"
	"log"
	"net/http"

	"github.com/Prokop6/simple_http_server/logger"
)



func getRoot(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Server works!\n"))
}

func postEcho(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)

	if err != nil {

		panic(err.Error())
	}

	w.Write(body)

}


func WrongMethodHandler(w http.ResponseWriter, r *http.Request, log *log.Logger) {
	w.WriteHeader(http.StatusMethodNotAllowed)

	new_logger := logger.GetLogger()

	new_logger.Printf("Received request with the method %s", r.Method)

}