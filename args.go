package gopp

import "log"

func First[T any](v T, _ ...any) T {
	return v
}

func Second[T any](_ any, v T, _ ...any) T {
	return v
}

func Third[T any](_, _ any, v T, _ ...any) T {
	return v
}

func Coalesce[T comparable](values ...T) (v T) {
	var zero T

	for _, v = range values {
		if v != zero {
			return
		}
	}

	return
}

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}

	return v
}

func Catch[T any](v T, err error) T {
	if err != nil {
		log.Printf("error in catch block - %v", err.Error())
	}

	return v
}
