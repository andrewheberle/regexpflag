package regexpflag_test

import (
	"flag"
	"fmt"
	"os"

	"github.com/andrewheberle/regexpflag"
	"github.com/spf13/pflag"
)

func ExampleMustCompile_pflag() {
	re := regexpflag.MustCompile(".*")

	fs := pflag.NewFlagSet("example", pflag.ExitOnError)
	fs.Var(re, "regexp", "Regular expression")
	fs.Parse([]string{})

	fmt.Printf("Regexp is %q\n", re)
	// Output: Regexp is ".*"
}

func ExampleCompile_pflag() {
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

func ExampleCompile_flag() {
	re, err := regexpflag.Compile(".*")
	if err != nil {
		panic(err)
	}

	flag.Var(re, "regexp", "Regular expression")
	flag.Parse()

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

func ExampleFlag_Regexp_string() {
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
