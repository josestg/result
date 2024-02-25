package result

import (
	"fmt"
)

// Result is a generic type that represents the result of an operation that may have an error.
type Result[T any] struct {
	data T
	err  error
}

// Ok returns a new Result with the provided data.
func Ok[T any](data T) Result[T] {
	return Result[T]{data: data}
}

// Err returns a new Result with the provided error.
func Err[T any](err error) Result[T] {
	return Result[T]{data: zeroValueOf[T](), err: err}
}

func (r Result[T]) String() string {
	if r.Err() {
		return fmt.Sprintf("Err(%s)", r.err)
	}
	return fmt.Sprintf("Ok(%#v)", r.data)
}

// Ok returns true if the result is Ok.
func (r Result[T]) Ok() bool { return r.err == nil }

// Err returns true if the result is Err.
func (r Result[T]) Err() bool { return r.err != nil }

// Unwrap returns the data if the result is Ok, otherwise, it panics.
func (r Result[T]) Unwrap() T {
	if r.Err() {
		panic(r.err)
	}
	return r.data
}

// UnwrapErr returns the error if the result is Err, otherwise, it panics.
func (r Result[T]) UnwrapErr() error {
	if r.Err() {
		return r.err
	}
	panic("result: called UnwrapErr() on an Ok result")
}

// UnwrapOr returns the data if the result is Ok, otherwise, it returns the fallback value.
func (r Result[T]) UnwrapOr(fallback T) T {
	if r.Err() {
		return fallback
	}
	return r.data
}

// UnwrapOrBy returns the data if the result is Ok, otherwise, it returns the value from the supplier.
func (r Result[T]) UnwrapOrBy(fallback func(error) T) T {
	if r.Err() {
		return fallback(r.err)
	}
	return r.data
}

// Alt returns the result if it is Ok, otherwise, it returns the alternative result.
func (r Result[T]) Alt(alt Result[T]) Result[T] {
	if r.Err() {
		return alt
	}
	return r
}

// AltBy returns the result if it is Ok, otherwise, it returns the result from the supplier.
func (r Result[T]) AltBy(supplier func() Result[T]) Result[T] {
	if r.Err() {
		return supplier()
	}
	return r
}

func zeroValueOf[T any]() (z T) { return }
