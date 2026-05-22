package profile

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Profile struct {
	Mode   string   // coverage mode (e.g., "set", "count", "atomic")
	Blocks []*Block // coverage blocks
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

	// read mode header
	var mode string
	if scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			return nil, fmt.Errorf("missing mode header %q", path)
		}

		md, ok := strings.CutPrefix(line, "mode:")
		if !ok {
			return nil, fmt.Errorf("invalid mode header %q", line)
		}

		mode = strings.TrimSpace(md)
		if mode == "" {
			return nil, fmt.Errorf("empty mode header %q", line)
		}
	}

	// read blocks
	var blocks []*Block
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		block, err := parseLine(line)
		if err != nil {
			return nil, fmt.Errorf("parse line %q: %w", line, err)
		}

		blocks = append(blocks, block)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("read coverage profile: %w", err)
	}

	return &Profile{
		Mode:   mode,
		Blocks: blocks,
	}, nil
}
