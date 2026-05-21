package report_test

import (
	"fmt"
	"testing"

	"github.com/itsubaki/gocov/report"
)

func ExampleMerge() {
	a := &report.Stats{
		TotalStatements:   10,
		CoveredStatements: 5,
		TotalLines:        20,
		CoveredLines:      10,
	}

	b := &report.Stats{
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
		s    *report.Stats
		want int
	}{
		{
			s: &report.Stats{
				TotalLines:      10,
				TotalStatements: 5,
			},
			want: 10,
		},
		{
			s: &report.Stats{
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
