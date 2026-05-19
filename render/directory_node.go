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

		root.addSummary(dir.Summary)
		if dir.Name == "root" {
			root.addSelfSummary(dir.Summary)
			continue
		}

		node, parts := root, pathParts(dir.Name)
		for idx, part := range parts {
			node = node.addChild(part, strings.Join(parts[:idx+1], "/"))
			node.addSummary(dir.Summary)
		}

		node.addSelfSummary(dir.Summary)
	}

	return root.reflesh().sort()
}

func (n *DirectoryNode) addSummary(s report.Summary) {
	n.summary.Add(s)
}

func (n *DirectoryNode) addSelfSummary(s report.Summary) {
	n.selfSummary.Add(s)
}

func (n *DirectoryNode) maxDepth() int {
	depth := n.depth
	if len(n.children) > 0 && n.selfSummary.Weight() > 0 {
		depth = max(depth, n.depth+1)
	}

	for _, child := range n.children {
		depth = max(depth, child.maxDepth())
	}

	return depth
}

func (n *DirectoryNode) addChild(name, displayPath string) *DirectoryNode {
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

func (n *DirectoryNode) reflesh() *DirectoryNode {
	n.summary.Refresh()
	n.selfSummary.Refresh()

	for _, child := range n.children {
		child.reflesh()
	}

	return n
}

func (n *DirectoryNode) sort() *DirectoryNode {
	sort.Slice(n.children, func(i, j int) bool {
		return n.children[i].displayPath < n.children[j].displayPath
	})

	for _, child := range n.children {
		child.sort()
	}

	return n
}
