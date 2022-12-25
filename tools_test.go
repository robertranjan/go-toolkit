package toolkit

import "testing"

func TestTools_RandomString(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		tr   *Tools
		args args
		want int
	}{
		{
			name: "success case",
			tr:   &Tools{},
			args: args{n: 10},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Tools{}
			if got := len(tr.RandomString(tt.args.n)); got != tt.want {
				t.Errorf("Tools.RandomString() = %v, want %v", got, tt.want)
			}
		})
	}
}
