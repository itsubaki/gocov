package coverage_test

import (
	"math"
	"testing"

	"github.com/itsubaki/gocov/render/coverage"
)

func TestPercent(t *testing.T) {
	cases := []struct {
		v    float64
		want string
	}{
		{
			v:    0,
			want: "0.0%",
		},
		{
			v:    50,
			want: "50.0%",
		},
		{
			v:    math.NaN(),
			want: "0.0%",
		},
		{
			v:    math.Inf(1),
			want: "0.0%",
		},
	}

	for _, c := range cases {
		if got := coverage.Percent(c.v); got != c.want {
			t.Errorf("Percent(%v) = %v, want %v", c.v, got, c.want)
		}
	}
}

func TestSharePercent(t *testing.T) {
	cases := []struct {
		part  int
		total int
		want  string
	}{
		{
			part:  0,
			total: 100,
			want:  "0.0%",
		},
		{
			part:  100,
			total: 0,
			want:  "0.0%",
		},
		{
			part:  50,
			total: 100,
			want:  "50.0%",
		},
	}

	for _, c := range cases {
		if got := coverage.SharePercent(c.part, c.total); got != c.want {
			t.Errorf("SharePercent(%v, %v) = %v, want %v", c.part, c.total, got, c.want)
		}
	}
}
