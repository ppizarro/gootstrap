package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/ppizarro/gootstrap"
)

// GoVersion the Go version
const GoVersion = "1.13.5"

// GoDigest the hash digest from golang image
const GoDigest = "sha256:3997ddbf0c313c5b629eccbdf14e098f4d3ad23cb1f2d9f8cb66707c9ee4ce79" // golang:1.13.5-stretch

// CILintVersion the golangci-lint version
const CILintVersion = "1.21.0"

func main() {

	outputdir := ""
	module := ""
	dockerimg := ""

	flag.StringVar(
		&dockerimg,
		"image",
		"",
		"docker image of the project",
	)
	flag.StringVar(
		&module,
		"module",
		"",
		"The module name of the project, like: 'github.com/ppizarro/gootstrap'",
	)
	flag.StringVar(
		&outputdir,
		"output-dir",
		getcwd(),
		"directory where the generated files are going to be saved",
	)

	flag.Parse()

	if module == "" {
		fmt.Println("-module is an obligatory parameter")
		os.Exit(1)
	}

	if dockerimg == "" {
		fmt.Println("-docker-image is an obligatory parameter")
		os.Exit(1)
	}

	fmt.Printf("creating project module[%s] docker-image[%s] files at dir[%s]\n",
		module, dockerimg, outputdir)

	project, err := parseNameFromModule(module)
	abortonerr(err)

	cfg := gootstrap.Config{
		Project:       project,
		Module:        module,
		DockerImg:     dockerimg,
		GoVersion:     GoVersion,
		GoDigest:      GoDigest,
		CILintVersion: CILintVersion,
	}
	err = gootstrap.CreateProject(cfg, outputdir)
	abortonerr(err)
}

func getcwd() string {
	wd, err := os.Getwd()
	abortonerr(err)
	return wd
}

func abortonerr(err error) {
	if err != nil {
		panic(err)
	}
}

func parseNameFromModule(module string) (string, error) {
	// Go modules are like this: github.com/ppizarro/gootstrap
	// Lets assume that the last component of the path is the project name
	parsed := strings.Split(module, "/")
	if len(parsed) == 1 {
		return "", fmt.Errorf("invalid module[%s] cant extract project name from it", module)
	}

	return parsed[len(parsed)-1], nil
}
