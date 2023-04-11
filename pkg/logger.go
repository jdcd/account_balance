package pkg

import (
	"log"
	"os"
)

const (
	reset  = "\033[0m"
	yellow = "\033[33m"
	red    = "\033[31m"
	green  = "\033[32m"
)

func WarningLogger() *log.Logger {
	return log.New(os.Stdout, yellow+"[WARNING] "+reset, log.Ldate|log.Ltime)
}

func InfoLogger() *log.Logger {
	return log.New(os.Stdout, green+"[INFO] "+reset, log.Ldate|log.Ltime)
}

func ErrorLogger() *log.Logger {
	return log.New(os.Stdout, red+"[ERROR] "+reset, log.Ldate|log.Ltime)
}
