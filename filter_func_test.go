package chanio

import (
	"testing"
)

func TestFilterFunc(t *testing.T) {
	var tests = []struct {
		line     string
		filter   FilterFunc
		expected bool
	}{
		{"", NotEmpty, false},
		{"a", NotEmpty, true},
		{"", NotEmptyAndNoComment, false},
		{"#", NotEmptyAndNoComment, false},
		{"//", NotEmptyAndNoComment, false},
		{"/", NotEmptyAndNoComment, true},
		{"a", NotEmptyAndNoComment, true},
		{"ab", NotEmptyAndNoComment, true},
		{"abc", NotEmptyAndNoComment, true},
		{"# a", NotEmptyAndNoComment, false},
		{"// ab", NotEmptyAndNoComment, false},
		{"/ abc", NotEmptyAndNoComment, true},
		{" /abc", NotEmptyAndNoComment, true},
	}

	for _, test := range tests {
		actual := test.filter(test.line)
		if actual != test.expected {
			t.Errorf("`%s`: %v != %v", test.line, test.expected, actual)
		}
	}
}
