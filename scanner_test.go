package chanio

import (
	"strings"
	"testing"
)

func TestScanner(t *testing.T) {
	var tests = []struct {
		lines    string
		filter   FilterFunc
		expected []string
	}{
		{"a\nb\n\n#c\n//d\ne\n", all, []string{"a", "b", "", "#c", "//d", "e"}},
		{"a\nb\n\n#c\n//d\ne\n", NotEmpty, []string{"a", "b", "#c", "//d", "e"}},
		{"a\nb\n\n#c\n//d\ne\n", NotEmptyAndNoComment, []string{"a", "b", "e"}},
	}

	for _, test := range tests {
		r := strings.NewReader(test.lines)
		s := NewScanner(r)

		i := 0
		for line := range s.ChFilter(test.filter) {
			if line != test.expected[i] {
				t.Errorf("%s != %s", test.expected[i], line)
			}
			i++
		}
		if s.Err() != nil {
			t.Error("No err expected")
		}
	}
}
