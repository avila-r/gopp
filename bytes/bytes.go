package bytes

func Empty() []byte {
	return []byte{}
}

func ToString(b []byte) string {
	return string(b)
}

func Size(b []byte) int {
	return len(b)
}
