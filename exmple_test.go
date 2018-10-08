package chanio

import (
	"fmt"
	"strings"
)

func ExampleScanner_Ch() {
	file := "// comment\n\n1\n2\n"
	reader := strings.NewReader(file)

	scan := NewScanner(reader)
	for line := range scan.Ch() {
		fmt.Println(line)
	}
	if err := scan.Err(); err != nil {
		fmt.Printf("Error: %v", err)
	}
	// Output: // comment
	//
	// 1
	// 2
}

func ExampleScanner_ChFilter() {
	file := "// comment\n\n1\n2\n"
	reader := strings.NewReader(file)

	scan := NewScanner(reader)
	for line := range scan.ChFilter(NotEmptyAndNoComment) {
		fmt.Println(line)
	}
	if err := scan.Err(); err != nil {
		fmt.Printf("Error: %v", err)
	}
	// Output: 1
	// 2
}
