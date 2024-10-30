package a14

import (
	"errors"
	"testing"
)

func TestNewError(t *testing.T) {
	originalError := errors.New("original error")
	type args struct {
		msg  string
		kind ErrorKind
		err  error
	}
	tests := []struct {
		name string
		args args
		want A14Error
	}{
		{
			name: "Test NewUnauthorizedError",
			args: args{
				msg:  "Test NewUnauthorizedError",
				kind: UnauthorizedError,
				err:  nil,
			},
			want: &a14ErrorInternal{
				kind: UnauthorizedError,
				msg:  "Test NewUnauthorizedError",
				err:  nil,
			},
		},
		{
			name: "Test NewUnauthorizedError with original error",
			args: args{
				msg:  "Test NewUnauthorizedError",
				kind: UnauthorizedError,
				err:  originalError,
			},
			want: &a14ErrorInternal{
				kind: UnauthorizedError,
				msg:  "Test NewUnauthorizedError",
				err:  originalError,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewError(tt.args.msg, tt.args.kind, tt.args.err); got.Error() != tt.want.Error() {
				t.Errorf("NewUnauthorizedError() = %v, want %v", got, tt.want)
			}
		})
	}
}
