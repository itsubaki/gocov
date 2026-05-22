package directory_test

import (
	"math"
	"testing"

	"github.com/itsubaki/gocov/render/directory"
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
		if got := directory.Percent(c.v); got != c.want {
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
		if got := directory.SharePercent(c.part, c.total); got != c.want {
			t.Errorf("SharePercent(%v, %v) = %v, want %v", c.part, c.total, got, c.want)
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
