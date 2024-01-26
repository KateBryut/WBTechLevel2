package main

import "testing"

func Test_unwrapString(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "test #1",
			input: `a4bc2d5e`,
			want:  "aaaabccddddde",
		},
		{
			name:  "test #2",
			input: `abcd`,
			want:  "abcd",
		},
		{
			name:  "test #3",
			input: `45`,
			want:  "",
		},
		{
			name:  "test #4",
			input: ``,
			want:  "",
		},
		{
			name:  "test #5",
			input: `qwe\4\5`,
			want:  "qwe45",
		},
		{
			name:  "test #6",
			input: `qwe\45`,
			want:  "qwe44444",
		},
		{
			name:  "test #6",
			input: `qwe\\5`,
			want:  `qwe\\\\\`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := unwrapString(tt.input)
			if got != tt.want {
				t.Errorf("unwrapString() = %v, want %v", got, tt.want)
			}
		})
	}
}
