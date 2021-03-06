package helpers

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Adapted from https://elithrar.github.io/article/generating-secure-random-numbers-crypto-rand/

func init() {
	assertAvailablePRNG()
}

func assertAvailablePRNG() {
	// Assert that a cryptographically secure PRNG is available.
	// Panic otherwise.
	buf := make([]byte, 1)

	_, err := io.ReadFull(rand.Reader, buf)
	if err != nil {
		panic(fmt.Sprintf("encryption/rand is unavailable: Read() failed with %#v", err))
	}
}

// RandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func RandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// RandomStringSet ...
func RandomStringSet(n int, set string) (string, error) {
	bytes, err := RandomBytes(n)
	if err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = set[b%byte(len(set))]
	}
	return string(bytes), nil
}

// RandomAlphaNumeric returns a securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func RandomAlphaNumeric(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	return RandomStringSet(n, letters)
}

// RandomNumericString returns a securely generated random numeric string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func RandomNumericString(n int) (string, error) {
	const letters = "0123456789"
	return RandomStringSet(n, letters)
}

// RandomStringURLSafe returns a URL-safe, base64 encoded
// securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func RandomStringURLSafe(n int) (string, error) {
	b, err := RandomBytes(n)
	return base64.URLEncoding.EncodeToString(b), err
}

func GetEnvAndValidateBool(key string) bool {
	value := os.Getenv(key)

	if len(value) == 0 {
		panic(fmt.Sprintf("env [%s] not found", key))
	}

	b, err := strconv.ParseBool(value)

	if err != nil {
		panic(err)
	}

	return b
}

func GetEnvAndValidateInt(key string) int {
	value := os.Getenv(key)

	if len(value) == 0 {
		panic(fmt.Sprintf("env [%s] not found", key))
	}

	b, err := strconv.Atoi(value)

	if err != nil {
		panic(err)
	}

	return b
}

func GetEnvAndValidate(key string) string {
	value := os.Getenv(key)

	if len(value) == 0 {
		panic(fmt.Sprintf("env [%s] not found", key))
	}

	return value
}

func IsPresent(str *string) bool {
	if str != nil && len(*str) > 0 {
		return true
	}

	return false
}
