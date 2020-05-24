package main

import (
	_ "webserver/boot"
	_ "webserver/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
