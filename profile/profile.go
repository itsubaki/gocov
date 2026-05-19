package profile

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Profile struct {
	Mode   string
	Blocks []*Block
}

func Parse(path string) (*Profile, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open coverage profile: %w", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "gocov: %v\n", err)
		}
	}()

	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 1024), 1024*1024)

	var prof Profile
	var lineNo int
	for scanner.Scan() {
		lineNo++
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		if prof.Mode == "" {
			mode, ok := strings.CutPrefix(line, "mode:")
			if !ok {
				return nil, fmt.Errorf("line %d: expected mode header", lineNo)
			}

			prof.Mode = strings.TrimSpace(mode)
			continue
		}

		block, err := parseLine(line)
		if err != nil {
			return nil, fmt.Errorf("line %d: %w", lineNo, err)
		}

		prof.Blocks = append(prof.Blocks, block)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("read coverage profile: %w", err)
	}

	if prof.Mode == "" {
		return nil, fmt.Errorf("missing mode header %q", path)
	}

	return &prof, nil
}
