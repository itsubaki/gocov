package render_test

import (
	"html/template"
	"math"
	"testing"
	"time"

	"github.com/itsubaki/gocov/render"
)

func Test_generatedAt(t *testing.T) {
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
			t.Errorf("got=%v, want=%v", got, c.want)
		}
	}
}

func Test_stylePct(t *testing.T) {
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
			t.Errorf("got=%v, want=%v", got, c.want)
		}
	}
}

func Test_covColor(t *testing.T) {
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
			t.Errorf("got=%v, want=%v", got, c.want)
		}
	}
}

func Test_lineClass(t *testing.T) {
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
			t.Errorf("got=%v, want=%v", got, c.want)
		}
	}
}
