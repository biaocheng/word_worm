package models

import (
	"github.com/xormplus/xorm"
	"github.com/astaxie/beego"
	"fmt"
	"github.com/astaxie/beego/logs"
	"os"
	"github.com/xormplus/core"
	_"github.com/go-sql-driver/mysql"
)

var DBWrite *xorm.Engine

func init() {
	dbName := beego.AppConfig.String("db.dbname")
	dbUser := beego.AppConfig.String("db.username")
	dbPassword := beego.AppConfig.String("db.password")
	dbHost := beego.AppConfig.DefaultString("db.host","127.0.0.1")
	dbPort := beego.AppConfig.DefaultInt("db.port",3306)

	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&loc=Local",dbUser,dbPassword,dbHost,dbPort,dbName)

	var err error
	DBWrite,err = xorm.NewEngine("mysql",dbUrl)

	if err!=nil{
		logs.Error(err)
		os.Exit(-1)
	}

	DBWrite.ShowSQL(true)
	DBWrite.Logger().SetLevel(core.LOG_DEBUG)


	//自动创建表
	err = DBWrite.CreateTables(new(Words),new(WordType))

	if err!=nil{
		logs.Warn(err)
	}
}