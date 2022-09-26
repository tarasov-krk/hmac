package hmac

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// Sha256 returns a encoding data by hmac 256 algo
func Sha256(data, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(data))
	return hex.EncodeToString(mac.Sum(nil))
}
