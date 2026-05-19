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

type Report struct {
	RootPath     string
	ModulePath   string
	ProfilePath  string
	OutputPath   string
	Mode         string
	GeneratedAt  time.Time
	Stats        Stats
	Files        []File
	Directories  []Directory
	MissingFiles []string
}

func (r *Report) AddFile(file File, profileFile string) {
	if !file.Found {
		r.MissingFiles = append(r.MissingFiles, profileFile)
	}

	r.Files = append(r.Files, file)
	r.Stats.TotalFiles = len(r.Files)
	r.Stats.Add(file.Stats)
	r.Stats.Refresh()
}

func (r *Report) AddDirs(dirs map[string]*Directory) {
	for _, dir := range dirs {
		dir.Stats.Refresh()
		r.Directories = append(r.Directories, *dir)
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

type Directory struct {
	Name  string
	Stats Stats
}

type Line struct {
	Number int
	Code   string
	Hits   string
	State  string
}

type Options struct {
	RootPath    string
	ProfilePath string
	OutputPath  string
	GeneratedAt time.Time
}

func New(prof profile.Profile, opts Options) (Report, error) {
	rep := Report{
		GeneratedAt: opts.GeneratedAt,
		RootPath:    opts.RootPath,
		ProfilePath: opts.ProfilePath,
		OutputPath:  opts.OutputPath,
		Mode:        prof.Mode,
		ModulePath:  modulePath(opts.RootPath),
	}

	blocks := make(map[string][]profile.Block)
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
			return Report{}, fmt.Errorf("create file report for %s: %w", f, err)
		}

		rep.AddFile(file, f)
	}

	// add directories
	dirs := make(map[string]*Directory)
	for _, f := range rep.Files {
		dir, ok := dirs[f.Directory]
		if !ok {
			dirs[f.Directory] = &Directory{
				Name:  f.Directory,
				Stats: f.Stats,
			}

			continue
		}

		dir.Stats.Add(f.Stats)
	}

	rep.AddDirs(dirs)
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
