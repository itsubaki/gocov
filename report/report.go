package report

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/itsubaki/gocov/profile"
)

type Options struct {
	RootPath    string
	ProfilePath string
	OutputPath  string
	GeneratedAt time.Time
}

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

	modulePath := modulePath(opts.RootPath)

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

	// stats
	stats := Stats{
		TotalFiles: len(files),
	}

	for _, f := range files {
		stats.Merge(f.Stats)
		stats.Update()
	}

	return &Report{
		GeneratedAt:  opts.GeneratedAt,
		RootPath:     opts.RootPath,
		ProfilePath:  opts.ProfilePath,
		OutputPath:   opts.OutputPath,
		Mode:         prof.Mode,
		ModulePath:   modulePath,
		Stats:        stats,
		Files:        files,
		Directories:  NewDirectory(files),
		MissingFiles: missing,
	}, nil
}

func modulePath(root string) string {
	file, err := os.Open(filepath.Join(root, "go.mod"))
	if err != nil {
		return ""
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "gocov: %v\n", err)
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if module, ok := strings.CutPrefix(line, "module "); ok {
			return strings.TrimSpace(module)
		}
	}

	return ""
}
