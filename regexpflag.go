package regexpflag

import (
	"flag"
	"regexp"

	"github.com/spf13/pflag"
)

var (
	_ pflag.Value = (*Flag)(nil)
	_ flag.Value  = (*Flag)(nil)
)

// Flag implements the [pflag.Value] interface
type Flag struct {
	r *regexp.Regexp
}

// MustCompile is like [Compile] but panics if the expression cannot be parsed.
func MustCompile(expr string) *Flag {
	r := regexp.MustCompile(expr)

	return &Flag{r}
}

// Compile parses a regular expression and returns, if successful, a Flag
// that can be used as a [pflag.Value]
func Compile(expr string) (*Flag, error) {
	r, err := regexp.Compile(expr)
	if err != nil {
		return nil, err
	}

	return &Flag{r}, nil
}

// Set is requried to satisfy the [pflag.Value] interface
func (f *Flag) Set(expr string) error {
	r, err := regexp.Compile(expr)
	if err != nil {
		return err
	}
	f.r = r

	return nil
}

// String is requried to satisfy the [pflag.Value] interface
func (f Flag) String() string {
	return f.r.String()
}

// String is requried to satisfy the [pflag.Value] interface
func (f Flag) Type() string {
	return "regexp"
}

// Regexp returns the underlying [regexp.Regexp]
func (f Flag) Regexp() *regexp.Regexp {
	return f.r
}
