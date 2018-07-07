package lvlerr

import (
	"fmt"

	"github.com/pkg/errors"
)

func Wrap(err error, level string) error {
	return &lvlError{
		err:   err,
		level: level,
	}
}

func WithMessage(err error, level, msg string) error {
	return &lvlError{
		err:   errors.Wrap(err, msg),
		level: level,
	}
}

func Level(err error) string {
	for err != nil {
		level, ok := err.(leveler)
		if ok {
			return level.Level()
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

type leveler interface {
	Level() string
}

var (
	_ error   = (*lvlError)(nil)
	_ causer  = (*lvlError)(nil)
	_ leveler = (*lvlError)(nil)
)

type lvlError struct {
	err   error
	level string
}

func (l *lvlError) Error() string {
	return fmt.Sprintf("%s: %s", l.level, l.err)
}

func (l *lvlError) Cause() error {
	return l.err
}

func (l *lvlError) Level() string {
	return l.level
}
