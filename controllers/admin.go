package controllers

import (
	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

func (c *AdminController) Get() {
	beego.TemplateLeft = "<<<"
	beego.TemplateRight = ">>>"
	c.TplNames = "admin.html"
}
