package testabilities

import (
	"crypto/sha256"
	"encoding/hex"
)

func hexHash(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}
