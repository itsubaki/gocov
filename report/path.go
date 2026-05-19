package report

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

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
