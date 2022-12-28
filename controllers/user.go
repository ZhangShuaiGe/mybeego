package controllers

import (
	"mybeego/dto"
	"mybeego/utils/common"

	"github.com/beego/beego/v2/adapter/logs"
	"github.com/beego/beego/v2/core/validation"
)

type User struct {
	BaseController
}

func (u *User) Register() {
	// 先引入结构体
	userInfo := &dto.UserInfo{}
	// 绑定结构体
	err := u.BindJSON(userInfo)
	if err != nil {
		logs.Info("报错", err)
	}
	// 先引入校验
	valid := validation.Validation{}
	//使用校验
	b, err := valid.Valid(userInfo)
	if !b || err != nil {
		logs.Error(1111, err, b, valid.Errors)
		u.JSON(common.ResJson{
			Code:    -1,
			Message: valid.Errors[0].Message,
		})
	}
	u.JSON(common.ResJson{
		Code:    0,
		Message: "ok",
	})
}
