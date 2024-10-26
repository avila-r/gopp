package gopp

import (
	"encoding/json"

	"github.com/avila-r/gopp/result"
)

type Json map[string]interface{}

func Marshal[T any](t T) result.R[[]byte] {
	return result.Of(json.Marshal(t))
}

func Unmarshal[T any](b []byte, r T) error {
	return json.Unmarshal(b, &r)
}
