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

	// 插入测试数据
	// t := time.Now().Unix()
	// article := models.Article{Title: "linux", PostDate: t}
	// o.Insert(&article)

	sql := "select * from article"
	o.Raw(sql).QueryRows(&articles)

	this.Data["json"] = &articles
	this.ServeJson()
}
