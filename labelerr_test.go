package labelerr

import (
	"testing"

	"github.com/pkg/errors"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name  string
		msg   string
		label string
		want  string
	}{
		{
			name:  "simple",
			msg:   "error message",
			label: "warning",
			want:  "warning: error message",
		},
		{
			name:  "empty msg",
			msg:   "",
			label: "warning",
			want:  "warning: ",
		},
		{
			name:  "empty label",
			msg:   "error message",
			label: "",
			want:  ": error message",
		},
		{
			name:  "empty",
			msg:   "",
			label: "",
			want:  ": ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.msg, tt.label); got.Error() != tt.want {
				t.Errorf("New.Error(): got: %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWrap(t *testing.T) {
	tests := []struct {
		name  string
		err   error
		label string
		want  string
	}{
		{
			name:  "simple",
			err:   errors.New("error"),
			label: "label",
			want:  "label: error",
		},
		{
			name:  "empty label",
			err:   errors.New("error"),
			label: "",
			want:  ": error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Wrap(tt.err, tt.label).Error(); got != tt.want {
				t.Errorf("Wrap.Error(): got: %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWrapNil(t *testing.T) {
	if got := Wrap(nil, "no error"); got != nil {
		t.Errorf("Wrap(nil, \"no error\"): got: %v, want nil", got)
	}
}

func TestWithMessage(t *testing.T) {
	tests := []struct {
		name  string
		err   error
		label string
		msg   string
		want  string
	}{
		{
			name:  "simple",
			err:   errors.New("error"),
			label: "label",
			msg:   "message",
			want:  "message: label: error",
		},
		{
			name:  "empty label",
			err:   errors.New("error"),
			label: "",
			msg:   "message",
			want:  "message: : error",
		},
		{
			name:  "empty message",
			err:   errors.New("error"),
			label: "label",
			msg:   "",
			want:  ": label: error",
		},
		{
			name:  "empty label and message",
			err:   errors.New("error"),
			label: "",
			msg:   "",
			want:  ": : error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithMessage(tt.err, tt.label, tt.msg); got.Error() != tt.want {
				t.Errorf("WithMessage.Error(): got: %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithMessageNil(t *testing.T) {
	if got := WithMessage(nil, "no error", "no msg"); got != nil {
		t.Errorf("WithMessage(nil, \"no error\", \"no msg\"): got: %v, want nil", got)
	}
}

func TestLabel(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want string
	}{
		{
			name: "New()",
			err:  New("message", "label"),
			want: "label",
		},
		{
			name: "Wrap()",
			err:  Wrap(errors.New("error"), "label"),
			want: "label",
		},
		{
			name: "WithMessage()",
			err:  WithMessage(errors.New("error"), "label", "message"),
			want: "label",
		},
		{
			name: "no label error",
			err:  errors.New("error"),
			want: "",
		},
		{
			name: "nil",
			err:  nil,
			want: "",
		},
		{
			name: "use errors.Wrap New()",
			err:  errors.Wrap(New("message", "label"), "wrapped"),
			want: "label",
		},
		{
			name: "use errors.Wrap Wrap()",
			err:  errors.Wrap(Wrap(errors.New("error"), "label"), "wrapped"),
			want: "label",
		},
		{
			name: "use errors.Wrap WithMessage()",
			err:  errors.Wrap(WithMessage(errors.New("error"), "label", "message"), "wrapped"),
			want: "label",
		},
		{
			name: "use multiple label error(Wrap and New)",
			err:  Wrap(New("message", "inner"), "outer"),
			want: "outer",
		},
		{
			name: "use multiple label error(WithMessage and New)",
			err:  WithMessage(New("message", "inner"), "outer", "message"),
			want: "outer",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Label(tt.err); got != tt.want {
				t.Errorf("Label(): got: %v, want %v", got, tt.want)
			}
		})
	}
}
