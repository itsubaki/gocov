package report_test

import (
	"fmt"
	"testing"

	"github.com/itsubaki/gocov/profile"
	"github.com/itsubaki/gocov/report"
)

func TestNewFileStats(t *testing.T) {
	cases := []struct {
		files []*report.File
		want  report.Stats
	}{
		{
			files: []*report.File{
				{
					Lines: []report.Line{
						{State: "covered"},
						{State: "missed"},
						{State: "partial"},
						{State: "neutral"},
					},
					Stats: report.Stats{
						TotalStatements:   4,
						CoveredStatements: 1,
						TotalLines:        4,
						CoveredLines:      1,
						PartialLines:      1,
						MissedLines:       1,
					},
				},
			},
			want: report.Stats{
				TotalStatements:   4,
				CoveredStatements: 1,
				TotalLines:        4,
				CoveredLines:      1,
				PartialLines:      1,
				MissedLines:       1,
				TotalFiles:        1,
				Percent:           25.0,
				Status:            "low",
			},
		},
	}

	for _, c := range cases {
		got := report.NewFileStats(c.files)
		if got.TotalStatements != c.want.TotalStatements ||
			got.CoveredStatements != c.want.CoveredStatements ||
			got.TotalLines != c.want.TotalLines ||
			got.CoveredLines != c.want.CoveredLines ||
			got.PartialLines != c.want.PartialLines ||
			got.MissedLines != c.want.MissedLines ||
			got.TotalFiles != c.want.TotalFiles ||
			got.Percent != c.want.Percent ||
			got.Status != c.want.Status {
			t.Errorf("NewFileStats() = %v, want %v", got, c.want)
		}
	}
}

func TestNewLineStats(t *testing.T) {
	type Want struct {
		coveredLines      int
		partialLines      int
		missedLines       int
		totalLines        int
		coveredStatements int
		totalStatements   int
	}

	cases := []struct {
		lines  []report.Line
		blocks []*profile.Block
		want   Want
	}{
		{
			lines: []report.Line{
				{State: "covered"},
				{State: "missed"},
			},
			blocks: []*profile.Block{
				{Statements: 3, Count: 1},
				{Statements: 2, Count: 0},
			},
			want: Want{
				coveredLines:      1,
				partialLines:      0,
				missedLines:       1,
				totalLines:        2,
				coveredStatements: 3,
				totalStatements:   5,
			},
		},
		{
			lines: []report.Line{
				{State: "partial"},
				{State: "covered"},
				{State: "missed"},
			},
			blocks: []*profile.Block{
				{Statements: 10, Count: 5},
			},
			want: Want{
				coveredLines:      1,
				partialLines:      1,
				missedLines:       1,
				totalLines:        3,
				coveredStatements: 10,
				totalStatements:   10,
			},
		},
	}

	for i, c := range cases {
		got := report.NewLineStats(c.lines, c.blocks)
		if got.CoveredLines != c.want.coveredLines ||
			got.PartialLines != c.want.partialLines ||
			got.MissedLines != c.want.missedLines ||
			got.TotalLines != c.want.totalLines ||
			got.CoveredStatements != c.want.coveredStatements ||
			got.TotalStatements != c.want.totalStatements {
			t.Errorf("case %d: NewLineStats() = %v, want %v", i, got, c.want)
		}
	}
}

func ExampleMerge() {
	a := report.Stats{
		TotalStatements:   10,
		CoveredStatements: 5,
		TotalLines:        20,
		CoveredLines:      10,
	}

	b := report.Stats{
		TotalStatements:   20,
		CoveredStatements: 15,
		TotalLines:        30,
		CoveredLines:      25,
	}

	c := report.Merge(a, b)
	fmt.Println(c.TotalStatements)
	fmt.Println(c.CoveredStatements)
	fmt.Println(c.TotalLines)
	fmt.Println(c.CoveredLines)
	fmt.Println(c.Percent)
	fmt.Println(c.Status)

	// Output:
	// 30
	// 20
	// 50
	// 35
	// 66.66666666666666
	// medium

}

func TestStats_Weight(t *testing.T) {
	cases := []struct {
		s    report.Stats
		want int
	}{
		{
			s: report.Stats{
				TotalLines:      10,
				TotalStatements: 5,
			},
			want: 10,
		},
		{
			s: report.Stats{
				TotalLines:      0,
				TotalStatements: 0,
			},
			want: 0,
		},
	}

	for _, c := range cases {
		if got := c.s.Weight(); got != c.want {
			t.Errorf("Weight() = %d, want %d", got, c.want)
		}
	}
}

func Test_div(t *testing.T) {
	cases := []struct {
		a, b int
		want float64
	}{
		{a: 1, b: 2, want: 0.5},
		{a: 1, b: 0, want: 0},
	}

	for _, c := range cases {
		if got := report.Div(c.a, c.b); got != c.want {
			t.Errorf("div(%d, %d) = %f, want %f", c.a, c.b, got, c.want)
		}
	}
}
