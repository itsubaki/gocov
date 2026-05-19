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
	RootPath     string
	ModulePath   string
	ProfilePath  string
	OutputPath   string
	Mode         string
	GeneratedAt  time.Time
	Stats        Stats
	MissingFiles []string
	Files        []*File
	Directories  []*Directory
}

func (r *Report) AddFile(f *File, profile string) {
	if !f.Found {
		r.MissingFiles = append(r.MissingFiles, profile)
	}

	r.Files = append(r.Files, f)
	r.Stats.TotalFiles = len(r.Files)
	r.Stats.Merge(f.Stats)
	r.Stats.Refresh()
}

func (r *Report) AddDirs(dirs map[string]*Directory) {
	for _, dir := range dirs {
		dir.Stats.Refresh()
		r.Directories = append(r.Directories, dir)
	}

	sort.Slice(r.Directories, func(i, j int) bool {
		return r.Directories[i].Name < r.Directories[j].Name
	})

	sort.Slice(r.Files, func(i, j int) bool {
		return r.Files[i].DisplayPath < r.Files[j].DisplayPath
	})

	for i := range r.Files {
		r.Files[i].ID = fmt.Sprintf("file-%d", i+1)
	}
}

func New(prof *profile.Profile, opts Options) (*Report, error) {
	rep := &Report{
		GeneratedAt: opts.GeneratedAt,
		RootPath:    opts.RootPath,
		ProfilePath: opts.ProfilePath,
		OutputPath:  opts.OutputPath,
		Mode:        prof.Mode,
		ModulePath:  modulePath(opts.RootPath),
	}

	blocks := make(map[string][]*profile.Block)
	for _, b := range prof.Blocks {
		blocks[b.File] = append(blocks[b.File], b)
	}

	files := make([]string, 0, len(blocks))
	for b := range blocks {
		files = append(files, b)
	}
	sort.Strings(files)

	// add files
	for i, f := range files {
		file, err := NewFile(opts.RootPath, rep.ModulePath, f, i, blocks[f])
		if err != nil {
			return nil, fmt.Errorf("create file report for %s: %w", f, err)
		}

		rep.AddFile(file, f)
	}

	// add directories
	rep.AddDirs(NewDirectory(rep.Files))
	return rep, nil
}

func modulePath(root string) string {
	file, err := os.Open(filepath.Join(root, "go.mod"))
	if err != nil {
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if module, ok := strings.CutPrefix(line, "module "); ok {
			return strings.TrimSpace(module)
		}
	}

	return ""
}
