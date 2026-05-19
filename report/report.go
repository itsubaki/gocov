package report

import (
	"fmt"
	"path/filepath"
	"sort"
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
	Summary      Summary
	Files        []File
	Directories  []Directory
	MissingFiles []string
}

type File struct {
	ID          string
	DisplayPath string
	ProfilePath string
	SourcePath  string
	Directory   string
	Found       bool
	Blocks      int
	Summary     Summary
	Lines       []Line
}

func NewFile(rootPath, modulePath, profileFile string, idx int, blocks []profile.Block) (File, error) {
	sourcePath, found := sourcePath(rootPath, modulePath, profileFile)
	displayPath := displayPath(rootPath, modulePath, profileFile, sourcePath, found)
	dirName := filepath.ToSlash(filepath.Dir(displayPath))
	if dirName == "." {
		dirName = "root"
	}

	file := File{
		ID:          fmt.Sprintf("file-%d", idx+1),
		DisplayPath: displayPath,
		ProfilePath: profileFile,
		SourcePath:  sourcePath,
		Directory:   dirName,
		Found:       found,
		Blocks:      len(blocks),
	}

	for _, block := range blocks {
		file.Summary.TotalStatements += block.Statements
		if block.Count > 0 {
			file.Summary.CoveredStatements += block.Statements
		}
	}

	if found {
		lines, err := sourceLines(sourcePath)
		if err != nil {
			return File{}, fmt.Errorf("read source %s: %w", sourcePath, err)
		}

		file.Lines = newLines(lines, blocks)
		file.Summary.Count(file.Lines)
	}

	file.Summary.Refresh()
	return file, nil
}

type Directory struct {
	Name    string
	Summary Summary
}

type Line struct {
	Number int
	Code   string
	Hits   string
	State  string
}

type Summary struct {
	TotalStatements   int
	CoveredStatements int
	TotalLines        int
	CoveredLines      int
	PartialLines      int
	MissedLines       int
	TotalFiles        int
	Percent           float64
	Status            string
}

func (s *Summary) Weight() int {
	if s.TotalLines > 0 {
		return s.TotalLines
	}

	return s.TotalStatements
}

func (s *Summary) Add(another Summary) {
	s.TotalStatements += another.TotalStatements
	s.CoveredStatements += another.CoveredStatements
	s.TotalLines += another.TotalLines
	s.CoveredLines += another.CoveredLines
	s.PartialLines += another.PartialLines
	s.MissedLines += another.MissedLines
	s.TotalFiles += another.TotalFiles
}

func (s *Summary) Refresh() {
	div := func(a, b int) float64 {
		if b == 0 {
			return 0
		}

		return float64(a) / float64(b)
	}

	s.Percent = 100
	if s.TotalStatements > 0 {
		s.Percent = div(s.CoveredStatements, s.TotalStatements) * 100
	}

	switch {
	case s.Percent >= 80:
		s.Status = "high"
	case s.Percent >= 50:
		s.Status = "medium"
	default:
		s.Status = "low"
	}
}

func (s *Summary) Count(lines []Line) {
	for _, line := range lines {
		switch line.State {
		case "covered":
			s.CoveredLines++
			s.TotalLines++
		case "partial":
			s.PartialLines++
			s.TotalLines++
		case "missed":
			s.MissedLines++
			s.TotalLines++
		}
	}
}

type Options struct {
	RootPath    string
	ProfilePath string
	OutputPath  string
	GeneratedAt time.Time
}

func New(prof profile.Profile, opts Options) (Report, error) {
	modulePath := modulePath(opts.RootPath)
	blocksByFile := make(map[string][]profile.Block)
	for _, block := range prof.Blocks {
		blocksByFile[block.File] = append(blocksByFile[block.File], block)
	}

	files := make([]string, 0, len(blocksByFile))
	for f := range blocksByFile {
		files = append(files, f)
	}
	sort.Strings(files)

	rep := Report{
		GeneratedAt: opts.GeneratedAt,
		RootPath:    opts.RootPath,
		ProfilePath: opts.ProfilePath,
		OutputPath:  opts.OutputPath,
		Mode:        prof.Mode,
		ModulePath:  modulePath,
	}

	dirs := make(map[string]*Directory)
	for idx, profileFile := range files {
		blocks := blocksByFile[profileFile]
		file, err := NewFile(opts.RootPath, modulePath, profileFile, idx, blocks)
		if err != nil {
			return Report{}, fmt.Errorf("create file report for %s: %w", profileFile, err)
		}

		if !file.Found {
			rep.MissingFiles = append(rep.MissingFiles, profileFile)
		}

		rep.Summary.Add(file.Summary)
		rep.Files = append(rep.Files, file)

		dir, ok := dirs[file.Directory]
		if !ok {
			dirs[file.Directory] = &Directory{
				Name:    file.Directory,
				Summary: file.Summary,
			}

			continue
		}

		dir.Summary.Add(file.Summary)
	}

	rep.Summary.TotalFiles = len(rep.Files)
	rep.Summary.Refresh()

	for _, dir := range dirs {
		dir.Summary.Refresh()
		rep.Directories = append(rep.Directories, *dir)
	}

	sort.Slice(rep.Directories, func(i, j int) bool {
		return rep.Directories[i].Name < rep.Directories[j].Name
	})

	sort.Slice(rep.Files, func(i, j int) bool {
		return rep.Files[i].DisplayPath < rep.Files[j].DisplayPath
	})

	for i := range rep.Files {
		rep.Files[i].ID = fmt.Sprintf("file-%d", i+1)
	}

	return rep, nil
}
