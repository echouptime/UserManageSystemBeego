package routers

import (
	"UserManagementSystem/controllers"
	"github.com/astaxie/beego"
)

func Register() {
	//默认路由
	beego.Router("/", &controllers.BaseController{}, "*:BaseInfo")
	//用户路由
	beego.AutoRouter(&controllers.UserController{})
}
