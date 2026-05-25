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
	ProfilePath string
	SourcePath  string
	DisplayPath string
	Directory   string
	Found       bool
	Blocks      int
	Stats       Stats
	Lines       []Line
}

func NewFile(rootPath, modulePath, profileFile string, blocks []*profile.Block) (*File, error) {
	var lines []Line
	srcPath, found := sourcePath(rootPath, modulePath, profileFile)
	if found {
		data, err := os.ReadFile(srcPath)
		if err != nil {
			return nil, fmt.Errorf("read source file %s: %w", srcPath, err)
		}

		lines = NewLines(sourceLines(string(data)), blocks)
	}

	dispPath := func() string {
		if rel, ok := relativeSourcePath(rootPath, srcPath); ok {
			return rel
		}

		return displayPath(modulePath, profileFile)
	}()

	dirName := func() string {
		name := filepath.ToSlash(filepath.Dir(dispPath))
		if name == "." {
			return "root"
		}

		return name
	}()

	return &File{
		ProfilePath: profileFile,
		SourcePath:  srcPath,
		DisplayPath: dispPath,
		Directory:   dirName,
		Found:       found,
		Blocks:      len(blocks),
		Lines:       lines,
		Stats:       NewLineStats(lines, blocks),
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

func relativeSourcePath(root, sourcePath string) (string, bool) {
	if sourcePath == "" {
		return "", false
	}

	rel, err := filepath.Rel(root, sourcePath)
	if err != nil || strings.HasPrefix(rel, "..") {
		return "", false
	}

	return filepath.ToSlash(rel), true
}

func displayPath(modulePath, profileFile string) string {
	slash := filepath.ToSlash(profileFile)
	if modulePath == "" {
		return slash
	}

	if rel, ok := strings.CutPrefix(slash, modulePath+"/"); ok {
		return rel
	}

	return slash
}

func sourceLines(data string) []string {
	// normalize line endings to \n
	text := strings.ReplaceAll(data, "\r\n", "\n")
	text = strings.ReplaceAll(text, "\r", "\n")
	lines := strings.Split(text, "\n")

	// remove trailing empty line
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	return lines
}
