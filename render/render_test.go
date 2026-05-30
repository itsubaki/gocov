package render_test

import (
	"bytes"
	"fmt"
	"html/template"
	"math"
	"testing"
	"time"

	"github.com/itsubaki/gocov/render"
	"github.com/itsubaki/gocov/report"
)

func ExampleHTML() {
	rep := &report.Report{
		GeneratedAt:  time.Now(),
		RootPath:     "/path/to/root",
		ProfilePath:  "/path/to/profile",
		OutputPath:   "/path/to/output",
		Mode:         "atomic",
		ModulePath:   "github.com/itsubaki/gocov",
		Stats:        report.Stats{},
		Files:        []*report.File{},
		Directories:  []*report.Directory{},
		MissingFiles: []string{},
	}

	var buf bytes.Buffer
	if err := render.HTML(&buf, rep); err != nil {
		panic(err)
	}

	fmt.Println(buf.String()[:16])

	// Output:
	// <!doctype html>
}

func TestGeneratedAt(t *testing.T) {
	cases := []struct {
		t    time.Time
		want string
	}{
		{
			time.Time{},
			"",
		},
		{
			time.Date(2024, 6, 1, 12, 0, 0, 0, time.UTC),
			"2024-06-01 12:00:00 UTC",
		},
	}

	for _, c := range cases {
		if got := render.GeneratedAt(c.t); got != c.want {
			t.Errorf("GeneratedAt(%v) = %v, want %v", c.t, got, c.want)
		}
	}
}

func TestStylePct(t *testing.T) {
	cases := []struct {
		v    float64
		want string
	}{
		{0, "0.0000%"},
		{50, "50.0000%"},
		{100, "100.0000%"},
		{math.NaN(), "0%"},
		{math.Inf(1), "0%"},
	}

	for _, c := range cases {
		if got := render.StylePct(c.v); got != c.want {
			t.Errorf("StylePct(%v) = %v, want %v", c.v, got, c.want)
		}
	}
}

func TestCovColor(t *testing.T) {
	cases := []struct {
		v    float64
		want template.CSS
	}{
		{0, "#c52626"},
		{50, "#c53726"},
		{59.9, "#c53b26"},
		{60, "#ca7021"},
		{79.9, "#ca9721"},
		{80, "#339d2f"},
		{100, "#2f9d5d"},
		{math.NaN(), "#c52626"},
		{math.Inf(1), "#c52626"},
	}

	for _, c := range cases {
		if got := render.CovColor(c.v); got != c.want {
			t.Errorf("CovColor(%v) = %v, want %v", c.v, got, c.want)
		}
	}
}

func TestLineClass(t *testing.T) {
	cases := []struct {
		state string
		want  string
	}{
		{"covered", "line-covered"},
		{"partial", "line-partial"},
		{"missed", "line-missed"},
	}

	for _, c := range cases {
		if got := render.LineClass(c.state); got != c.want {
			t.Errorf("LineClass(%q) = %q, want %q", c.state, got, c.want)
		}
	}
}
