package template

const Cmd = `package main

import (
	"flag"
	"fmt"
	"{{.Module}}/pkg/version"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	showVersion := flag.Bool("version", false, "Show version")
	flag.Parse()

	if *showVersion {
		fmt.Print(version.Version())
		return nil
	}
	return nil
}
`
