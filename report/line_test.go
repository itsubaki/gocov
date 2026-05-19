package report_test

import (
	"testing"

	"github.com/itsubaki/gocov/report"
)

func Test_formatHits(t *testing.T) {
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
			t.Errorf("formatHits(%d) = %q, want %q", c.hits, got, c.want)
		}
	}
}
