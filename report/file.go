package report

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/itsubaki/gocov/profile"
)

type File struct {
	ID          string
	DisplayPath string
	ProfilePath string
	SourcePath  string
	Directory   string
	Found       bool
	Blocks      int
	Stats       Stats
	Lines       []Line
}

func NewFile(rootPath, modulePath, profileFile string, idx int, blocks []*profile.Block) (*File, error) {
	sourcePath, found := sourcePath(rootPath, modulePath, profileFile)
	displayPath := displayPath(rootPath, modulePath, profileFile, sourcePath, found)

	dirName := filepath.ToSlash(filepath.Dir(displayPath))
	if dirName == "." {
		dirName = "root"
	}

	var lines []Line
	if found {
		srcLines, err := sourceLines(sourcePath)
		if err != nil {
			return nil, fmt.Errorf("read source %s: %w", sourcePath, err)
		}

		lines = NewLines(srcLines, blocks)
	}

	var stats Stats
	for _, v := range blocks {
		stats.TotalStatements += v.Statements
		if v.Count > 0 {
			stats.CoveredStatements += v.Statements
		}
	}
	stats.Increment(lines)
	stats.Refresh()

	return &File{
		ID:          fmt.Sprintf("file-%d", idx+1),
		DisplayPath: displayPath,
		ProfilePath: profileFile,
		SourcePath:  sourcePath,
		Directory:   dirName,
		Found:       found,
		Blocks:      len(blocks),
		Lines:       lines,
		Stats:       stats,
	}, nil
}

func sourcePath(root, modulePath, profileFile string) (string, bool) {
	var paths []string
	if filepath.IsAbs(profileFile) {
		paths = append(paths, profileFile)
	}
	slash := filepath.ToSlash(profileFile)

	paths = append(paths, filepath.Join(root, filepath.FromSlash(slash)))
	if modulePath != "" {
		if rel, ok := strings.CutPrefix(slash, modulePath+"/"); ok {
			p := filepath.Join(root, filepath.FromSlash(rel))
			paths = append(paths, p)
		}

		if idx := strings.Index(slash, modulePath+"/"); idx > 0 {
			rel := strings.TrimPrefix(slash[idx:], modulePath+"/")
			p := filepath.Join(root, filepath.FromSlash(rel))
			paths = append(paths, p)
		}
	}

	for _, v := range paths {
		info, err := os.Stat(v)
		if err != nil || info.IsDir() {
			continue
		}

		abs, err := filepath.Abs(v)
		if err != nil {
			return v, true
		}

		return abs, true
	}

	return "", false
}

func displayPath(root, modulePath, profileFile, sourcePath string, found bool) string {
	if found {
		rel, err := filepath.Rel(root, sourcePath)
		if err == nil && !strings.HasPrefix(rel, "..") {
			return filepath.ToSlash(rel)
		}
	}

	slash := filepath.ToSlash(profileFile)
	if modulePath == "" {
		return slash
	}

	if rel, ok := strings.CutPrefix(slash, modulePath+"/"); ok {
		return rel
	}

	return slash
}

func sourceLines(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// normalize line endings to \n
	text := strings.ReplaceAll(string(data), "\r\n", "\n")
	text = strings.ReplaceAll(text, "\r", "\n")
	lines := strings.Split(text, "\n")

	// remove trailing empty line
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	return lines, nil
}
