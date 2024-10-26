package gopp

type Optional[T any] struct {
	value *T
}

func OptionalOf[T any](t T) Optional[T] {
	return Optional[T]{
		value: &t,
	}
}

func (o *Optional[T]) IsPresent() bool {
	return o.value != nil
}

func (o *Optional[T]) Get() T {
	return *o.value
}

func (o *Optional[T]) GetOrNull() *T {
	return o.value
}
