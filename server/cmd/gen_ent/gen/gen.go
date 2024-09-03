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
		Comment    string `gorm:"column:COLUMN_COMMENT"`
	}

	db.Raw("SELECT COLUMN_NAME, DATA_TYPE, COLUMN_COMMENT FROM information_schema.columns WHERE table_schema = DATABASE() AND table_name = ?", tableName).Scan(&columns)

	var modelCode strings.Builder
	camelCaseTableName := CamelCase(tableName)
	modelCode.WriteString(fmt.Sprintf("package ent\n\n"))
	modelCode.WriteString(fmt.Sprintf("type %s struct {\n", camelCaseTableName))

	for _, col := range columns {
		goType := sqlTypeToGoType(col.DataType)
		camelCaseColName := CamelCase(col.ColumnName)
		comment := col.Comment
		if comment == "" {
			comment = "无备注"
		}
		modelCode.WriteString(fmt.Sprintf("\t%s %s `gorm:\"column:%s\"` // %s\n", camelCaseColName, goType, col.ColumnName, comment))
	}

	modelCode.WriteString("}\n\n")
	modelCode.WriteString(fmt.Sprintf("func (%s) TableName() string {\n\treturn \"%s\"\n}\n", camelCaseTableName, tableName))

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

// CamelCase 将下划线分隔的字符串转换为大驼峰命名法
func CamelCase(s string) string {
	parts := strings.Split(s, "_")
	for i, part := range parts {
		parts[i] = strings.Title(part)
	}
	return strings.Join(parts, "")
}
