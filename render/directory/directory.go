package directory

import (
	"fmt"
	"math"

	"github.com/itsubaki/gocov/report"
)

type Directory struct {
	Name     string
	Depth    int
	Lines    int
	Path     string
	Color    string
	Coverage string
	Share    string
}

func New(rep *report.Report) []*Directory {
	root := NewNode(rep.Directories)
	total := root.Weight()
	if total <= 0 {
		return nil
	}

	return appendDir(
		[]*Directory{},
		root,
		-90,
		270,
		root.MaxDepth()+1,
		total,
	)
}

func appendDir(dirs []*Directory, node *Node, start, end float64, ringCount, total int) []*Directory {
	// append current directory
	weight := node.Weight()
	slices := append(dirs, &Directory{
		Name:     node.displayPath,
		Depth:    node.depth,
		Lines:    weight,
		Path:     donutSegmentPath(start, end, node.depth, ringCount),
		Color:    CoverageColor(node.Stats.Percent),
		Coverage: Percent(node.Stats.Percent),
		Share:    SharePercent(weight, total),
	})

	if len(node.children) == 0 || weight < 1 {
		// no child or no coverable lines, skip children
		return slices
	}

	// append child directories
	next, span := start, end-start
	for _, v := range node.children {
		if w := v.Stats.Weight(); w > 0 {
			start, end := next, next+span*float64(w)/float64(weight)
			next = end

			// append child directory
			slices = appendDir(slices, v, start, end, ringCount, total)
		}
	}

	return slices
}

// donutSegmentPath returns the SVG path for a donut segment defined by start and end angles, depth, and total ring count.
func donutSegmentPath(start, end float64, depth, ringCount int) string {
	if end <= start {
		return ""
	}

	inner, outer := ringRadii(depth, ringCount)
	if end-start >= 359.999 {
		return fmt.Sprintf("M 0 %.4f A %.4f %.4f 0 1 1 0 %.4f A %.4f %.4f 0 1 1 0 %.4f M 0 %.4f A %.4f %.4f 0 1 0 0 %.4f A %.4f %.4f 0 1 0 0 %.4f Z",
			-outer,
			outer,
			outer,
			outer,
			outer,
			outer,
			-outer,
			-inner,
			inner,
			inner,
			inner,
			inner,
			inner,
			-inner,
		)
	}

	polarPoint := func(radius, degrees float64) (float64, float64) {
		radians := degrees * math.Pi / 180
		return radius * math.Cos(radians), radius * math.Sin(radians)
	}

	startOuterX, startOuterY := polarPoint(outer, start)
	endOuterX, endOuterY := polarPoint(outer, end)
	startInnerX, startInnerY := polarPoint(inner, start)
	endInnerX, endInnerY := polarPoint(inner, end)

	var largeArc int
	if end-start > 180 {
		largeArc = 1
	}

	return fmt.Sprintf("M %.4f %.4f A %.4f %.4f 0 %d 1 %.4f %.4f L %.4f %.4f A %.4f %.4f 0 %d 0 %.4f %.4f Z",
		startOuterX,
		startOuterY,
		outer,
		outer,
		largeArc,
		endOuterX,
		endOuterY,
		endInnerX,
		endInnerY,
		inner,
		inner,
		largeArc,
		startInnerX,
		startInnerY,
	)
}

// ringRadii returns the inner and outer radii for a given depth and total ring count.
func ringRadii(depth, count int) (float64, float64) {
	const centerRadius = 0.34
	if count <= 0 {
		return centerRadius, 1
	}

	width := (1 - centerRadius) / float64(count)
	inner := centerRadius + float64(depth)*width
	outer := centerRadius + float64(depth+1)*width
	return inner, min(outer, 1)
}

// SharePercent returns the percentage of part over total as a string with one decimal place.
func SharePercent(part, total int) string {
	if total <= 0 || part <= 0 {
		return "0.0%"
	}

	return fmt.Sprintf("%.1f%%", float64(part)*100/float64(total))
}

// Percent formats a float64 as a percentage string with one decimal place.
func Percent(v float64) string {
	if math.IsNaN(v) || math.IsInf(v, 0) {
		return "0.0%"
	}

	return fmt.Sprintf("%.1f%%", v)
}

// CoverageColor returns the color for a given coverage percentage.
func CoverageColor(v float64) string {
	if math.IsNaN(v) || math.IsInf(v, 0) {
		v = 0
	}

	switch v = min(max(v, 0), 100); {
	case v >= 80:
		return coverageBandColor(v, 80, 100, 118, 145, 54, 40)
	case v >= 60:
		return coverageBandColor(v, 60, 80, 28, 42, 72, 46)
	default:
		return coverageBandColor(v, 0, 60, 0, 8, 68, 46)
	}
}

// coverageBandColor returns the color for a given coverage percentage within a specified band.
func coverageBandColor(v, start, end, startHue, endHue, saturation, lightness float64) string {
	progress := 0.0
	if end > start {
		progress = (v - start) / (end - start)
	}

	hue := startHue + (endHue-startHue)*min(max(progress, 0), 1)
	return hslHex(hue, saturation, lightness)
}

// hslHex converts HSL color to hex string.
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

// hueToRGB converts hue to RGB component.
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
