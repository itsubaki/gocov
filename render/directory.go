package render

import (
	"fmt"
	"math"

	"github.com/itsubaki/gocov/report"
)

func NewDirectorySlice(rep report.Report) []Directory {
	root := NewDirectoryNode(rep.Directories)
	total := root.summary.Weight()
	if total <= 0 {
		return nil
	}

	return appendDir(
		[]Directory{},
		root,
		-90,
		270,
		root.maxDepth()+1,
		total,
	)
}

func appendDir(dirs []Directory, node *DirectoryNode, start, end float64, ringCount, total int) []Directory {
	slices := append(dirs, NewDirectory(
		node.summary,
		node.displayPath,
		node.depth,
		start, end,
		ringCount, total,
	))

	if len(node.children) == 0 {
		return slices
	}

	weight := node.summary.Weight()
	if weight <= 0 {
		return slices
	}

	next, span := start, end-start
	if selfWeight := node.selfSummary.Weight(); selfWeight > 0 {
		path := "root files"
		if node.displayPath != "root" {
			path = node.displayPath + " files"
		}

		end := next + span*float64(selfWeight)/float64(weight)
		slices = append(slices, NewDirectory(
			node.selfSummary,
			path,
			node.depth+1,
			next,
			end,
			ringCount,
			total,
		))

		next = end
	}

	for _, child := range node.children {
		childWeight := child.summary.Weight()
		if childWeight <= 0 {
			continue
		}

		end := next + span*float64(childWeight)/float64(weight)
		slices = appendDir(
			slices,
			child,
			next,
			end,
			ringCount,
			total,
		)

		next = end
	}

	return slices
}

type Directory struct {
	Name     string
	Path     string
	Color    string
	Coverage string
	Share    string
	Lines    int
	Depth    int
}

func NewDirectory(summary report.Summary, name string, depth int, start, end float64, ringCount, total int) Directory {
	return Directory{
		Name:     name,
		Path:     donutSegmentPath(start, end, depth, ringCount),
		Color:    coverageColor(summary.Percent),
		Coverage: percent(summary.Percent),
		Share:    sharePercent(summary.Weight(), total),
		Lines:    summary.Weight(),
		Depth:    depth,
	}
}

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

func ringRadii(depth, ringCount int) (float64, float64) {
	const centerRadius = 0.34
	if ringCount <= 0 {
		return centerRadius, 1
	}

	ringWidth := (1 - centerRadius) / float64(ringCount)
	inner := centerRadius + float64(depth)*ringWidth
	outer := centerRadius + float64(depth+1)*ringWidth
	return inner, min(outer, 1)
}

func sharePercent(part, total int) string {
	if total <= 0 || part <= 0 {
		return "0.0%"
	}

	return fmt.Sprintf("%.1f%%", float64(part)*100/float64(total))
}

func percent(v float64) string {
	if math.IsNaN(v) || math.IsInf(v, 0) {
		return "0.0%"
	}

	return fmt.Sprintf("%.1f%%", v)
}

func coverageColor(v float64) string {
	if math.IsNaN(v) || math.IsInf(v, 0) {
		v = 0
	}

	v = min(max(v, 0), 100)

	switch {
	case v >= 80:
		return coverageBandColor(v, 80, 100, 118, 145, 54, 40)
	case v >= 60:
		return coverageBandColor(v, 60, 80, 28, 42, 72, 46)
	default:
		return coverageBandColor(v, 0, 60, 0, 8, 68, 46)
	}
}

func coverageBandColor(v, start, end, startHue, endHue, saturation, lightness float64) string {
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
