package auth

import (
	"testing"
)

func TestHashAndCheckPassword(t *testing.T) {
	password := "mysecret"
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("unexpected error hashing password: %v", err)
	}

	if len(hash) == 0 {
		t.Fatal("hash should not be empty")
	}

	if !CheckPasswordHash(password, hash) {
		t.Error("CheckPasswordHash should return true for correct password")
	}

	if CheckPasswordHash("wrongpass", hash) {
		t.Error("CheckPasswordHash should return false for incorrect password")
	}
}
