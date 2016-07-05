package hello

import (
	"fmt"
	"os"

	"github.com/go-kit/kit/log"
)

func init() {
	fmt.Println("hello")
	logger := log.NewJSONLogger(os.Stdout)
	logger.Log("foo", 123)
}
