package labelerr

import (
	"fmt"

	"github.com/pkg/errors"
)

func New(msg, label string) error {
	return &labelError{
		err:   errors.New(msg),
		label: label,
	}
}

func Wrap(err error, label string) error {
	if err == nil {
		return nil
	}
	return &labelError{
		err:   err,
		label: label,
	}
}

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
