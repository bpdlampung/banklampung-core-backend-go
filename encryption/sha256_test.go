package encryption

import (
	"testing"
)

func TestSHA256(t *testing.T) {
	password := "supersecretpassword"
	hPassword := SHA256(password)

	if hPassword == "" {
		t.Error("empty hashed password is not expected.")
	}
	if hPassword != "5ac152b6f8bdb8bb12959548d542cb237c4a730064bf88bbb8dd6e204912baad" {
		t.Error("different hash result")
	}
}

func TestHMACSHA256(t *testing.T) {
	password := "supersecretpassword"
	key := "mysuperkey"
	hPassword := SHA256(password)

	hmac, err := HMACSHA256(key, hPassword)

	if err != nil {
		t.Error(err)
	}

	if hmac == "" {
		t.Error("empty hashed password is not expected.")
	}
	if hmac != "418e457aca6540d35a0c6686f1783b6afe4a02f4530179b4696a5245196c651b" {
		t.Error("different hash result")
	}
}
