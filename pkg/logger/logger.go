package logger

import (
	"log"
)

func Info(msg string) {
	log.Println(msg)
}

func Error(msg string) {
	log.Println(msg)
}