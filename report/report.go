package report

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/itsubaki/gocov/profile"
)

type Report struct {
	GeneratedAt  time.Time
	RootPath     string
	ProfilePath  string
	OutputPath   string
	Mode         string
	ModulePath   string
	Stats        Stats
	Files        []*File
	Directories  []*Directory
	MissingFiles []string
}

type Options struct {
	RootPath    string
	ProfilePath string
	OutputPath  string
	GeneratedAt time.Time
}

func New(prof *profile.Profile, opts Options) (*Report, error) {
	blocks := make(map[string][]*profile.Block)
	for _, b := range prof.Blocks {
		blocks[b.File] = append(blocks[b.File], b)
	}

	blockFiles := make([]string, 0, len(blocks))
	for b := range blocks {
		blockFiles = append(blockFiles, b)
	}
	sort.Strings(blockFiles)

	// module path
	file, err := os.Open(filepath.Join(opts.RootPath, "go.mod"))
	if err != nil {
		return nil, fmt.Errorf("open go.mod: %w", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "gocov: %v\n", err)
		}
	}()
	modulePath := modulePath(file)

	// files
	var files []*File
	var missing []string
	for _, f := range blockFiles {
		file, err := NewFile(opts.RootPath, modulePath, f, blocks[f])
		if err != nil {
			return nil, fmt.Errorf("create file report for %s: %w", f, err)
		}

		files = append(files, file)
		if !file.Found {
			missing = append(missing, f)
		}
	}

	// sort
	sort.Slice(files, func(i, j int) bool {
		return files[i].DisplayPath < files[j].DisplayPath
	})

	// update file IDs
	for i := range files {
		files[i].ID = fmt.Sprintf("file-%d", i+1)
	}

	return &Report{
		GeneratedAt:  opts.GeneratedAt,
		RootPath:     opts.RootPath,
		ProfilePath:  opts.ProfilePath,
		OutputPath:   opts.OutputPath,
		Mode:         prof.Mode,
		ModulePath:   modulePath,
		Files:        files,
		Directories:  NewDirectory(files),
		Stats:        NewFileStats(files),
		MissingFiles: missing,
	}, nil
}

func modulePath(r io.Reader) string {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if module, ok := strings.CutPrefix(line, "module "); ok {
			return strings.TrimSpace(module)
		}
	}

	return ""
}
