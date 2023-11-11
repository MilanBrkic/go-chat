package util

import "testing"

func TestHashPassword(t *testing.T) {
	password := "milan"
	firstHash, err := HashPassword(password)

	if err != nil {
		t.Fatalf("HashPassword has failed on first hash: %v", err)
	}

	secondHash, err := HashPassword(password)

	if err != nil {
		t.Fatalf("HashPassword has failed on second hash: %v", err)
	}

	if firstHash == secondHash {
		t.Errorf("Hashing the same password two different times should result in a different hash value. Got %s and %s", firstHash, secondHash)
	}
}

func TestCompareHash(t *testing.T) {
	password := "milan"

	hashedPassword, err := HashPassword(password)

	if err != nil {
		t.Fatalf("HashPassword has failed: %v", err)
	}

	ok := CompareHash(password, hashedPassword)

	if !ok {
		t.Errorf("CompareHash should return true for same password, but got false")
	}

	ok = CompareHash("incorrect_password", hashedPassword)

	if ok {
		t.Errorf("CompareHash should return false for different passwords, but got true")
	}
}
