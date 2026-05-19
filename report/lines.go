package report

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/itsubaki/gocov/profile"
)

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

func newLines(sourceLines []string, blocks []profile.Block) []Line {
	type LineCoverage struct {
		covered bool
		missed  bool
		hits    int64
	}

	coverage := make([]LineCoverage, len(sourceLines)+1)
	for _, block := range blocks {
		start, end := max(1, block.StartLine), block.EndLine
		if block.EndCol == 1 && end > start {
			end--
		}

		end = min(end, len(sourceLines))
		if start > end {
			continue
		}

		for line := start; line <= end; line++ {
			if block.Count > 0 {
				coverage[line].covered = true
				coverage[line].hits += block.Count
				continue
			}

			coverage[line].missed = true
		}
	}

	lines := make([]Line, 0, len(sourceLines))
	for idx, code := range sourceLines {
		lineNo := idx + 1
		cov := coverage[lineNo]
		state := "neutral"

		var hits string
		switch {
		case cov.covered && cov.missed:
			state = "partial"
			hits = formatHits(cov.hits)
		case cov.covered:
			state = "covered"
			hits = formatHits(cov.hits)
		case cov.missed:
			state = "missed"
			hits = "0"
		}

		lines = append(lines, Line{
			Number: lineNo,
			Code:   code,
			Hits:   hits,
			State:  state,
		})
	}

	return lines
}

func formatHits(hits int64) string {
	if hits <= 0 {
		return ""
	}

	if hits < 1000 {
		return strconv.FormatInt(hits, 10)
	}

	if hits < 1000000 {
		return fmt.Sprintf("%.1fk", float64(hits)/1000)
	}

	return fmt.Sprintf("%.1fm", float64(hits)/1000000)
}
