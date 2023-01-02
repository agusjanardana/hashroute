package hashRoute

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashRoute(data string) string {
	bytes := []byte(data)
	hash := sha256.Sum256(bytes)
	return hex.EncodeToString(hash[:])
}