# regexpflag

[![Go Report Card](https://goreportcard.com/badge/github.com/andrewheberle/regexpflag?logo=go&style=flat-square)](https://goreportcard.com/report/github.com/andrewheberle/regexpflag)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/andrewheberle/regexpflag)

This package provides a `Flag` type that can be used as a custom flag for
`flag` and `github.com/sp13/pflag` as it satisifes the `flag.Value` and
`pflag.Value` interfaces.

Command line flag values are a string that is compiled into a
[*regexp.Regexp](https://pkg.go.dev/regexp#Regexp) using
[regexp.Compile](https://pkg.go.dev/regexp#Compile).

The resulting [*regexp.Regexp](https://pkg.go.dev/regexp#Regexp) can be
accessed in yor code using `Regexp()`

## Example

```go
package main

import (
    "flag"

    "github.com/andrewheberle/regexpflag"
)

func main() {
	re := regexpflag.MustCompile(".*")

	flag.Var(re, "regexp", "Regular expression")
	flag.Parse()

	fmt.Printf("Regexp is %q\n", re)
	// Output: Regexp is ".*"
}
```
