/*
Package chanio provides a Scanner for reading newline-delimited text from an
io.Reader or file line by line from a channel.

The chanio.Scanner is a lightweight wrapper of the bufio.Scanner.

Reading a file line by line from a channel is that easy:

	scan := chanio.NewFileScanner("file.txt")
	for line := range scan.Ch() {
		log.Println(line)
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}
*/
package chanio

import (
	"bufio"
	"io"
	"os"
)

// Scanner provides a simple interface for reading newline-delimited text from
// an io.Reader or file line by line from a channel.
type Scanner struct {
	r   io.Reader
	err error
}

// NewScanner returns a new Scanner to read from io.Reader r.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: r}
}

// NewFileScanner returns a new Scanner to read from the named file.
func NewFileScanner(name string) *Scanner {
	file, err := os.Open(name)
	if err != nil {
		return &Scanner{err: err}
	}
	return NewScanner(file)
}

// Ch returns a channel of scanned lines. The channel is closed when the scan
// stops, either by reaching the end of the input or an error. If the scanned
// io.Reader also implements io.Closer the Close method will be called when the
// scan has stopped.
// After the returned channel has been closed, the Err method will return any
// error that occurred during scanning.
func (s *Scanner) Ch() <-chan string {
	return s.ChFilter(all)
}

// ChFilter works analogously to the Ch method but takes a filter function as
// argument that will filter out all lines for which the filter function
// returns false.
func (s *Scanner) ChFilter(filter FilterFunc) <-chan string {
	ch := make(chan string)
	go s.ch(ch, filter)
	return ch
}

func (s *Scanner) ch(ch chan string, filter FilterFunc) {
	defer close(ch)
	if s.r == nil || s.err != nil {
		return
	}

	if c, ok := s.r.(io.Closer); ok {
		defer c.Close()
	}

	scanner := bufio.NewScanner(s.r)
	for scanner.Scan() {
		text := scanner.Text()
		if filter(text) {
			ch <- text
		}
	}
	if scanner.Err() != nil {
		s.err = scanner.Err()
	}
}

// Err returns the first error that was encountered by the Scanner.
func (s *Scanner) Err() error {
	return s.err
}
