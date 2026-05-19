package profile

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
)

type Block struct {
	File       string
	StartLine  int
	StartCol   int
	EndLine    int
	EndCol     int
	Statements int
	Count      int64
}

func parseLine(line string) (*Block, error) {
	fields := strings.Fields(line)
	if len(fields) != 3 {
		return nil, fmt.Errorf("invalid line: %q", line)
	}

	location := fields[0]
	colon := strings.LastIndex(location, ":")
	if colon < 1 || colon == len(location)-1 {
		return nil, fmt.Errorf("invalid location %q", location)
	}

	startEnd := location[colon+1:]
	points := strings.Split(startEnd, ",")
	if len(points) != 2 {
		return nil, fmt.Errorf("invalid range %q", startEnd)
	}

	startLine, startCol, err := parsePoint(points[0])
	if err != nil {
		return nil, fmt.Errorf("invalid start point: %w", err)
	}

	endLine, endCol, err := parsePoint(points[1])
	if err != nil {
		return nil, fmt.Errorf("invalid end point: %w", err)
	}

	statements, err := strconv.Atoi(fields[1])
	if err != nil {
		return nil, fmt.Errorf("invalid statement count %q", fields[1])
	}

	count, err := strconv.ParseInt(fields[2], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid execution count %q", fields[2])
	}

	return &Block{
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
