package report_test

import (
	"testing"

	"github.com/itsubaki/gocov/report"
)

func TestNewDirectory(t *testing.T) {
	cases := []struct {
		files []*report.File
		want  []string
	}{
		{
			files: []*report.File{
				{Directory: "b", Stats: &report.Stats{CoveredLines: 1}},
				{Directory: "a", Stats: &report.Stats{CoveredLines: 2}},
				{Directory: "b", Stats: &report.Stats{CoveredLines: 3}},
			},
			want: []string{"a", "b"},
		},
		{
			files: []*report.File{
				{Directory: "x", Stats: &report.Stats{CoveredLines: 1}},
				{Directory: "y", Stats: &report.Stats{CoveredLines: 2}},
			},
			want: []string{"x", "y"},
		},
	}

	for i, c := range cases {
		got := report.NewDirectory(c.files)
		if len(got) != len(c.want) {
			t.Errorf("case %d: len = %d want %d", i, len(got), len(c.want))
		}

		for j := range got {
			if got[j].Name != c.want[j] {
				t.Errorf("case %d: index %d got %s want %s", i, j, got[j].Name, c.want[j])
			}
		}
	}
}
