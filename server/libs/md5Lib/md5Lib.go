package md5Lib

import (
	"crypto/md5"
	"encoding/hex"
)

// MakeMd5 制作MD5
func MakeMd5(data string) string {
	sum := md5.Sum([]byte(data))
	return hex.EncodeToString(sum[:])
}
