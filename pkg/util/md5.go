package util

import (
	"crypto/md5"
	"encoding/hex"
)

// 将文件名 MD5 后再进行写入，防止直接把原始名称就暴露出去
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}