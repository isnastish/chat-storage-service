package utils

import (
	"testing"
)

func TestSHA256(t *testing.T) {
	const testString = "saml@12349@_"
	hash, err := SHA256(testString)
	if err != nil {
		t.Fatalf("Failed to compute hash %v", err)
	}
	if len(hash) != 64 {
		t.Fatalf("Invalid hash length")
	}
	hash2, err := SHA256(testString)
	if err != nil {
		t.Fatalf("Failed to compute hash %v", err)
	}
	if hash2 != hash {
		t.Fatalf("Hashes not equal")
	}
}

func TestSHA256EmptyString(t *testing.T) {
	hash, err := SHA256("")
	if err != nil {
		t.Fatalf("Failed to compute hash %v", err)
	}
	if len(hash) != 64 {
		t.Fatalf("Invalid hash length")
	}
}
