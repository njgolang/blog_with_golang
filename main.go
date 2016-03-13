package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// _ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	_ "site/routers"
	"site/models"
	"time"
	// "fmt"
	"os"
	"strings"
)

func init() {
	// beego.SetStaticPath("/", "views")
	beego.SetStaticPath("/articles", "articles")
	/**
	 * 数据库设置
	 * 数据库名: test
	 * user: root
	 * password: root
	 */
	orm.RegisterDriver("sqlite", orm.DR_Sqlite)
	orm.RegisterDataBase("default", "sqlite3", "data.db")
	// orm.RegisterDataBase("default", "mysql", "root:root@/test?charset=utf8", 30)
	// orm.RegisterDataBase("default", "sqlite", "root:root@/test?charset=utf8", 30)
	// orm.DefaultTimeLoc = time.UTC

	// regiseter model
	//	orm.RegisterModel(new(Article))

	// create table
	orm.RunSyncdb("default", false, false)

}

func main() {
	go syncToDb()
	beego.Run()
}


/**
 * sync md file to database
 */ 
func syncToDb() {
	ticker := time.NewTicker(time.Minute * 1)
	for _ = range ticker.C {
		articlesPath := "articles"
		mdList := getFilelist(articlesPath)
		for _, value := range mdList {
			insertArticle(value)
		}
		
	}
}


/**
 * 根据给定的数据获取article
 */
func insertArticle(filepath string) {
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		panic(err)
	}
	modTime := fileInfo.ModTime()
	filename := fileInfo.Name()
	titles := strings.Split(filename, ".")
	sz := len(titles)
	article := models.Article{Title: titles[sz-2], PostDate: modTime.Unix(), Path: filepath}
	
 
	o := orm.NewOrm()
	


	var articles []models.Article
	sql := "select * from article where Path='" + filepath + "'"
	o.Raw(sql).QueryRows(&articles)

	isExist := false
	for _, item := range articles {
		if item.Path == filepath {
			isExist = true
			break
		}
	}
	if !isExist {
		o.Insert(&article)
	}

	// fmt.Println(articles)
	//插入测试数据
	// article := models.Article{Title: "linux", PostDate: modTime.Unix(), Author: "my", Path: filepath}
	// o.Insert(&article)
}