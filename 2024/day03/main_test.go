package main

import "testing"

func TestCalculateMultiplicationResult(t *testing.T) {
	cases := []struct {
		name             string
		isDoDontsEnabled bool
		input            string
		want             int
	}{
		{
			name:             "without do and donts",
			isDoDontsEnabled: false,
			input:            "./test_part1.txt",
			want:             161,
		},
		{
			name:             "with do and donts",
			isDoDontsEnabled: true,
			input:            "./test_part2.txt",
			want:             48,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculateMultiplicationResult(tt.input, tt.isDoDontsEnabled)
			if err != nil {
				t.Errorf("an error occurred, %+v", err)
			}

			if got != tt.want {
				t.Errorf("test: %s, got %d, want %d", tt.name, got, tt.want)
			}
		})
	}
}
