package labelerr

import "testing"

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
			name:  "empty level",
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
