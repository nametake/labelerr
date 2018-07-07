package lvlerr

import "fmt"

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
