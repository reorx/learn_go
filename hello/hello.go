package hello

import (
	"fmt"
	"os"

	"github.com/go-kit/kit/log"
)

var RespG = Greetings()

// Hello print hello
func Hello() {
	logger := log.NewJSONLogger(os.Stdout)
	logger.Log("SayHello")
	// Greetings()
}

func Hi() int {
	fmt.Println("well hi")
	return 2
}

// func init() {
// 	fmt.Println("hello")
// 	logger := log.NewJSONLogger(os.Stdout)
// 	logger.Log("foo", 123)
// }
