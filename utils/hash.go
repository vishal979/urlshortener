package utils

import (
	"crypto/sha1"
	"encoding/base64"
)

func CalculateHash(originalurl string) string{
	hasher:=sha1.New()
	hasher.Write([]byte(originalurl))
	sha:=base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha[21:]
}