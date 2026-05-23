package directory_test

import (
	"fmt"
	"testing"

	"github.com/itsubaki/gocov/render/directory"
	"github.com/itsubaki/gocov/report"
)

func ExampleNew() {
	dirs := directory.New(&report.Report{Directories: []*report.Directory{
		{
			Name: "root",
			Stats: report.Stats{
				CoveredStatements: 10,
				TotalStatements:   100,
			},
		},
		{
			Name: "dir1",
			Stats: report.Stats{
				CoveredStatements: 25,
				TotalStatements:   120,
			},
		},
		{
			Name: "dir1/subdir1",
			Stats: report.Stats{
				CoveredStatements: 25,
				TotalStatements:   50,
			},
		},
		{
			Name: "dir1/subdir2",
			Stats: report.Stats{
				CoveredStatements: 30,
				TotalStatements:   30,
			},
		},
	}})

	for _, d := range dirs {
		fmt.Println(d.Name, d.Depth, d.Coverage, d.Statements, d.Share, d.Color)
	}

	// Output:
	// root 0 30.0% 300 100.0% #c53026
	// dir1 1 40.0% 200 66.7% #c53426
	// dir1/subdir1 2 50.0% 50 16.7% #c53726
	// dir1/subdir2 2 100.0% 30 10.0% #2f9d5d
}

func Test_donutSegmentPath(t *testing.T) {
	cases := []struct {
		start     float64
		end       float64
		depth     int
		ringCount int
		want      string
	}{
		{
			start:     0,
			end:       90,
			depth:     0,
			ringCount: 1,
			want:      "M 1.0000 0.0000 A 1.0000 1.0000 0 0 1 0.0000 1.0000 L 0.0000 0.3400 A 0.3400 0.3400 0 0 0 0.3400 0.0000 Z",
		},
		{
			start:     -90,
			end:       270,
			depth:     1,
			ringCount: 2,
			want:      "M 0 -1.0000 A 1.0000 1.0000 0 1 1 0 1.0000 A 1.0000 1.0000 0 1 1 0 -1.0000 M 0 -0.6700 A 0.6700 0.6700 0 1 0 0 0.6700 A 0.6700 0.6700 0 1 0 0 -0.6700 Z",
		},
		{
			start:     0,
			end:       200,
			depth:     0,
			ringCount: 1,
			want:      "M 1.0000 0.0000 A 1.0000 1.0000 0 1 1 -0.9397 -0.3420 L -0.3195 -0.1163 A 0.3400 0.3400 0 1 0 0.3400 0.0000 Z",
		},
		{
			start:     100,
			end:       100,
			depth:     0,
			ringCount: 1,
			want:      "",
		},
	}

	for _, c := range cases {
		if got := directory.DonutSegmentPath(c.start, c.end, c.depth, c.ringCount); got != c.want {
			t.Errorf("donutSegmentPath(%v, %v, %v, %v) = %v, want %v", c.start, c.end, c.depth, c.ringCount, got, c.want)
		}
	}
}

func Test_ringRadii(t *testing.T) {
	cases := []struct {
		depth int
		count int
		want  [2]float64
	}{
		{
			depth: 0,
			count: 0,
			want:  [2]float64{0.34, 1},
		},
		{
			depth: 0,
			count: 2,
			want:  [2]float64{0.34, 0.67},
		},
		{
			depth: 1,
			count: 2,
			want:  [2]float64{0.67, 1},
		},
	}

	for _, c := range cases {
		inner, outer := directory.RingRadii(c.depth, c.count)
		if inner != c.want[0] {
			t.Errorf("ringRadii(%v, %v) inner = %v, want %v", c.depth, c.count, inner, c.want[0])
		}

		if outer != c.want[1] {
			t.Errorf("ringRadii(%v, %v) outer = %v, want %v", c.depth, c.count, outer, c.want[1])
		}
	}
}
