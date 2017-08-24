# Gootstrap

Gootstrap aims to help you bootstrap Go projects.

It provides just a set of scripts that you can copy
to your Go project to help you:

* Vendor things, like conan (no semver, you have integration tests right ?)
* Run tests recursively without including the vendor directory
* Run tests with coverage, coalescing all the packages reports in one
* Cool static analysis
* Embedding --version on binaries using git commit tag

All commands happens on top of docker since it
is the basis of our development environments and production deployment.

It has much less usefulness in Go than other languages (since Go is very
simple, specially with vendoring) but at least avoids differences on Go
versions between developers and CI servers, also promotes uniformity on how
we work with other languages.

We are also striving to do governance through code, spreading
good practices and good tooling out to new projects, reducing time
to bootstrap new projects.

## Installation

Run:

```
go get github.com/NeowayLabs/gootstrap
```

And that is it =)

## Usage

After creating the local directory for your Go project
(or cloning it), you will have something like:

```
$GOPATH/src/domain/yourproject
```

Just:

```
cd $GOPATH/src/domain/yourproject
./gootstrap <option1> <option2>
```

The help of the project should be enough to elaborate
on available options.

Just be aware that besides running everything inside docker
containers it is important that the project is inside the
GOPATH of your host.

Some details on how the code is mapped to the containers
depend on this, also if you do so you will be able
to build and run tests directly on your host too if
you want (autocomplete and code navigation will also
work properly in vendored dependencies).

Also be aware that gootstrap WILL overwrite files if it
finds files with the same name, so have everything commited before
running it. Files that are unknown to gootstrap will be left alone.
