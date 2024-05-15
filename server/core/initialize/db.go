package initialize

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GormMysql() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/hyper?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db
}
