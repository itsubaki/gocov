package report

import (
	"github.com/itsubaki/gocov/profile"
)

type Stats struct {
	TotalStatements   int
	CoveredStatements int
	TotalLines        int
	CoveredLines      int
	PartialLines      int
	MissedLines       int
	TotalFiles        int
	Percent           float64
	Status            string
}

func NewFileStats(files []*File) Stats {
	var s Stats
	for _, f := range files {
		s = Merge(s, f.Stats)
	}

	s.TotalFiles = len(files)
	return s
}

func NewLineStats(lines []Line, blocks []*profile.Block) Stats {
	var s Stats
	for _, v := range blocks {
		s.TotalStatements += v.Statements
		if v.Count > 0 {
			s.CoveredStatements += v.Statements
		}
	}

	for _, line := range lines {
		switch line.State {
		case "covered":
			s.CoveredLines++
			s.TotalLines++
		case "partial":
			s.PartialLines++
			s.TotalLines++
		case "missed":
			s.MissedLines++
			s.TotalLines++
		}
	}

	s.update()
	return s
}

func Merge(a, b Stats) Stats {
	a.TotalStatements += b.TotalStatements
	a.CoveredStatements += b.CoveredStatements
	a.TotalLines += b.TotalLines
	a.CoveredLines += b.CoveredLines
	a.PartialLines += b.PartialLines
	a.MissedLines += b.MissedLines
	a.TotalFiles += b.TotalFiles

	a.update()
	return a
}

func (s *Stats) update() {
	s.Percent = 100
	if s.TotalStatements > 0 {
		a := float64(s.CoveredStatements)
		b := float64(s.TotalStatements)
		s.Percent = a / b * 100
	}

	switch v := s.Percent; {
	case v >= 80:
		s.Status = "high"
	case v >= 50:
		s.Status = "medium"
	default:
		s.Status = "low"
	}
}
