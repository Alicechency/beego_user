package routers

import (
	"beego_login/controllers"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego"
)

func init() {
	beego.InsertFilter("/user/*",beego.BeforeExec,filterFunc)
    beego.Router("/", &controllers.MainController{})
    beego.Router("/register",&controllers.UserController{},"get:ShowRegister;post:HandleRegister")
    beego.Router("/login",&controllers.UserController{},"get:ShowLogin;post:HandleLogin")
	beego.Router("/user/update",&controllers.UserController{},"get:ShowUpdate;post:HandleUpdate")
}

var filterFunc = func(ctx *context.Context) {
	userName := ctx.Input.Session("userName")
	if userName==nil{
		ctx.Redirect(302,"/login")
		return
	}
}