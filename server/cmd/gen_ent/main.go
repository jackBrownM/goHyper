package main

import (
	"go/format"
	"goHyper/cmd/gen_ent/gen"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

func main() {
	tableName := "user"

	dsn := "root:root@tcp(127.0.0.1:3306)/coin?charset=utf8mb4&parseTime=True&loc=Local"
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
