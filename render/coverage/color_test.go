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

func TestColor(t *testing.T) {
	cases := []struct {
		v    float64
		want string
	}{
		{
			v:    0,
			want: "#c52626",
		},
		{
			v:    60,
			want: "#ca7021",
		},
		{
			v:    80,
			want: "#339d2f",
		},
		{
			v:    math.NaN(),
			want: "#c52626",
		},
		{
			v:    math.Inf(1),
			want: "#c52626",
		},
	}

	for _, c := range cases {
		if got := coverage.Color(c.v); got != c.want {
			t.Errorf("Color(%v) = %v, want %v", c.v, got, c.want)
		}
	}
}

func Test_hslHex(t *testing.T) {
	cases := []struct {
		hue        float64
		saturation float64
		lightness  float64
		want       string
	}{
		{
			hue:        0,
			saturation: 100,
			lightness:  50,
			want:       "#ff0000",
		},
		{
			hue:        120,
			saturation: 100,
			lightness:  50,
			want:       "#00ff00",
		},
		{
			hue:        240,
			saturation: 100,
			lightness:  50,
			want:       "#0000ff",
		},
	}

	for _, c := range cases {
		if got := coverage.HSLHex(c.hue, c.saturation, c.lightness); got != c.want {
			t.Errorf("hslHex(%v, %v, %v) = %v, want %v", c.hue, c.saturation, c.lightness, got, c.want)
		}
	}
}

func Test_hueToRGB(t *testing.T) {
	cases := []struct {
		p    float64
		q    float64
		t    float64
		want float64
	}{
		{
			p:    0,
			q:    1,
			t:    2,
			want: 0,
		},
		{
			p:    0,
			q:    1,
			t:    1.0 / 3.0,
			want: 1,
		},
		{
			p:    0,
			q:    1,
			t:    -1.0 / 3.0,
			want: 0,
		},
	}

	for _, c := range cases {
		if got := coverage.HueToRGB(c.p, c.q, c.t); got != c.want {
			t.Errorf("hueToRGB(%v, %v, %v) = %v, want %v", c.p, c.q, c.t, got, c.want)
		}
	}
}
