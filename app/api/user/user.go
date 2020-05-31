package user

import (
	"webserver/app/service/response"
	"webserver/app/service/user"

	"github.com/gogf/gf/net/ghttp"
)

// 用户API管理对象
type Controller struct{}

// 注册请求参数， 用于前后端交互参数格式约定
type SignUpRequest struct {
	user.SignUpInput
}

func (c *Controller) SignUp(r *ghttp.Request) {
	var data *SignUpRequest
	// 这里没有使用Parse而是仅用GetStruct获取对象
	// 数据校验交给后续的service层同意处理
	if err := r.GetStruct(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := user.SignUp(&data.SignUpInput); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok")
	}
}
