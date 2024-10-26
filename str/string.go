package str

import "strings"

type String string

func Of(s string) String {
	return String(s)
}

func (s String) S() string {
	return string(s)
}

func (s String) Size() int {
	return Size(s.S())
}

func (s String) Equals(other string) bool {
	return Equals(s.S(), other)
}

func (s String) IsEmpty() bool {
	return IsEmpty(s.S())
}

func (s String) Bytes() []byte {
	return ToBytes(s.S())
}

func (s String) Replace(old, new string) string {
	return strings.Replace(s.S(), old, new, -1)
}
