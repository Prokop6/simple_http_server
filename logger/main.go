package logger

import (
	"log"
	"os"
)

var Log_info = log.New(os.Stdout, "LOG: ", log.Ldate|log.Ltime)
var Log_error = log.New(
	os.Stdout,
	"ERROR: ",
	log.Ldate|log.Ltime,
)
