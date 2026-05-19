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

	cov := make([]LineCoverage, len(sourceLines)+1)
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
				cov[line].covered = true
				cov[line].hits += v.Count
				continue
			}

			cov[line].missed = true
		}
	}

	lines := make([]Line, 0, len(sourceLines))
	for i, code := range sourceLines {
		lineNo, state, hits := i+1, "neutral", "0"

		switch v := cov[lineNo]; {
		case v.covered && v.missed:
			state = "partial"
			hits = formatHits(v.hits)
		case v.covered:
			state = "covered"
			hits = formatHits(v.hits)
		case v.missed:
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
