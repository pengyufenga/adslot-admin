package model

import (
	"adslot-admin/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

var DB *gorm.DB

func init()  {
	log.Println("开始初始化数据库连接，gorm...........")
	var conn = config.Config.GetString("sqlConn")
	db, err := gorm.Open("mysql", conn)
	if err != nil {
		log.Fatal("gorm open 失败:",err)
		return
	}
	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(50)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)
	DB = db
	log.Println("数据库连接初始化成功...............")
}
