package lvlerr

import "testing"

func TestNew(t *testing.T) {
	tests := []struct {
		name  string
		msg   string
		level string
		want  string
	}{
		{
			name:  "simple",
			msg:   "error message",
			level: "warning",
			want:  "warning: error message",
		},
		{
			name:  "empty msg",
			msg:   "",
			level: "warning",
			want:  "warning: ",
		},
		{
			name:  "empty level",
			msg:   "error message",
			level: "",
			want:  ": error message",
		},
		{
			name:  "empty",
			msg:   "",
			level: "",
			want:  ": ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.msg, tt.level); got.Error() != tt.want {
				t.Errorf("New.Error(): got: %v, want %v", got, tt.want)
			}
		})
	}
}
