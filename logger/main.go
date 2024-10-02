package logger

import (
	"log"
	"os"
)

var logger *log.Logger

func GetLogger() (*log.Logger){
	new_logger := log.New(os.Stdout, "LOG: ", log.Ldate|log.Ltime)

	return new_logger

}