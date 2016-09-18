package bcGolang

import (
	"strings"
)

func StrEmpty(str string) bool {
	if nil == str {
		return true
	}
	if "" == strings.TrimSpace(str) {
		return true
	}
	return false
}

