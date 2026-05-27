package regexpflag_test

import (
	"fmt"
	"os"

	"github.com/andrewheberle/regexpflag"
	"github.com/spf13/pflag"
)

func ExampleMustCompile() {
	re := regexpflag.MustCompile(".*")

	fs := pflag.NewFlagSet("example", pflag.ExitOnError)
	fs.Var(re, "regexp", "Regular expression")
	fs.Parse([]string{})

	fmt.Printf("Regexp is %q\n", re)
	// Output: Regexp is ".*"
}

func ExampleCompile() {
	re, err := regexpflag.Compile(".*")
	if err != nil {
		panic(err)
	}

	fs := pflag.NewFlagSet("example", pflag.ExitOnError)
	fs.Var(re, "regexp", "Regular expression")
	fs.Parse([]string{})

	fmt.Printf("Regexp is %q\n", re)
	// Output: Regexp is ".*"
}

func ExampleFlag_Type() {
	// This example doesn't explicitly show the calling of Type(),
	// however the pflag package calls Type() as part of it's display
	// of command line flag defaults that is being manually called below.
	re, err := regexpflag.Compile(".*")
	if err != nil {
		panic(err)
	}

	fs := pflag.NewFlagSet("example", pflag.ExitOnError)
	fs.SetOutput(os.Stdout)
	fs.Var(re, "regexp", "Regular expression")
	fs.Parse([]string{})
	fs.PrintDefaults()
	// Output:
	//		--regexp regexp   Regular expression (default .*)
}

func ExampleFlag_Regexp() {
	re, err := regexpflag.Compile(".*")
	if err != nil {
		panic(err)
	}

	fmt.Println(re.Regexp().String())
	// Output:
	// .*
}

func ExampleFlag_Regexp_find() {
	re, err := regexpflag.Compile("^why ([a-z]*) there$")
	if err != nil {
		panic(err)
	}

	match := re.Regexp().FindStringSubmatch("why hello there")
	if match == nil {
		panic("no match")
	}

	if len(match) != 2 {
		panic("didn't get submatch")
	}
	fmt.Printf("%s", match[1])
	// Output:
	// hello
}
