package router

import (
	"webserver/app/api/user"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		ctlUser := new(user.Controller)
		group.ALL("/user", ctlUser)
	})
}
