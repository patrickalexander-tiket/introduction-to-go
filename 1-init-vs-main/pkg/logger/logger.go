package logger

import (
	"log"
)

type Logger struct {
	Message string
}

func (l Logger) Print() string {
	return l.Message
}

func init() {
	log.Printf("init() function in logger package")
}

func main() {
	log.Println("main() function in logger package")
}
