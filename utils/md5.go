package utils

import (
	"crypto/md5"
)

func Md5Encode(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	return string(has[:])
}
