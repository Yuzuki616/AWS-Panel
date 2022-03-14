package utils

import "os"

func IsNotFound(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return true
	} else {
		return false
	}

}
