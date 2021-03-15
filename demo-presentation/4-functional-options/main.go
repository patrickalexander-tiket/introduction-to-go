package main

import (
	"time"

	"github.com/patrickalexchan/introduction-to-go/4-functional-options/pkg/server"
)

//START1 OMIT
func main() {
	srv := server.New(
		server.WithHost("localhost"),      // H12
		server.WithMaxConn(100),           // H12
		server.WithPort(8080),             // H12
		server.WithTimeout(time.Second*5), // H12
	)
	srv.Start()
}

//END1 OMIT
