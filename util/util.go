package util

import (
	"strconv"
)

func IsStringInt(str string) bool {
	if _, err := strconv.Atoi(str); err != nil {
		return false
	}
	return true
}
