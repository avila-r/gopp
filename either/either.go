package either

import "fmt"

type Either[L, R any] struct {
	Left  L
	Right R
}

func Of[L, R any](l L, r R) Either[L, R] {
	return Either[L, R]{
		Left:  l,
		Right: r,
	}
}

func (e Either[L, R]) ToString() string {
	format := "Either[%T, %T](%v, %v)"

	return fmt.Sprintf(format, e.Left, e.Left, e.Right, e.Right)
}
