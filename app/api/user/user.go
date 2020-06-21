package user

import (
	"webserver/app/service/response"
	"webserver/app/service/user"

	"github.com/gogf/gf/net/ghttp"
)

// 用户API管理对象
type Controller struct{}

/********************用户注册********************/

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

/********************用户登录********************/

//登录请求参数
type SignInRequest struct {
	Username string `v:"required#用户名不能为空"`
	Password string `v:"required#密码不能为空"`
}

func (c *Controller) SignIn(r *ghttp.Request) {
	var data *SignInRequest
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := user.SignIn(data.Username, data.Password, r.Session); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok")
	}
}

// 检查用户是否已经登录
func (c *Controller) IsSignedIn(r *ghttp.Request) {
	var data *SignInRequest
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if isSignedIn := user.IsSignedIn(data.Username, r.Session); isSignedIn == true {
		response.JsonExit(r, 1, "该用户已登录!")
	} else {
		response.JsonExit(r, 0, "该用户未登录!")
	}
}
