package utilLib

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/google/uuid"
	"math/rand"
	"strings"
)

var (
	allRandomStr = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func RandomString(length int) string {
	byteList := make([]byte, length)
	for i := 0; i < length; i++ {
		byteList[i] = allRandomStr[rand.Intn(62)]
	}
	return string(byteList)
}

// MakeMd5 制作MD5
func MakeMd5(data string) string {
	sum := md5.Sum([]byte(data))
	return hex.EncodeToString(sum[:])
}

// MakeUuid 制作UUID
func MakeUuid() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
