package fileLib

import (
	"os"
	"path/filepath"
)

// Exists 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// IsDir 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// IsFile 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}

// CreateDirIfNotExist 目录如果不存在则创建
func CreateDirIfNotExist(path string) error {
	if !Exists(path) {
		if err := os.MkdirAll(path, 0755); err != nil {
			return err
		}
	}
	return nil
}

// GetWorkingDir 获取当前程序的工作目录
func GetWorkingDir() string {
	// 获取当前工作目录
	dir, _ := os.Getwd()
	return dir
}

var WebDirName = "wwwroot"

// GetWebDir 获取当前程序的工作目录
func GetWebDir() string {
	return filepath.Join(GetWorkingDir(), WebDirName)
}
