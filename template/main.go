package template

// Main template for main.go file
const Main = `package main

import (
	"flag"
	"fmt"
	"os"

	"{{.Module}}/pkg/version"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	fmt.Println("That's it folks!")
}

func run() error {
	flagSet := flag.NewFlagSet("{{.Project}}", flag.ExitOnError)
	showVersion := flagSet.Bool("version", false, "Show version")

	if err := flagSet.Parse(os.Args[1:]); err != nil {
		return fmt.Errorf("failure on parse command line argunments: %q", err)
	}

	if *showVersion {
		fmt.Print(version.Version())
		return nil
	}

	return nil
}
`
