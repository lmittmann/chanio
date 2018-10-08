# Package chanio [![Build Status](https://travis-ci.com/lmittmann/chanio.svg?branch=master)](https://travis-ci.com/lmittmann/chanio) [![GoDoc](https://godoc.org/github.com/lmittmann/chanio?status.svg)](https://godoc.org/github.com/lmittmann/chanio) [![Go Report Card](https://goreportcard.com/badge/github.com/lmittmann/chanio)](https://goreportcard.com/report/github.com/lmittmann/chanio)


```go
import "github.com/lmittmann/chanio"
```

Package chanio provides a Scanner for reading newline-delimited text from an io.Reader or file line by line from a channel.

The chanio.Scanner is a lightweight wrapper of the bufio.Scanner.

Reading a file line by line from a channel is that easy:
```go
scan := chanio.NewFileScanner("file.txt")
for line := range scan.Ch() {
	log.Println(line)
}
if err := scan.Err(); err != nil {
	log.Fatal(err)
}
```

## Index
* [type Scanner](#type-scanner)
    * [func NewScanner(r io.Reader) *Scanner](#func-newscanner)
    * [func NewFileScanner(name string) *Scanner](#func-newfilescanner)
    * [func (s *Scanner) Ch() <-chan string](#func-scanner-ch)
    * [func (s *Scanner) ChFilter(filter FilterFunc) <-chan string](#func-scanner-chfilter)
    * [func (s *Scanner) Err() error](#func-scanner-err)
* [type FilterFunc](#type-filterfunc)
* [func NotEmpty(line string) bool](#func-notempty)
* [func NotEmptyAndNoComment(line string) bool](#func-notemptyandnocomment)


## type [Scanner](scanner.go#L27)
```go
type Scanner struct {
    // contains filtered or unexported fields
}
```
Scanner provides a simple interface for reading newline-delimited text from an io.Reader or file line by line from a channel.


## func [NewScanner](scanner.go#L33)
```go
func NewScanner(r io.Reader) *Scanner
```
NewScanner returns a new Scanner to read from io.Reader r.


## func [NewFileScanner](scanner.go#L38)
```go
func NewFileScanner(name string) *Scanner
```
NewFileScanner returns a new Scanner to read from the named file.


## func (*Scanner) [Ch](scanner.go#L52)
```go
func (s *Scanner) Ch() <-chan string
```
Ch returns a channel of scanned lines. The channel is closed when the scan stops, either by reaching the end of the input or an error. If the scanned io.Reader also implements io.Closer the Close method will be called when the scan has stopped. After the returned channel has been closed, the Err method will return any error that occurred during scanning.


## func (*Scanner) [ChFilter](scanner.go#L59)
```go
func (s *Scanner) ChFilter(filter FilterFunc) <-chan string
```
ChFilter works analogously to the Ch method but takes a filter function as argument that will filter out all lines for which the filter function returns false.


## func (*Scanner) [Err](scanner.go#L88)
```go
func (s *Scanner) Err() error
```
Err returns the first error that was encountered by the Scanner.


## type [FilterFunc](filter_func.go#L8)
```go
type FilterFunc func(line string) bool
```
FilterFunc is the signature of the filter function. A filter function is called on each scanned line of the Scanner. A line is put in the channel if the filter function returns true and is filtered out otherwise.


## func [NotEmpty](filter_func.go#L14)
```go
func NotEmpty(line string) bool
```
NotEmpty is a filter function for a Scanner that filters out empty lines of length zero.


## func [NotEmptyAndNoComment](filter_func.go#L20)
```go
func NotEmptyAndNoComment(line string) bool
```
NotEmptyAndNoComment is a filter function for a Scanner that filters out both empty lines of length zero and comment lines starting with "#" or "//".
