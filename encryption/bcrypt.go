package encryption

import "golang.org/x/crypto/bcrypt"

// BCryptIsMatch , match hashed value with input, returns true if match
func BCryptIsMatch(input string, hashed string) (bool, error) {
	errCompare := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(input))
	if errCompare != nil {
		return false, nil
	}
	return true, nil
}

// BCrypt ...
func BCrypt(input string) (string, error) {
	hash, errHash := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
	if errHash != nil {
		return "", errHash
	}
	return string(hash), nil
}
