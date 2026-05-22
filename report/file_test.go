package report_test

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/itsubaki/gocov/report"
)

//go:embed file.go
var filedotgo []byte

func Example_sourceLines() {
	tmpFile, err := os.CreateTemp("", "coverage-*.txt")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.Write(filedotgo); err != nil {
		panic(err)
	}
	tmpFile.Close()

	lines, err := report.SourceLines(tmpFile.Name())
	if err != nil {
		panic(err)
	}

	fmt.Println(lines[0])
	fmt.Println(lines[2])

	// Output:
	// package report
	// import (
}
