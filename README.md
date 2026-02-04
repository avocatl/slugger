# Slugger

An extensible and configurable library with language specific support for generating slugs.

## Motivation

> TL;DR This is yet another slug generation library.

Creating a slug is common for applications and services and
althought there are several existing libraries the maintenance or purpose
doesn't clearly align with our needs.

## Usage

To generate slugs you will need to define a configuration, the library ships several
options to have a fluent approach for this task.

Here is a basic example for an slugger using English and no suffixes.

```go
package main

import "github.com/avocatl/slugger"

func main() {
	cfg := slugger.NewConfig(
		slugger.WithLanguage(slugger.English),
		slugger.WithNoSuffix(),
	)

	s := slugger.New(cfg)

	slug := s.Slugify("Hello, World!")
	println(slug) // Output: hello-world
}
```