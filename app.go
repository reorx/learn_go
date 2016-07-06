package learngo

import (
	"fmt"

	"github.com/jinzhu/now"

	"github.com/reorx/learngo/hello"
)

var bar = "app.go"

// Runit what you want
func Runit() {
	fmt.Println(bar)
	nowtime := now.BeginningOfMinute()
	fmt.Println(nowtime)

	// only `package <name>` matters, not the name of the files
	hello.Hello()
	hello.Greetings()
}
