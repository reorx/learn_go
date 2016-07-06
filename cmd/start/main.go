package main

import (
	"os"

	"github.com/go-kit/kit/log"
	"github.com/reorx/learngo"
)

func main() {
	logger := log.NewJSONLogger(os.Stdout)
	logger.Log("cmd main")
	learngo.Runit()
}
