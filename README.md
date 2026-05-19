# gocov

`gocov` turns a Go coverage profile into a rich, local HTML report. Run it from the root of the Go repository that produced the profile:

```sh
go install github.com/itsubaki/gocov@latest
gocov -f coverage.txt
```

By default it writes `coverage.html` in the current directory.

## Usage

```sh
gocov -f coverage.txt -o coverage.html
```

Flags:

- `-f`: path to a Go coverage profile, usually created by `go test ./... -coverprofile=coverage.txt`
- `-o`: output path for the generated HTML report
- `-root`: Go repository root; defaults to the current directory

The report is a single self-contained HTML file with summary charts, a hierarchical directory coverage pie weighted by coverable lines, directory bars, searchable file navigation, and line-level coverage highlighting.
