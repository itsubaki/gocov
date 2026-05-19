package profile

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Profile struct {
	Mode   string
	Blocks []Block
}

type Block struct {
	File       string
	StartLine  int
	StartCol   int
	EndLine    int
	EndCol     int
	Statements int
	Count      int64
}

func Parse(path string) (Profile, error) {
	file, err := os.Open(path)
	if err != nil {
		return Profile{}, fmt.Errorf("open coverage profile: %w", err)
	}
	defer file.Close()

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
				return Profile{}, fmt.Errorf("line %d: expected mode header", lineNo)
			}

			prof.Mode = strings.TrimSpace(mode)
			continue
		}

		block, err := parseLine(line)
		if err != nil {
			return Profile{}, fmt.Errorf("line %d: %w", lineNo, err)
		}

		prof.Blocks = append(prof.Blocks, block)
	}

	if err := scanner.Err(); err != nil {
		return Profile{}, fmt.Errorf("read coverage profile: %w", err)
	}

	if prof.Mode == "" {
		return Profile{}, fmt.Errorf("missing mode header %q", path)
	}

	return prof, nil
}

func parseLine(line string) (Block, error) {
	fields := strings.Fields(line)
	if len(fields) != 3 {
		return Block{}, fmt.Errorf("invalid line: %q", line)
	}

	location := fields[0]
	colon := strings.LastIndex(location, ":")
	if colon < 1 || colon == len(location)-1 {
		return Block{}, fmt.Errorf("invalid location %q", location)
	}

	startEnd := location[colon+1:]
	points := strings.Split(startEnd, ",")
	if len(points) != 2 {
		return Block{}, fmt.Errorf("invalid range %q", startEnd)
	}

	startLine, startCol, err := parsePoint(points[0])
	if err != nil {
		return Block{}, fmt.Errorf("invalid start point: %w", err)
	}

	endLine, endCol, err := parsePoint(points[1])
	if err != nil {
		return Block{}, fmt.Errorf("invalid end point: %w", err)
	}

	statements, err := strconv.Atoi(fields[1])
	if err != nil {
		return Block{}, fmt.Errorf("invalid statement count %q", fields[1])
	}

	count, err := strconv.ParseInt(fields[2], 10, 64)
	if err != nil {
		return Block{}, fmt.Errorf("invalid execution count %q", fields[2])
	}

	return Block{
		File:       filepath.ToSlash(location[:colon]),
		StartLine:  startLine,
		StartCol:   startCol,
		EndLine:    endLine,
		EndCol:     endCol,
		Statements: statements,
		Count:      count,
	}, nil
}

func parsePoint(point string) (int, int, error) {
	lineCol := strings.Split(point, ".")
	if len(lineCol) != 2 {
		return 0, 0, fmt.Errorf("invalid point: %q", point)
	}

	line, err := strconv.Atoi(lineCol[0])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid line: %q", lineCol[0])
	}

	col, err := strconv.Atoi(lineCol[1])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid column: %q", lineCol[1])
	}

	return line, col, nil
}
