package validation

import "testing"

func TestIsValidStatus(t *testing.T) {
	cases := []struct {
		name  string
		input string
		valid bool
	}{
		{"new", "new", true},
		{"in_progress", "in_progress", true},
		{"done", "done", true},
		{"empty", "", false},
		{"weird", "inprogress", false},
		{"caps", "NEW", false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := IsValidStatus(tc.input)
			if got != tc.valid {
				t.Fatalf("IsValidStatus(%q) = %v; want %v", tc.input, got, tc.valid)
			}
		})
	}
}
