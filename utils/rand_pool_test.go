package utils

import (
	"testing"
)

func TestMyRand_NextInt64(t *testing.T) {
	my := NewMyRand()
	tests := []struct {
		name string
		want int64
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			want: 0,
		},
		{
			name: "test1",
			want: 0,
		},
		{
			name: "test1",
			want: 0,
		},
		{
			name: "test1",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(my.NextInt64())
		})
	}
}
