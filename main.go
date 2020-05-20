package main

import (
	_ "beego_login/routers"
	//_"beego_login/models"
	_"beego_login/sysinit"
	_ "github.com/astaxie/beego/session/redis"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

