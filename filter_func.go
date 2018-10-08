package chanio

// FilterFunc is the signature of the filter function. A filter function is
// called on each scanned line of the Scanner. A line is put in the channel if
// the filter function returns true and is filtered out otherwise.
type FilterFunc func(line string) bool

func all(line string) bool {
	return true
}

// NotEmpty is a filter function for a Scanner that filters out empty lines
// of length zero.
func NotEmpty(line string) bool {
	return len(line) > 0
}

// NotEmptyAndNoComment is a filter function for a Scanner that filters out
// both empty lines of length zero and comment lines starting with "#" or "//".
func NotEmptyAndNoComment(line string) bool {
	switch len(line) {
	case 0:
		return false
	case 1:
		return line[0] != '#'
	default:
		return line[0] != '#' && (line[0] != '/' || line[1] != '/')
	}
}
