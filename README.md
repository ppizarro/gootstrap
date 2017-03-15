# Gootstrap

Gootstrap aims to help you bootstrap Go projects.

It provides just a set of scripts that you can copy
to your Go project to help you:

* Vendor things, like conan (no semver, you have integration tests right ?)
* Run tests recursively without including the vendor directory
* Run tests with coverage, coalescing all the packages reports in one

Perhaps it may have more scripts on the future, but the idea is to
keep it very simple and without strong opinions (like Makefiles, docker stuff).

## Installation

Really, just copy the scripts to wherever you want them to be and use them.