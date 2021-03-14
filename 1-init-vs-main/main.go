package main

import (
	"log"

	"github.com/patrickalexchan/introduction-to-go/1-init-vs-main/example"
)

func init() {
	log.Println("init() function in main() - 3")
}

func init() {
	log.Println("init() function in main() - 1")
}

func init() {
	log.Println("init() function in main() - 2")
}

func main() {
	e := example.Test{Message: "Hello from Example!"}
	log.Println(e.Echo())
	log.Println("main func()")
}
