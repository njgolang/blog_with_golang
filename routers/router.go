package routers

import (
	"github.com/astaxie/beego"
	"site/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/about", &controllers.AboutController{})
	beego.Router("/pages", &controllers.PostsController{})
	beego.Router("/api/test", &controllers.TestController{}, "*:List")
	beego.Router("/api/articles", &controllers.ArticleController{}, "get:GetMoreArticle")
	beego.Router("/page/*.*", &controllers.PagesController{})
}
