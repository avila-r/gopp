package array

import (
	"github.com/avila-r/gopp"
	"github.com/avila-r/gopp/optional"
)

func Of[T any](values ...T) []T {
	return values
}

func New[T any]() []T {
	return []T{}
}

func Empty[T any]() []T {
	return New[T]()
}

func Append[T any](at []T, t ...T) []T {
	return append(at, t...)
}

func Add[T any](at *[]T, t ...T) {
	*at = append(*at, t...)
}

func Filter[S ~[]E, E any](values S, filter func(v E) bool) S {
	result := S{}

	for _, v := range values {
		if filter(v) {
			result = append(result, v)
		}
	}

	return result
}

func IsEmpty[T any](t []T) bool {
	return len(t) == 0
}

func First[T any](t []T) gopp.Optional[T] {
	return optional.Of(t[0])
}

func Last[T any](t []T) gopp.Optional[T] {
	if IsEmpty(t) {
		return optional.None[T]()
	}

	v := t[len(t)-1]

	return optional.Of(v)
}

func Size[T any](t []T) int {
	return len(t)
}

func IsNil[T any](t []T) bool {
	return t == nil
}
