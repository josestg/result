# result

A generic result type for Go to wrap a value or an error.

## Install

```shell
go get github.com/josestg/result
```

## Usage

```go
package main

import (
	"errors"
	"fmt"

	"github.com/josestg/result"
)

var anError = errors.New("a dummy error for example")

func main() {
	r1 := div(1, 0)
	// r1.Unwrap() // will panic!
	fmt.Println(r1.Err())                                                      // true
	fmt.Println(r1.Ok())                                                       // false
	fmt.Println(errors.Is(r1.UnwrapErr(), anError))                            // true
	fmt.Println(r1.UnwrapOr(42))                                               // 42
	fmt.Println(r1.UnwrapOrBy(fallbackSupplier))                               // 42
	fmt.Println(r1.Alt(result.Ok(123)))                                        // Ok(123)
	fmt.Println(r1.AltBy(func() result.Result[int] { return result.Ok(123) })) // Ok(123)

	r2 := div(4, 2)
	// r2.UnwrapErr() // will panic!
	fmt.Println(r2.Err())                                                      // false
	fmt.Println(r2.Ok())                                                       // true
	fmt.Println(r2.UnwrapOr(-1))                                               // 2
	fmt.Println(r2.UnwrapOrBy(fallbackSupplier))                               // 2
	fmt.Println(r2.Alt(result.Ok(123)))                                        // Ok(2)
	fmt.Println(r2.AltBy(func() result.Result[int] { return result.Ok(123) })) // Ok(2)
}

func fallbackSupplier(err error) int {
	if errors.Is(err, anError) {
		return 42
	}
	return -1
}

func div(a, b int) result.Result[int] {
	if b == 0 {
		return result.Err[int](anError)
	}
	return result.Ok(a / b)
}
```