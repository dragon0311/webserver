package users

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/util/gvalid"
)

const (
	USER_SESSION_MARK = "user_info"
)

// 注册账户输入参数
type SignUpInput struct {
	Username  string `v:"required|length:6,16#账号不能为空|账号长度应当在:min到:max之间"`
	Email			string
	Password  string `v:"required|length:6,16#请输入确认密码|密码长度应当在:min到:max之间"`
	Password2 string `v:"required|length:6,16|same:Password#密码不能为空|密码长度应当在:min到:max之间|两次密码输入不相等"`
}

// 用户注册
func SignUp(data *SignUpInput) error {
	if  e := gvalid.CheckStruct(data, nil); e != nil {
		return errors.New(e.FirstString())
	}

}

func CheckPassport(passport string) bool {
	if i, err := user.
}