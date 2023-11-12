package util

import (
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	defaultValue := "difffff"
	err := os.Setenv("milan", "brkic")

	if err != nil {
		t.Fatalf("Setenv failed, can not continue with the test %v", err)
	}

	value := GetEnv("milan", defaultValue)

	if value == defaultValue {
		t.Errorf("GetEnv should find the set env, but it did not")
	}

	value = GetEnv("ne_postoji_ovaj", defaultValue)

	if value != defaultValue {
		t.Errorf("GetEnv should return the default value because wanted one was not set")
	}
}
