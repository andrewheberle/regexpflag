package regexpflag_test

import (
	"flag"
	"fmt"
	"os"

	"github.com/andrewheberle/regexpflag"
	"github.com/spf13/pflag"
)

func ExampleMustCompile() {
	re := regexpflag.MustCompile(".*")

	flag.Var(re, "regexp", "Regular expression")
	//
	// In this example flag.Parse() is commented out as this forms part of the tests
	// of this module, so parsing the command line flags is disabled.
	//
	// In a real program you would need to call flag.Parse()
	//
	// flag.Parse()

	fmt.Printf("Regexp is %q\n", re)
	// Output: Regexp is ".*"
}

func ExampleCompile() {
	re, err := regexpflag.Compile(".*")
	if err != nil {
		panic(err)
	}

	flag.Var(re, "regexp", "Regular expression")
	//
	// In this example flag.Parse() is commented out as this forms part of the tests
	// of this module, so parsing the command line flags is disabled.
	//
	// In a real program you would need to call flag.Parse()
	//
	// flag.Parse()

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
	// .*
}
