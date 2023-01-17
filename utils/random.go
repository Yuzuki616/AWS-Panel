package utils

import (
	"encoding/hex"
	"math/rand"
)

func GenRandomString(size int) string {
	tmp := make([]byte, size)
	rand.Read(tmp)
	return hex.EncodeToString(tmp)
}
