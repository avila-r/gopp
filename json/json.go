package json

import "encoding/json"

func Of[T any](t T) ([]byte, error) {
	return json.Marshal(t)
}

func As[T any](b []byte) (t T, err error) {
	err = json.Unmarshal(b, &t)
	return
}
