package report_test

import (
	"testing"

	"github.com/itsubaki/gocov/profile"
	"github.com/itsubaki/gocov/report"
)

func TestNewLines(t *testing.T) {
	cases := []struct {
		source []string
		blocks []*profile.Block
		want   []report.Line
	}{
		{
			source: []string{
				"package main",
				"",
				"func main() {",
				"    println(\"Hello, world!\")",
				"}",
			},
			blocks: []*profile.Block{
				{StartLine: 1, EndLine: 1, Count: 1},
				{StartLine: 3, EndLine: 4, Count: 0},
			},
			want: []report.Line{
				{Number: 1, Code: "package main", Hits: "1", State: "covered"},
				{Number: 2, Code: "", Hits: "0", State: "neutral"},
				{Number: 3, Code: "func main() {", Hits: "0", State: "missed"},
				{Number: 4, Code: "    println(\"Hello, world!\")", Hits: "0", State: "missed"},
				{Number: 5, Code: "}", Hits: "0", State: "neutral"},
			},
		},
		{
			source: []string{
				"if x > 0 {",
				"    println(\"positive\")",
				"} else {",
				"    println(\"non-positive\")",
				"}",
			},
			blocks: []*profile.Block{
				{StartLine: 1, EndLine: 4, Count: 1},
				{StartLine: 3, EndLine: 4, Count: 0},
			},
			want: []report.Line{
				{Number: 1, Code: "if x > 0 {", Hits: "1", State: "covered"},
				{Number: 2, Code: "    println(\"positive\")", Hits: "1", State: "covered"},
				{Number: 3, Code: "} else {", Hits: "1", State: "partial"},
				{Number: 4, Code: "    println(\"non-positive\")", Hits: "1", State: "partial"},
				{Number: 5, Code: "}", Hits: "0", State: "neutral"},
			},
		},
	}

	for _, c := range cases {
		got := report.NewLines(c.source, c.blocks)
		if len(got) != len(c.want) {
			t.Errorf("got %d lines, want %d lines", len(got), len(c.want))
			continue
		}

		for i := range got {
			if got[i].Number != c.want[i].Number ||
				got[i].Code != c.want[i].Code ||
				got[i].Hits != c.want[i].Hits ||
				got[i].State != c.want[i].State {
				t.Errorf("line %d: got %+v, want %+v", i+1, got[i], c.want[i])
			}
		}
	}
}

func TestFormatHits(t *testing.T) {
	cases := []struct {
		hits int64
		want string
	}{
		{-1, ""},
		{100, "100"},
		{12300, "12.3k"},
		{12300000, "12.3m"},
	}
	for _, c := range cases {
		if got := report.FormatHits(c.hits); got != c.want {
			t.Errorf("FormatHits(%d) = %q, want %q", c.hits, got, c.want)
		}
	}
}
