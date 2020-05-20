package models

import "github.com/astaxie/beego/orm"

func init()  {
	orm.RegisterModel(new(User))
}

func TNUser() string{
	return "md_user"
}
