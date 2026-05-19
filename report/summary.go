package report

type Summary struct {
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

func (s *Summary) Weight() int {
	if s.TotalLines > 0 {
		return s.TotalLines
	}

	return s.TotalStatements
}

func (s *Summary) Add(another Summary) {
	s.TotalStatements += another.TotalStatements
	s.CoveredStatements += another.CoveredStatements
	s.TotalLines += another.TotalLines
	s.CoveredLines += another.CoveredLines
	s.PartialLines += another.PartialLines
	s.MissedLines += another.MissedLines
	s.TotalFiles += another.TotalFiles
}

func (s *Summary) Increment(lines []Line) {
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

func (s *Summary) Refresh() {
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
