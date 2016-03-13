package controllers

import (
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/orm"
	// "site/models"
	// "time"
)

type PagesController struct {
	beego.Controller
}

/**
 * 根据给定的数据获取article
 */
func (this *PagesController) Get() {
	beego.TemplateLeft = "<<<"
	beego.TemplateRight = ">>>"
	this.TplNames = "page.html"
}
