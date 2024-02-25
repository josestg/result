package result

import (
	"io"
	"testing"
)

type Record struct{ ID int }

func TestResult_String(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		r := Ok(42)
		expectEq(t, r.String(), "Ok(42)")

		r2 := Ok(Record{ID: 1})
		expectEq(t, r2.String(), "Ok(result.Record{ID:1})")
	})

	t.Run("err", func(t *testing.T) {
		r := Err[int](io.EOF)
		expectEq(t, r.String(), "Err(EOF)")
	})
}

func TestOk(t *testing.T) {
	r := Ok(42)
	expectEq(t, r.Ok(), true)
	expectEq(t, r.Err(), false)
}

func TestErr(t *testing.T) {
	r := Err[int](io.EOF)
	expectEq(t, r.Ok(), false)
	expectEq(t, r.Err(), true)
}

func TestUnwrap(t *testing.T) {
	r := Ok(42)
	expectEq(t, r.Unwrap(), 42)

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic, got nil")
		}
	}()
	Err[int](io.EOF).Unwrap()
}

func TestUnwrapErr(t *testing.T) {
	r := Err[int](io.EOF)
	expectEq(t, r.UnwrapErr(), io.EOF)

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic, got nil")
		}
	}()
	_ = Ok(42).UnwrapErr()
}

func TestUnwrapOr(t *testing.T) {
	r := Ok(42)
	expectEq(t, r.UnwrapOr(0), 42)
	expectEq(t, Err[int](io.EOF).UnwrapOr(0), 0)
}

func TestUnwrapOrBy(t *testing.T) {
	r := Ok(42)
	expectEq(t, r.UnwrapOrBy(func(error) int { return 0 }), 42)
	expectEq(t, Err[int](io.EOF).UnwrapOrBy(func(error) int { return 0 }), 0)
}

func TestAlt(t *testing.T) {
	r := Ok(42)
	expectEq(t, r.Alt(Err[int](io.EOF)).Unwrap(), 42)
	expectEq(t, Err[int](io.EOF).Alt(Ok(42)).Unwrap(), 42)
	expectEq(t, Err[int](io.EOF).Alt(Err[int](io.ErrUnexpectedEOF)).UnwrapErr(), io.ErrUnexpectedEOF)
}

func TestAltBy(t *testing.T) {
	r := Ok(42)
	expectEq(t, r.AltBy(func() Result[int] { return Err[int](io.EOF) }).Unwrap(), 42)
	expectEq(t, Err[int](io.EOF).AltBy(func() Result[int] { return Ok(42) }).Unwrap(), 42)
	expectEq(t, Err[int](io.EOF).AltBy(func() Result[int] { return Err[int](io.ErrUnexpectedEOF) }).UnwrapErr(), io.ErrUnexpectedEOF)
}

func expectEq(t *testing.T, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
