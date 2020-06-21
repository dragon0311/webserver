package router

import (
	"webserver/app/api/user"
	"webserver/app/api/wallpaper"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		ctlUser := new(user.Controller)
		ctlWallpaper := new(wallpaper.Controller)
		group.ALL("/user", ctlUser)
		group.ALL("/wallpaper", ctlWallpaper)
	})
}
