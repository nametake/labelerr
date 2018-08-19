// Package labelerr provides to add label to error.
//
// The labelerr.Wrap function returns new erorr that added label to the original error.
// The label can be retrieved with labelerr.Label function.
//
// This is supported to be used with github.com/pkg/errors.
// So, labeled errors implement the Cause method and also supported to errors.Cause.
// And, it dosen't lose stacktrace that given by github.com/pkg/errors
package labelerr

import (
	"fmt"

	"github.com/pkg/errors"
)

// New returns an error with supplied message and label.
// New also records the stack trace at the point it was called.
func New(msg, label string) error {
	return &labelError{
		err:   errors.New(msg),
		label: label,
	}
}

// Wrap returns an error annotating err with a stack trace at the point Wrap is called,
// and the supplied label. If err is nil, Wrap returns nil.
func Wrap(err error, label string) error {
	if err == nil {
		return nil
	}
	return errors.WithStack(&labelError{
		err:   err,
		label: label,
	})
}

// WithMessage returns an error annotating err with a stack trace at the point Wrap is called,
// and the supplied label and message. If err is nil, Wrap returns nil.
func WithMessage(err error, label, msg string) error {
	if err == nil {
		return nil
	}
	return errors.Wrap(
		&labelError{
			err:   err,
			label: label,
		},
		msg,
	)
}

// Label returns the last assigned label.
// Even if errors are wrapped using errors.Wrap.
func Label(err error) string {
	for err != nil {
		label, ok := err.(labeler)
		if ok {
			return label.Label()
		}
		cause, ok := err.(causer)
		if !ok {
			break
		}
		err = cause.Cause()
	}
	return ""
}

type causer interface {
	Cause() error
}

type labeler interface {
	Label() string
}

var (
	_ error         = (*labelError)(nil)
	_ fmt.Formatter = (*labelError)(nil)
	_ causer        = (*labelError)(nil)
	_ labeler       = (*labelError)(nil)
)

type labelError struct {
	err   error
	label string
}

func (l *labelError) Error() string {
	return fmt.Sprintf("%s: %s", l.label, l.err)
}

func (l *labelError) Cause() error {
	return l.err
}

func (l *labelError) Label() string {
	return l.label
}

func (l *labelError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%+v", l.err)
			return
		}
		fallthrough
	case 's':
		fmt.Fprintf(s, "%s", l.err)
	case 'q':
		fmt.Fprintf(s, "%q", l.err)
	}
}
