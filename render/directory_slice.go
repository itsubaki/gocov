package render

import "github.com/itsubaki/gocov/report"

func NewDirectorySlice(rep *report.Report) []*Directory {
	root := NewDirectoryNode(rep.Directories)
	total := root.Stats.Weight()
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

func appendDir(
	dirs []*Directory,
	node *DirectoryNode,
	start float64,
	end float64,
	ringCount int,
	total int,
) []*Directory {
	slices := append(dirs, NewDirectory(
		node.Stats,
		node.displayPath,
		node.depth,
		start, end,
		ringCount, total,
	))

	if len(node.children) == 0 {
		return slices
	}

	weight := node.Stats.Weight()
	if weight <= 0 {
		return slices
	}

	next, span := start, end-start
	if w := node.SelfStats.Weight(); w > 0 {
		path := "root files"
		if node.displayPath != "root" {
			path = node.displayPath + " files"
		}

		end := next + span*float64(w)/float64(weight)
		slices = append(slices, NewDirectory(
			node.SelfStats,
			path,
			node.depth+1,
			next,
			end,
			ringCount,
			total,
		))

		next = end
	}

	for _, v := range node.children {
		w := v.Stats.Weight()
		if w <= 0 {
			continue
		}

		end := next + span*float64(w)/float64(weight)
		slices = appendDir(
			slices,
			v,
			next,
			end,
			ringCount,
			total,
		)

		next = end
	}

	return slices
}
