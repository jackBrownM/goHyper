package gen

import (
	"fmt"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func GenerateModel(db *gorm.DB, tableName string) string {
	var columns []struct {
		ColumnName string `gorm:"column:COLUMN_NAME"`
		DataType   string `gorm:"column:DATA_TYPE"`
	}

	db.Raw("SELECT COLUMN_NAME, DATA_TYPE FROM information_schema.columns WHERE table_schema = DATABASE() AND table_name = ?", tableName).Scan(&columns)

	var modelCode strings.Builder
	modelCode.WriteString(fmt.Sprintf("package ent\n\n"))
	modelCode.WriteString(fmt.Sprintf("type %s struct {\n", strings.Title(tableName)))

	for _, col := range columns {
		goType := sqlTypeToGoType(col.DataType)
		modelCode.WriteString(fmt.Sprintf("\t%s %s `gorm:\"column:%s\"`\n", strings.Title(col.ColumnName), goType, col.ColumnName))
	}

	modelCode.WriteString("}\n\n")
	modelCode.WriteString(fmt.Sprintf("func (%s) TableName() string {\n\treturn \"%s\"\n}\n", strings.Title(tableName), tableName))

	return modelCode.String()
}

func SaveToFile(code []byte, tableName string) {
	fileName := fmt.Sprintf("./internal/ent/%s.go", tableName)
	err := os.MkdirAll("./internal/ent", os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create directory: %v", err)
	}

	err = ioutil.WriteFile(fileName, code, 0644)
	if err != nil {
		log.Fatalf("Failed to write model to file: %v", err)
	}

	fmt.Printf("Model saved to %s\n", fileName)
}

func sqlTypeToGoType(sqlType string) string {
	switch sqlType {
	case "varchar", "text", "longtext":
		return "string"
	case "int", "smallint", "tinyint":
		return "int"
	case "bigint":
		return "int64"
	case "float", "double":
		return "float64"
	case "datetime", "date", "timestamp":
		return "time.Time"
	default:
		return "interface{}" // 未知类型
	}
}
