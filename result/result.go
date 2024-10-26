package result

type R[T any] struct {
	Value T
	Err   error
}

func (r R[T]) Get() T {
	return r.Value
}

func (r R[T]) Error() error {
	return r.Err
}

func (r R[T]) HasError() bool {
	return r.Err != nil
}

func (r R[T]) HasErr() bool {
	return r.HasError()
}

func Of[T any](v T, err error) R[T] {
	return R[T]{
		Value: v,
		Err:   err,
	}
}
