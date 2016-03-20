package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"site/models"
	// "time"
)

type ArticleController struct {
	beego.Controller
}

/**
 * 根据给定的数据获取article
 */
func (this *ArticleController) GetMoreArticle() {
	var articles []models.Article

	o := orm.NewOrm()

	sql := "select * from article"
	o.Raw(sql).QueryRows(&articles)

	this.Data["json"] = &articles
	this.ServeJson()
}
