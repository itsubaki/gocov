package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/itsubaki/gocov/profile"
	"github.com/itsubaki/gocov/render"
	"github.com/itsubaki/gocov/report"
)

func main() {
	var profilePath, outputPath, root string
	flag.StringVar(&profilePath, "f", "coverage.txt", "path to a Go coverage profile")
	flag.StringVar(&outputPath, "o", "coverage.html", "path to the HTML report to write")
	flag.StringVar(&root, "root", ".", "path to the Go repository root")
	flag.Parse()

	if strings.TrimSpace(profilePath) == "" {
		fmt.Fprintln(os.Stderr, "gocov: coverage profile path is required")
		os.Exit(1)
	}

	input, err := filepath.Abs(profilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "gocov: %v\n", err)
		os.Exit(1)
	}

	output, err := filepath.Abs(outputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "gocov: %v\n", err)
		os.Exit(1)
	}

	dir, err := filepath.Abs(root)
	if err != nil {
		fmt.Fprintf(os.Stderr, "gocov: %v\n", err)
		os.Exit(1)
	}

	if err := run(input, output, dir); err != nil {
		fmt.Fprintf(os.Stderr, "gocov: %v\n", err)
		os.Exit(1)
	}
}

func run(input, output, dir string) error {
	prof, err := profile.Parse(input)
	if err != nil {
		return err
	}

	rep, err := report.New(prof, report.Options{
		RootPath:    dir,
		ProfilePath: input,
		OutputPath:  output,
		GeneratedAt: time.Now(),
	})
	if err != nil {
		return err
	}

	file, close, err := touch(output)
	if err != nil {
		return err
	}
	defer func() {
		if err := close(); err != nil {
			fmt.Fprintf(os.Stderr, "gocov: %v\n", err)
		}
	}()

	if err := render.HTML(file, rep); err != nil {
		return err
	}

	return nil
}

func touch(path string) (*os.File, func() error, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil, nil, err
	}

	file, err := os.Create(path)
	if err != nil {
		return nil, nil, err
	}

	return file, file.Close, nil
}
