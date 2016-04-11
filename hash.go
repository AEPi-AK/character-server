package main

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashString(str string) string {
	hash := sha256.New()
	hash.Write([]byte(str))

	return hex.EncodeToString(hash.Sum(nil))
}
