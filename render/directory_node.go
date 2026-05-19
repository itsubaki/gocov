package render

import (
	"sort"
	"strings"

	"github.com/itsubaki/gocov/report"
)

type DirectoryNode struct {
	name        string
	displayPath string
	depth       int
	summary     report.Summary
	selfSummary report.Summary
	children    []*DirectoryNode
	childByName map[string]*DirectoryNode
}

func NewDirectoryNode(dirs []report.Directory) *DirectoryNode {
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
		if dir.Summary.Weight() == 0 {
			continue
		}

		root.AddSummary(dir.Summary)
		if dir.Name == "root" {
			root.AddSelfSummary(dir.Summary)
			continue
		}

		node, parts := root, pathParts(dir.Name)
		for idx, part := range parts {
			node = node.AddChild(part, strings.Join(parts[:idx+1], "/"))
			node.AddSummary(dir.Summary)
		}

		node.AddSelfSummary(dir.Summary)
	}

	return root.Reflesh().Sort()
}

func (n *DirectoryNode) AddSummary(s report.Summary) {
	n.summary.Add(s)
}

func (n *DirectoryNode) AddSelfSummary(s report.Summary) {
	n.selfSummary.Add(s)
}

func (n *DirectoryNode) MaxDepth() int {
	depth := n.depth
	if len(n.children) > 0 && n.selfSummary.Weight() > 0 {
		depth = max(depth, n.depth+1)
	}

	for _, child := range n.children {
		depth = max(depth, child.MaxDepth())
	}

	return depth
}

func (n *DirectoryNode) AddChild(name, displayPath string) *DirectoryNode {
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

func (n *DirectoryNode) Reflesh() *DirectoryNode {
	n.summary.Refresh()
	n.selfSummary.Refresh()

	for _, child := range n.children {
		child.Reflesh()
	}

	return n
}

func (n *DirectoryNode) Sort() *DirectoryNode {
	sort.Slice(n.children, func(i, j int) bool {
		return n.children[i].displayPath < n.children[j].displayPath
	})

	for _, child := range n.children {
		child.Sort()
	}

	return n
}
