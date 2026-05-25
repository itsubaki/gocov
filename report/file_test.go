package report_test

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"testing"

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

	lines := report.SourceLines(string(filedotgo))

	fmt.Println(lines[0])
	fmt.Println(lines[2])

	// Output:
	// package report
	// import (
}

func Test_displayPath(t *testing.T) {
	cases := []struct {
		modulePath  string
		profileFile string
		want        string
	}{
		{
			profileFile: "x/y.go",
			modulePath:  "mod",
			want:        "x/y.go",
		},
		{
			profileFile: "pkg/a.go",
			modulePath:  "",
			want:        "pkg/a.go",
		},
		{
			profileFile: "mod/pkg/a.go",
			modulePath:  "mod",
			want:        "pkg/a.go",
		},
		{
			profileFile: "other/pkg/a.go",
			modulePath:  "mod",
			want:        "other/pkg/a.go",
		},
	}

	for i, tt := range cases {
		got := report.DisplayPath(tt.modulePath, tt.profileFile)
		if got != tt.want {
			t.Fatalf("case %d: got %q, want %q", i, got, tt.want)
		}
	}
}

func Test_sourcePath(t *testing.T) {
	root := t.TempDir()
	os.MkdirAll(filepath.Join(root, "pkg"), 0755)
	os.MkdirAll(filepath.Join(root, "mod"), 0755)
	os.WriteFile(filepath.Join(root, "a.go"), []byte("package x"), 0644)
	os.WriteFile(filepath.Join(root, "pkg/b.go"), []byte("package x"), 0644)
	os.WriteFile(filepath.Join(root, "mod/c.go"), []byte("package x"), 0644)

	cases := []struct {
		modulePath  string
		profileFile string
		ok          bool
	}{
		{
			profileFile: "a.go",
			modulePath:  "",
			ok:          true,
		},
		{
			profileFile: "pkg/b.go",
			modulePath:  "",
			ok:          true,
		},
		{
			profileFile: "mod/c.go",
			modulePath:  "mod",
			ok:          true,
		},
		{
			profileFile: "notfound.go",
			modulePath:  "mod",
			ok:          false,
		},
	}

	for i, c := range cases {
		got, ok := report.SourcePath(root, c.modulePath, c.profileFile)
		if ok != c.ok {
			t.Errorf("case %d: ok = %v want %v", i, ok, c.ok)
		}

		if c.ok && got == "" {
			t.Errorf("case %d: expected path but got empty", i)
		}
	}
}
