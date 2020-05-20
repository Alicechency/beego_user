package models

type User struct {
	Id int
	Username string`orm:"size(20)"`
	Password string`orm:"size(100)"`
	Email string`orm:"size(50)"`
	Active bool`orm:"default(false)"`
}

