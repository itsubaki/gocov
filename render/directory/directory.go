package directory

import (
	"fmt"
	"math"

	"github.com/itsubaki/gocov/render/coverage"
	"github.com/itsubaki/gocov/report"
)

type Directory struct {
	Name       string
	Depth      int
	Statements int
	Path       string
	Color      string
	Coverage   string
	Share      string
}

func New(rep *report.Report) []*Directory {
	root := NewNode(rep.Directories)
	total := root.TotalStatements()
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
	weight := node.TotalStatements()
	slices := append(dirs, &Directory{
		Name:       node.displayPath,
		Depth:      node.depth,
		Statements: weight,
		Path:       donutSegmentPath(start, end, node.depth, ringCount),
		Color:      coverage.Color(node.Stats.Percent),
		Coverage:   coverage.Percent(node.Stats.Percent),
		Share:      coverage.SharePercent(weight, total),
	})

	if len(node.children) == 0 || weight < 1 {
		// no child or no coverable lines, skip children
		return slices
	}

	// append child directories
	next, span := start, end-start
	for _, v := range node.children {
		if w := v.TotalStatements(); w > 0 {
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
