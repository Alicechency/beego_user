package controllers

import (
	"beego_login/models"
	"beego_login/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"regexp"
)

type UserController struct {
	beego.Controller
}

//ShowRegister
func (this *UserController)ShowRegister(){
	this.TplName = "register.html"
}
//HandleRegister
func (this *UserController)HandleRegister()  {
	//get data
	userName := this.GetString("username")
	pwd := this.GetString("password")
	cpwd := this.GetString("cpassword")
	email := this.GetString("email")

	//valid data
	if userName==""||pwd==""||cpwd==""||email==""{
		this.Data["errormsg"] = "data can not be empty"
		this.TplName = "register.html"
		return
	}
	if pwd != cpwd{
		this.Data["errormsg"] = "password is different"
		this.TplName = "register.html"
		return
	}
	reg,_ := regexp.Compile("^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$")
	res := reg.FindString(email)
	if res==""{
		this.Data["errormsg"] = "email format is not true"
		this.TplName = "register.html"
		return
	}

	//insert into database
	o := orm.NewOrm()
	//MD5加密
	pwd = utils.MD5(pwd)
	var user models.User
	user.Username = userName
	user.Password = pwd
	user.Email = email

	_,err := o.Insert(&user)
	if err != nil{
		this.Data["exiterrormsg"] = "UserName has already exits"
		this.TplName = "register.html"
		return
	}
	//active user

	//return view
	this.Ctx.WriteString("注册成功！请登录")
	this.Redirect("/login",302)
}

//ShowLogin
func (this *UserController) ShowLogin() {
	this.TplName = "login.html"
}

//HandleLogin
func (this *UserController)HandleLogin(){
	//get data
	userName := this.GetString("username")
	pwd := this.GetString("password")
	pwd = utils.MD5(pwd)

	//valid data
	if userName==""||pwd==""{
		this.Data["errormsg"] = "data can not be empty"
		this.TplName = "register.html"
		return
	}

	o := orm.NewOrm()
	var user models.User
	user.Username = userName


	err := o.Read(&user,"Username")
	if err != nil{
		this.Data["errormsg"] = "user doesn't exit"
		this.TplName = "login.html"
		return
	}
	if user.Password != pwd{
		this.Data["errormsg"] = "password is wrong"
		this.TplName = "login.html"
		return
	}
	//return view

	this.SetSession("userName",userName)
	this.SetSession("password",pwd)
	this.Ctx.WriteString("登录成功")
}

func GetUser(this *beego.Controller)string{
	userName := this.GetSession("userName")
	if userName == nil{
		this.Data["userName"]=""
	}else{
		this.Data["userName"]= userName.(string)
	}
	return userName.(string)
}

//ShowUpdate
func (this *UserController)ShowUpdate(){
	this.TplName = "update.html"
}

//HandleUpdate
func (this *UserController)HandleUpdate(){
	//get data
	userName := this.GetSession("userName")
	username := userName.(string)
	pwd := this.GetString("password")
	//valid data

	if pwd==""{
		this.Data["errormsg"] = "data can not be empty"
		this.TplName = "update.html"
		return
	}
	o := orm.NewOrm()
	var user models.User
	user.Username = username
	err := o.Read(&user,"Username")
	if err != nil{
		this.Data["errormsg"] = "user doesn't exit"
		this.TplName = "login.html"
		return
	}
	user.Password = pwd
	o.Update(&user)

	//修改成功
	this.Redirect("/login",302)
}

