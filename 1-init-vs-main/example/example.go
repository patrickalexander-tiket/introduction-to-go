package example

import (
	"log"

	"github.com/patrickalexchan/introduction-to-go/1-init-vs-main/pkg/logger"
)

type Test struct {
	Message string
}

func (t Test) Echo() string {
	return t.Message
}

func init() {
	log.Printf("init() function in example package")
}

func main() {
	l := logger.Logger{Message: "Hello from Logger!"}
	log.Println(l.Print())
	log.Println("main() function in example package")
}
