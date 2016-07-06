package hello

import (
	"fmt"
	"os"

	"github.com/go-kit/kit/log"
)

// SayHello print hello
func Hello() {
	logger := log.NewJSONLogger(os.Stdout)
	logger.Log("SayHello")
}

func init() {
	fmt.Println("hello")
	logger := log.NewJSONLogger(os.Stdout)
	logger.Log("foo", 123)
}
