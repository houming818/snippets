package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// hashPassword 使用 SHA-256 对密码进行加盐和加密
func HashBySha256(password, salt string) string {
	hash := sha256.New()
	hash.Write([]byte(password + salt))
	hashInBytes := hash.Sum(nil)
	return hex.EncodeToString(hashInBytes)
}
