package coverage

import (
	"fmt"
	"math"
)

// Percent formats a float64 as a percentage string with one decimal place.
func Percent(v float64) string {
	if math.IsNaN(v) || math.IsInf(v, 0) {
		return "0.0%"
	}

	return fmt.Sprintf("%.1f%%", v)
}

// SharePercent returns the percentage of part over total as a string with one decimal place.
func SharePercent(part, total int) string {
	if total <= 0 || part <= 0 {
		return "0.0%"
	}

	return fmt.Sprintf("%.1f%%", float64(part)*100/float64(total))
}

// Color returns the color for a given coverage percentage.
func Color(v float64) string {
	if math.IsNaN(v) || math.IsInf(v, 0) {
		v = 0
	}

	switch v = min(max(v, 0), 100); {
	case v >= 80:
		return bandColor(v, 80, 100, 118, 145, 54, 40)
	case v >= 60:
		return bandColor(v, 60, 80, 28, 42, 72, 46)
	default:
		return bandColor(v, 0, 60, 0, 8, 68, 46)
	}
}

func bandColor(v, start, end, startHue, endHue, saturation, lightness float64) string {
	progress := 0.0
	if end > start {
		progress = (v - start) / (end - start)
	}

	hue := startHue + (endHue-startHue)*min(max(progress, 0), 1)
	return hslHex(hue, saturation, lightness)
}

func hslHex(hue, saturation, lightness float64) string {
	h := math.Mod(hue, 360) / 360
	s := min(max(saturation/100, 0), 1)
	l := min(max(lightness/100, 0), 1)

	r, g, b := l, l, l
	if s != 0 {
		q := l * (1 + s)
		if l >= 0.5 {
			q = l + s - l*s
		}

		p := 2*l - q
		r = hueToRGB(p, q, h+1.0/3.0)
		g = hueToRGB(p, q, h)
		b = hueToRGB(p, q, h-1.0/3.0)
	}

	return fmt.Sprintf("#%02x%02x%02x", int(math.Round(r*255)), int(math.Round(g*255)), int(math.Round(b*255)))
}

func hueToRGB(p, q, t float64) float64 {
	if t < 0 {
		t++
	}

	if t > 1 {
		t--
	}

	switch {
	case t < 1.0/6.0:
		return p + (q-p)*6*t
	case t < 1.0/2.0:
		return q
	case t < 2.0/3.0:
		return p + (q-p)*(2.0/3.0-t)*6
	default:
		return p
	}
}
