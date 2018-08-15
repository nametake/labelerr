package labelerr_test

import (
	"fmt"

	"github.com/nametake/labelerr"
	"github.com/pkg/errors"
)

func cause() error {
	return errors.New("cause")
}

func first() error {
	err := cause()
	err = labelerr.Wrap(err, "warning")
	return errors.Wrap(err, "first")

	// one line:
	// return labelerr.WithMessage(err, "warning", "first")
}

func second() error {
	err := first()
	return errors.Wrap(err, "second")
}

func Example() {
	err := second()

	fmt.Println(err)
	fmt.Println(labelerr.Label(err))
	fmt.Println(errors.Cause(err))

	// Output:
	// second: first: warning: cause
	// warning
	// cause
}
