package user

import (
	"errors"
	"fmt"
	"webserver/app/model/user"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
)

const (
	USER_SESSION_MARK = "user_info"
)

/********************用户注册********************/

// 注册账户输入参数
type SignUpInput struct {
	Username  string `v:"required|length:6,16#账号不能为空|账号长度应当在:min到:max之间"`
	Email     string
	Password  string `v:"required|length:6,16#请输入确认密码|密码长度应当在:min到:max之间"`
	Password2 string `v:"required|length:6,16|same:Password#密码不能为空|密码长度应当在:min到:max之间|两次密码输入不相等"`
}

// 用户注册
func SignUp(data *SignUpInput) error {
	if e := gvalid.CheckStruct(data, nil); e != nil {
		return errors.New(e.FirstString())
	}

	// 用户名唯一性检查
	if !CheckUsername(data.Username) {
		return errors.New(fmt.Sprintf("用户名 %s 已经存在", data.Username))
	}

	// Email唯一性检查
	if !CheckUsername(data.Username) {
		return errors.New(fmt.Sprintf("用户名 %s 已经存在", data.Username))
	}

	// 将输入参数赋值到数据库实体对象上
	var entity *user.Entity
	if err := gconv.Struct(data, &entity); err != nil {
		return err
	}

	// 记录账号创建、注册时间
	entity.CreateTime = gtime.Now()
	if _, err := user.Save(entity); err != nil {
		return err
	}
	return nil
}

// 检查用户名唯一性
func CheckUsername(username string) bool {
	if i, err := user.FindCount("username", username); err != nil {
		return false
	} else {
		return i == 0
	}
}

// 检查Email唯一性
func CheckEmail(email string) bool {
	if i, err := user.FindCount("email", email); err != nil {
		return false
	} else {
		return i == 0
	}
}

/********************用户登录********************/

// 用户登录
func SignIn(username, password string, session *ghttp.Session) error {
	one, err := user.FindOne("username=? and password=?", username, password)
	if err != nil {
		return err
	}
	if one == nil {
		return errors.New("用户名密码错误")
	}
	return session.Set(username, one)
}

// 用户是否已经登录
func IsSignedIn(username string, session *ghttp.Session) bool {
	return session.Contains(username)
}
