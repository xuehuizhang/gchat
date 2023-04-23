package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/url"
)

var UDB *gorm.DB

func Setup(name string, password string, host string, port int, dbName string) {
	//执行main之前 先执行init方法
	dataSourceName := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&loc=%s&parseTime=true", name, password, host, port, dbName, url.QueryEscape("Asia/Shanghai"))
	config := &gorm.Config{}
	db, err := gorm.Open(mysql.Open(dataSourceName), config)
	if err != nil {
		log.Println("连接数据库异常")
		panic(err)
	}

	////最大空闲连接数，默认不配置，是2个最大空闲连接
	//db.SetMaxIdleConns(5)
	////最大连接数，默认不配置，是不限制最大连接数
	//db.SetMaxOpenConns(100)
	//// 连接最大存活时间
	//db.SetConnMaxLifetime(time.Minute * 3)
	////空闲连接最大存活时间
	//db.SetConnMaxIdleTime(time.Minute * 1)
	//err = db.Ping()
	//if err != nil {
	//	log.Println("数据库无法连接")
	//	_ = db.Close()
	//	panic(err)
	//}
	UDB = db
}
