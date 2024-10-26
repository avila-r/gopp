package optional

import "github.com/avila-r/gopp"

func None[T any]() gopp.Optional[T] {
	return gopp.Optional[T]{}
}

func Of[T any](v T) gopp.Optional[T] {
	return gopp.OptionalOf(v)
}
