package models

import (
	"github.com/astaxie/beego/orm"
)

type Article struct {
	Id       int64
	Title    string
	PostDate int64
	Author   string
	Path     string
}

func init() {
	orm.RegisterModel(new(Article))
}
