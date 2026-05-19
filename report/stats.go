package report

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

func (s *Stats) Weight() int {
	if s.TotalLines > 0 {
		return s.TotalLines
	}

	return s.TotalStatements
}

func (s *Stats) Add(a Stats) {
	s.TotalStatements += a.TotalStatements
	s.CoveredStatements += a.CoveredStatements
	s.TotalLines += a.TotalLines
	s.CoveredLines += a.CoveredLines
	s.PartialLines += a.PartialLines
	s.MissedLines += a.MissedLines
	s.TotalFiles += a.TotalFiles
}

func (s *Stats) Increment(lines []Line) {
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
}

func (s *Stats) Refresh() {
	div := func(a, b int) float64 {
		if b == 0 {
			return 0
		}

		return float64(a) / float64(b)
	}

	s.Percent = 100
	if s.TotalStatements > 0 {
		s.Percent = div(s.CoveredStatements, s.TotalStatements) * 100
	}

	switch {
	case s.Percent >= 80:
		s.Status = "high"
	case s.Percent >= 50:
		s.Status = "medium"
	default:
		s.Status = "low"
	}
}
