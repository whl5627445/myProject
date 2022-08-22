package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func openMySql() *gorm.DB {
	//dsn := "root:simtek_cloud_sim@tcp(127.0.0.1:3307)/simtek_cloud?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:simtek_cloud_sim@tcp(119.3.155.11:3307)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		panic(fmt.Sprintf("数据库连接失败： %s", err))
	}
	SqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Sprintf("数据库初始化失败： %s", err))
	}
	SqlDB.SetMaxIdleConns(50)
	SqlDB.SetMaxOpenConns(300)
	SqlDB.SetConnMaxLifetime(time.Hour * 5)
	Session := db.Session(&gorm.Session{PrepareStmt: true})
	return Session
}

var DB = openMySql()
