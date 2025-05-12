package main

import "testing"

func TestCalculateNumberOfSafeReport(t *testing.T) {
	cases := []struct {
		name       string
		isDampened bool
		want       int
	}{
		{
			name:       "without dampener",
			isDampened: false,
			want:       2,
		},
		{
			name:       "with dampener",
			isDampened: true,
			want:       4,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculateNumberOfSafeReport("./test.txt", tt.isDampened)
			if err != nil {
				t.Errorf("an error occurred, %+v", err)
			}

			if got != tt.want {
				t.Errorf("test: %s, got %d, want %d", tt.name, got, tt.want)
			}
		})
	}
}
