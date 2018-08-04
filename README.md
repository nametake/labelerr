# labelerr

[![CircleCI](https://circleci.com/gh/nametake/labelerr.svg?style=svg)](https://circleci.com/gh/nametake/labelerr)
[![GoDoc](https://godoc.org/github.com/nametake/labelerr?status.svg)](https://godoc.org/github.com/nametake/labelerr)

Package labelerr provides to add label to error.

## Install

`go get github.com/nametake/labelerr`

## Usage

```go
package main

import (
	"fmt"

	"github.com/nametake/labelerr"
	"github.com/pkg/errors"
)

func main() {
	err := errors.New("cause")

	err = errors.Wrap(err, "first")
	err = labelerr.Wrap(err, "warning")
	err = errors.Wrap(err, "second")
	err = errors.Wrap(err, "third")

	fmt.Println(err)                 // third: second: warning: first: cause
	fmt.Println(labelerr.Label(err)) // warning
	fmt.Println(errors.Cause(err))   // cause
}
```
