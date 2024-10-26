package str

import (
	"fmt"
	"strings"
)

var (
	ToUpperCase = strings.ToUpper
	ToLowerCase = strings.ToLower
)

func Equals(l, r string) bool {
	return l == r
}

func ToBytes(s string) []byte {
	return []byte(s)
}

func IsEmpty(s string) bool {
	return len(s) == 0
}

func Size(s string) int {
	return len(s)
}

func ReplaceIn(a ...any) func(format string) string {
	return func(format string) string {
		return fmt.Sprintf(format, a...)
	}
}

type StringFormat string

func (f StringFormat) Values(v ...any) string {
	return fmt.Sprintf(string(f), v...)
}

func Format(format string) StringFormat {
	return StringFormat(format)
}
