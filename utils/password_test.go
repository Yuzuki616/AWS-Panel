package utils

import "testing"

func TestGenPasswordHash(t *testing.T) {
	password := "test123"
	hash, err := GenPasswordHash(password)
	if err != nil {
		t.Error(err)
	}
	t.Log(hash)
}

func TestVerifyPasswordHash(t *testing.T) {
	password := "test123"
	hash := "$2a$10$BADuTsY.rYTQvp73yut7LuinBULy8.2TJdntEBv77wZNH5FU372J6"
	if VerifyPasswordHash(password, hash) {
		t.Log("ok")
	}
}
