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
