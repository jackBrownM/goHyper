package mysqlLib

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func HasDbErr(err error) bool {
	return err != nil && !errors.Is(err, gorm.ErrRecordNotFound)
}

// Exists 查询数据是否存在
// 返回值可能存在3种情况
// 1. bool,nil
// 2. false,errs
func Exists[T any](db *gorm.DB, where string, args ...any) (bool, error) {
	v := new(T)
	var num int64
	result := db.Model(v).Where(where, args...).Count(&num)
	if HasDbErr(result.Error) {
		return false, result.Error
	}
	return num > 0, nil
}

func GetOne[T any](db *gorm.DB, where string, args ...any) (*T, error) {
	v := new(T)
	result := db.Where(where, args...).First(v)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if HasDbErr(result.Error) {
		return nil, result.Error
	}
	return v, nil
}
