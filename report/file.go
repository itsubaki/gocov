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

func (f *File) SetLines(lines []Line) {
	f.Lines = lines
	f.Stats.Increment(lines)
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
		file.Stats.TotalStatements += block.Statements
		if block.Count > 0 {
			file.Stats.CoveredStatements += block.Statements
		}
	}

	if found {
		lines, err := sourceLines(sourcePath)
		if err != nil {
			return File{}, fmt.Errorf("read source %s: %w", sourcePath, err)
		}

		file.SetLines(NewLines(lines, blocks))
	}

	file.Stats.Refresh()
	return file, nil
}

func sourcePath(root, modulePath, profileFile string) (string, bool) {
	var candidates []string
	if filepath.IsAbs(profileFile) {
		candidates = append(candidates, profileFile)
	}

	slash := filepath.ToSlash(profileFile)
	candidates = append(candidates, filepath.Join(root, filepath.FromSlash(slash)))
	if modulePath != "" {
		if rel, ok := strings.CutPrefix(slash, modulePath+"/"); ok {
			candidates = append(candidates, filepath.Join(root, filepath.FromSlash(rel)))
		}

		if idx := strings.Index(slash, modulePath+"/"); idx > 0 {
			rel := strings.TrimPrefix(slash[idx:], modulePath+"/")
			candidates = append(candidates, filepath.Join(root, filepath.FromSlash(rel)))
		}
	}

	for _, candidate := range candidates {
		if info, err := os.Stat(candidate); err == nil && !info.IsDir() {
			abs, err := filepath.Abs(candidate)
			if err != nil {
				return candidate, true
			}

			return abs, true
		}
	}

	return "", false
}

func displayPath(root, modulePath, profileFile, sourcePath string, found bool) string {
	if found {
		if rel, err := filepath.Rel(root, sourcePath); err == nil && !strings.HasPrefix(rel, "..") {
			return filepath.ToSlash(rel)
		}
	}

	slash := filepath.ToSlash(profileFile)
	if modulePath != "" {
		if rel, ok := strings.CutPrefix(slash, modulePath+"/"); ok {
			return rel
		}
	}

	return slash
}

func sourceLines(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	text := strings.ReplaceAll(string(data), "\r\n", "\n")
	text = strings.ReplaceAll(text, "\r", "\n")
	lines := strings.Split(text, "\n")
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	return lines, nil
}
