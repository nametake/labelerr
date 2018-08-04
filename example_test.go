package labelerr_test

import (
	"fmt"

	"github.com/nametake/labelerr"
	"github.com/pkg/errors"
)

func ExampleLabel() {
	err := errors.New("cause")

	err = errors.Wrap(err, "first")
	err = labelerr.Wrap(err, "warning")
	err = errors.Wrap(err, "second")
	err = errors.Wrap(err, "third")

	fmt.Println(err)
	fmt.Println(labelerr.Label(err))
	fmt.Println(errors.Cause(err))

	// Output:
	// third: second: warning: first: cause
	// warning
	// cause
}
