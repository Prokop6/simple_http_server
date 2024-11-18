package http_handlers

import (
	"net/http"
	"strings"

	"github.com/Prokop6/simple_http_server/logger"

	_ "embed"
)

func RootHandler(w http.ResponseWriter,
	r *http.Request) {

	switch strings.ToUpper(r.Method) {
	case http.MethodGet:
		getRoot(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		logger.Log_info.Printf("Received request with the method %s", r.Method)
	}
}

func FileHandler(w http.ResponseWriter,
	r *http.Request) {

	switch strings.ToUpper(r.Method) {
	case http.MethodPost:

	default:
		WrongMethodHandler(w, r, logger.Log_error)
	}

}

func EchoHandler(w http.ResponseWriter,
	r *http.Request) {

	switch r.Method {

	case http.MethodPost:

		postEcho(w, r)

	default:
		WrongMethodHandler(w, r, logger.Log_error)
	}
}
