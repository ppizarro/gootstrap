package template

const Readme = `# {{.Project}}

TODO brief intro

## Why

TODO

## Testing

Run:` +

	"\n\n```\nmake test\n```\n" +

	"Open generated coverage on a browser:\n\n" +

	"```\nmake coverage\n```\n" +

	"To perform static analysis:\n\n```\nmake lint\n```\n" +

	`
## Releasing

Run:` +

	"\n\n```\nmake release version=<version>\n```\n" +

	`
It will create a git tag with the provided **<version>**
and build and publish a docker image.

## Git Hooks

To install the project githooks run:` + "\n\n```\nmake githooks\n```\n"
