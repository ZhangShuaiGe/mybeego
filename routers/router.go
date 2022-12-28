package routers

import (
	"mybeego/controllers"

	web "github.com/beego/beego/v2/server/web"
)

// 登录  获取用户信息  添加用户  删除用户
func init() {
	ns := web.NewNamespace("/v1",
		web.NSCtrlPost("/userRegister", (*controllers.User).Register),
	)
	//注册 namespace
	web.AddNamespace(ns)
	// web.Router("/", &controllers.MainController{})
	// web.CtrlPost("/orderList", (*user.LoginController).OrderList)
	// web.CtrlPost("/createdOrder", (*user.LoginController).CreatedOrder)
}
