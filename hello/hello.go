package hello

import (
	"fmt"
	"os"

	"github.com/go-kit/kit/log"
)

func SayHello() {
	fmt.Println("SayHello")
}

func init() {
	fmt.Println("hello")
	logger := log.NewJSONLogger(os.Stdout)
	logger.Log("foo", 123)
}
