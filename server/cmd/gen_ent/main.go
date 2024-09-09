package main

import (
	"go/format"
	"goHyper/cmd/gen_ent/gen"
	"goHyper/core/svc/base"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

func main() {
	tableName := "system_auth_dept"
	// 初始化配置
	config, err := base.NewConfig()
	if err != nil {
		panic(err)
	}
	dsn := config.MySQL.Conn
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，根据实际情况配置
		},
	})
	if err != nil {
		panic("failed to connect database")
	}

	// 生成模型并保存到文件
	modelCode := gen.GenerateModel(db, tableName)
	formattedCode, err := format.Source([]byte(modelCode))
	if err != nil {
		log.Fatalf("Failed to format source code: %v", err)
	}
	gen.SaveToFile(formattedCode, tableName)
}
