package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	web.Controller
}

func (b *BaseController) JSON(data interface{}) {
	b.Data["json"] = data
	b.ServeJSON()
}
