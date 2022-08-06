package string

import "strings"

func TrimSpace(str string) string {
	return strings.Replace(str, " ", "", -1)
}
