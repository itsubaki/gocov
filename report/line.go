package report

import (
	"fmt"
	"strconv"

	"github.com/itsubaki/gocov/profile"
)

type Line struct {
	Number int
	Code   string
	Hits   string
	State  string
}

func NewLines(sourceLines []string, blocks []*profile.Block) []Line {
	type LineCoverage struct {
		covered bool
		missed  bool
		hits    int64
	}

	coverage := make([]LineCoverage, len(sourceLines)+1)
	for _, v := range blocks {
		start, end := max(1, v.StartLine), v.EndLine
		if v.EndCol == 1 && end > start {
			end--
		}

		end = min(end, len(sourceLines))
		if start > end {
			continue
		}

		for line := start; line <= end; line++ {
			if v.Count > 0 {
				coverage[line].covered = true
				coverage[line].hits += v.Count
				continue
			}

			coverage[line].missed = true
		}
	}

	lines := make([]Line, 0, len(sourceLines))
	for i, code := range sourceLines {
		lineNo, state, hits := i+1, "neutral", "0"

		switch cov := coverage[lineNo]; {
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
