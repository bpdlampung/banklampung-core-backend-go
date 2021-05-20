package encryption

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// SHA256 ...
func SHA256(password string) string {
	sum := sha256.Sum256([]byte(password))
	return hex.EncodeToString(sum[:32])
}

// HMACSHA256 ...
func HMACSHA256(key, secret string) (string, error) {
	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte(key))
	_, errWrite := h.Write([]byte(secret))
	if errWrite != nil {
		return "", errWrite
	}

	// Get result and encode as hexadecimal string
	sha := hex.EncodeToString(h.Sum(nil))
	return sha, nil
}
