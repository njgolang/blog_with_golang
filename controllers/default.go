package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	beego.TemplateLeft = "<<<"
	beego.TemplateRight = ">>>"
	c.TplNames = "index.html"
}

type AboutController struct {
	beego.Controller
}

func (c *AboutController) Get() {
	c.TplNames = "about.html"
}

type TestController struct {
	beego.Controller
}

func (c *TestController) List() {
	c.TplNames = "test.html"
}
