# gocov

[![PkgGoDev](https://pkg.go.dev/badge/github.com/itsubaki/gocov)](https://pkg.go.dev/github.com/itsubaki/gocov)
[![Go Report Card](https://goreportcard.com/badge/github.com/itsubaki/gocov?style=flat-square)](https://goreportcard.com/report/github.com/itsubaki/gocov)
[![tests](https://github.com/itsubaki/gocov/workflows/tests/badge.svg)](https://github.com/itsubaki/gocov/actions)
[![coverage](https://itsubaki.github.io/gocov/coverage.svg)](https://itsubaki.github.io/gocov/)

`gocov` turns a Go coverage profile into a rich, local HTML report. Run it from the root of the Go repository that produced the profile:

```sh
go test -coverprofile=coverage.txt -covermode=atomic
gocov -f coverage.txt
```

By default it writes `coverage.html` in the current directory.

## Installation

```sh
go install github.com/itsubaki/gocov@latest
```

## Usage

```sh
gocov -f coverage.txt -o coverage.html
```

Flags:

- `-f`: path to a Go coverage profile, usually created by `go test ./... -coverprofile=coverage.txt`
- `-o`: output path for the generated HTML report

The report is a single self-contained HTML file with summary charts, a hierarchical directory coverage pie weighted by coverable lines, directory bars, searchable file navigation, and line-level coverage highlighting.
