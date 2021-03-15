package main

import (
	"time"

	"github.com/patrickalexchan/introduction-to-go/4-functional-options/pkg/server"
)

func main() {
	srv := server.New(
		server.WithHost("localhost"),
		server.WithMaxConn(100),
		server.WithPort(8080),
		server.WithTimeout(time.Second*5),
	)
	srv.Start()
}
