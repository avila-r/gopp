package result

type R[T any] struct {
	Value T
	Err   error
}

func (r R[T]) Get() T {
	return r.Value
}

func (r R[T]) Error() string {
	return r.Err.Error()
}

func (r R[T]) HasError() bool {
	return r.Err != nil
}

func (r R[T]) HasErr() bool {
	return r.HasError()
}

func (r R[T]) IsSuccess() bool {
	return !r.HasError()
}

func (r R[T]) Unwrap() T {
	if r.HasError() {
		panic(r.Error())
	}

	return r.Value
}

func (r R[T]) Expect(msg string) T {
	if r.HasError() {
		panic(msg)
	}

	return r.Value
}

func Of[T any](v T, err error) R[T] {
	return R[T]{
		Value: v,
		Err:   err,
	}
}
