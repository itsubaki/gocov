package render

import (
	"sort"
	"strings"

	"github.com/itsubaki/gocov/report"
)

func NewDirectoryNode(dirs []*report.Directory) *DirectoryNode {
	pathParts := func(path string) []string {
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

	root := &DirectoryNode{
		name:        "root",
		displayPath: "root",
		childByName: make(map[string]*DirectoryNode),
	}

	for _, dir := range dirs {
		if dir.Stats.Weight() == 0 {
			continue
		}

		root.Merge(dir.Stats)
		if dir.Name == "root" {
			root.MergeSelf(dir.Stats)
			continue
		}

		node, parts := root, pathParts(dir.Name)
		for i, p := range parts {
			node = node.NewChiled(p, strings.Join(parts[:i+1], "/"))
			node.Merge(dir.Stats)
		}

		node.MergeSelf(dir.Stats)
	}

	root.Refresh()
	return root
}

type DirectoryNode struct {
	name        string
	displayPath string
	depth       int
	Stats       report.Stats
	SelfStats   report.Stats
	children    []*DirectoryNode
	childByName map[string]*DirectoryNode
}

func (n *DirectoryNode) Merge(s report.Stats) {
	n.Stats.Merge(s)
}

func (n *DirectoryNode) MergeSelf(s report.Stats) {
	n.SelfStats.Merge(s)
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

func (n *DirectoryNode) NewChiled(name, displayPath string) *DirectoryNode {
	if child := n.childByName[name]; child != nil {
		return child
	}

	child := &DirectoryNode{
		name:        name,
		displayPath: displayPath,
		depth:       n.depth + 1,
		childByName: make(map[string]*DirectoryNode),
	}

	n.children = append(n.children, child)
	n.childByName[name] = child
	return child
}

func (n *DirectoryNode) Refresh() {
	n.Stats.Refresh()
	n.SelfStats.Refresh()

	for _, v := range n.children {
		v.Refresh()
	}

	nsort(n)
}

func nsort(n *DirectoryNode) {
	sort.Slice(n.children, func(i, j int) bool {
		return n.children[i].displayPath < n.children[j].displayPath
	})

	for _, v := range n.children {
		nsort(v)
	}
}
