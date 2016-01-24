package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "site/routers"
)

func init() {
	/**
	 * 数据库设置
	 * 数据库名: test
	 * user: root
	 * password: root
	 */
	orm.RegisterDataBase("default", "mysql", "root:root@/test?charset=utf8", 30)

	// regiseter model
	//	orm.RegisterModel(new(Article))

	// create table
	orm.RunSyncdb("default", false, true)

}

func main() {
	beego.Run()
}
