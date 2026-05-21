package profile_test

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/itsubaki/gocov/profile"
)

//go:embed testdata/coverage.txt
var coverage []byte

func ExampleParse() {
	tmpFile, err := os.CreateTemp("", "coverage-*.txt")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.Write(coverage); err != nil {
		panic(err)
	}
	tmpFile.Close()

	prof, err := profile.Parse(tmpFile.Name())
	if err != nil {
		panic(err)
	}

	fmt.Println(prof.Mode)
	fmt.Println(len(prof.Blocks))

	// Output:
	// atomic
	// 249
}
