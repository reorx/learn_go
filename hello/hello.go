package hello

import (
	"fmt"
	"os"

	"github.com/go-kit/kit/log"
)

// SayHello just say hello
func SayHello() {
	logger := log.NewJSONLogger(os.Stdout)
	logger.Log("SayHello")
}

func init() {
	fmt.Println("hello")
	logger := log.NewJSONLogger(os.Stdout)
	logger.Log("foo", 123)
}
