package gopp

type Condition[T any] struct {
	Cond  bool
	Value *T
}

func If[T any](c bool, t T) Condition[T] {
	var v *T
	if !c {
		v = nil
	}

	return Condition[T]{
		Cond:  c,
		Value: v,
	}
}

func (c Condition[T]) Get() *T {
	return c.Value
}

func (c Condition[T]) Else(f T) T {
	if !c.Cond {
		return f
	}

	return *c.Value
}

func (c Condition[T]) OrNull() *T {
	if !c.Cond {
		return nil
	}

	return c.Value
}

func (c Condition[T]) OrElse(f func() T) T {
	if !c.Cond {
		return f()
	}

	return *c.Value
}

func Ternary[T any](condition bool, t, f T) T {
	if condition {
		return t
	}

	return f
}
