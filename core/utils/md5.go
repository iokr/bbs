package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(value []byte, b ...byte) string {
	h := md5.New()
	h.Write(value)
	return hex.EncodeToString(h.Sum(b))
}
