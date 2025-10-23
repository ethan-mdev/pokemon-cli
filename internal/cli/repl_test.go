package cli

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "Hello World",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  Go is Awesome!  ",
			expected: []string{"go", "is", "awesome!"},
		},
	}

	for _, c := range cases {
		result := cleanInput(c.input)
		if len(result) != len(c.expected) {
			t.Errorf("cleanInput(%q) = %v; want %v", c.input, result, c.expected)
			continue
		}
	}
}
