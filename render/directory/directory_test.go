package directory_test

import (
	"testing"

	"github.com/itsubaki/gocov/render/directory"
)

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
