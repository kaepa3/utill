package file

import (
	"os"
	"testing"
)

func TestExists(t *testing.T) {
	testfile := "test"
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "exist normal test", args: args{testfile}, want: true},
	}
	os.Create(testfile)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Exists(tt.args.path); got != tt.want {
				t.Errorf("Exists() = %v, want %v", got, tt.want)
			}
		})
	}
	os.Remove(testfile)
}
