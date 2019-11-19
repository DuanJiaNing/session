package utils

import "strings"

func BlankString(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}
