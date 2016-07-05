package learngo

import (
	"fmt"

	"github.com/jinzhu/now"

	"github.com/reorx/learngo/hello"
)

func Run() {
	fmt.Println("app.go")
	nowtime := now.BeginningOfMinute()
	fmt.Println(nowtime)
	hello.SayHello()
}
