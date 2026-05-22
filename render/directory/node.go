package directory

import (
	"sort"
	"strings"

	"github.com/itsubaki/gocov/report"
)

type Node struct {
	Stats       report.Stats
	SelfStats   report.Stats
	name        string
	displayPath string
	depth       int
	children    []*Node
	childByName map[string]*Node
}

func NewNode(dirs []*report.Directory) *Node {
	root := &Node{
		name:        "root",
		displayPath: "root",
		childByName: make(map[string]*Node),
	}

	for _, dir := range dirs {
		if dir.Stats.Weight() == 0 {
			continue
		}

		if dir.Name == "root" {
			root.SelfStats = report.Merge(root.SelfStats, dir.Stats)
			continue
		}

		next, parts := root, pathParts(dir.Name)
		for i, name := range parts {
			if child := next.childByName[name]; child != nil {
				next = child
				continue
			}
			next = next.Add(&Node{
				name:        name,
				displayPath: strings.Join(parts[:i+1], "/"),
				depth:       next.depth + 1,
				childByName: make(map[string]*Node),
			})
		}
		next.SelfStats = report.Merge(next.SelfStats, dir.Stats)
	}

	accumulate(root)
	nsort(root)
	return root
}

// Weight returns the total number of coverable lines in the directory and its subdirectories.
func (n *Node) Weight() int {
	var sum int
	for _, c := range n.children {
		sum += c.Stats.Weight()
	}

	return sum + n.SelfStats.Weight()
}

func (n *Node) MaxDepth() int {
	depth := n.depth
	if len(n.children) > 0 && n.SelfStats.Weight() > 0 {
		depth = max(depth, n.depth+1)
	}

	for _, v := range n.children {
		depth = max(depth, v.MaxDepth())
	}

	return depth
}

func (n *Node) Add(child *Node) *Node {
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

func accumulate(n *Node) report.Stats {
	sum := n.SelfStats
	for _, c := range n.children {
		sum = report.Merge(sum, accumulate(c))
	}

	n.Stats = sum
	return sum
}

func nsort(n *Node) {
	sort.Slice(n.children, func(i, j int) bool {
		return n.children[i].displayPath < n.children[j].displayPath
	})

	for _, v := range n.children {
		nsort(v)
	}
}
