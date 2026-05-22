package render

import (
	"sort"
	"strings"

	"github.com/itsubaki/gocov/report"
)

type DirectoryNode struct {
	Stats       *report.Stats
	SelfStats   *report.Stats
	name        string
	displayPath string
	depth       int
	children    []*DirectoryNode
	childByName map[string]*DirectoryNode
}

func NewDirectoryNode(dirs []*report.Directory) *DirectoryNode {
	root := &DirectoryNode{
		Stats:       &report.Stats{},
		SelfStats:   &report.Stats{},
		name:        "root",
		displayPath: "root",
		childByName: make(map[string]*DirectoryNode),
	}

	for _, dir := range dirs {
		if dir.Stats.Weight() == 0 {
			// no coverable lines, skip directory
			continue
		}

		// root
		root.Stats = report.Merge(root.Stats, dir.Stats)
		if dir.Name == "root" {
			root.SelfStats = report.Merge(root.SelfStats, dir.Stats)
			continue
		}

		// child
		next, parts := root, pathParts(dir.Name)
		for i, name := range parts {
			if child := next.childByName[name]; child != nil {
				next = child
				continue
			}

			// create child
			next = next.Add(&DirectoryNode{
				Stats:       dir.Stats,
				SelfStats:   &report.Stats{},
				name:        name,
				displayPath: strings.Join(parts[:i+1], "/"),
				depth:       next.depth + 1,
				childByName: make(map[string]*DirectoryNode),
			})
		}

		// self directory
		next.SelfStats = report.Merge(next.SelfStats, dir.Stats)
	}

	nsort(root)
	return root
}

func (n *DirectoryNode) Name() string {
	if n.displayPath == "root" {
		return "root files"
	}

	return n.displayPath + " files"
}

func (n *DirectoryNode) MaxDepth() int {
	depth := n.depth
	if len(n.children) > 0 && n.SelfStats.Weight() > 0 {
		depth = max(depth, n.depth+1)
	}

	for _, v := range n.children {
		depth = max(depth, v.MaxDepth())
	}

	return depth
}

func (n *DirectoryNode) Add(child *DirectoryNode) *DirectoryNode {
	n.childByName[child.name] = child
	n.children = append(n.children, child)
	return child
}

func pathParts(path string) []string {
	slash := strings.ReplaceAll(path, "\\", "/")
	parts := strings.Split(slash, "/")

	out := make([]string, 0, len(parts))
	for _, part := range parts {
		if part == "" || part == "." {
			continue
		}

		out = append(out, part)
	}

	return out
}

func nsort(n *DirectoryNode) {
	sort.Slice(n.children, func(i, j int) bool {
		return n.children[i].displayPath < n.children[j].displayPath
	})

	for _, v := range n.children {
		nsort(v)
	}
}
