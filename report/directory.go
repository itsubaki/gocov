package report

import "sort"

type Directory struct {
	Name  string
	Stats *Stats
}

func NewDirectory(files []*File) []*Directory {
	dirs := make(map[string]*Directory)
	for _, f := range files {
		if dir, ok := dirs[f.Directory]; ok {
			dir.Stats.Merge(f.Stats)
			dir.Stats.Update()
			continue
		}

		dirs[f.Directory] = &Directory{
			Name:  f.Directory,
			Stats: f.Stats,
		}
	}

	out := make([]*Directory, 0, len(dirs))
	for _, dir := range dirs {
		out = append(out, dir)
	}

	sort.Slice(out, func(i, j int) bool {
		return out[i].Name < out[j].Name
	})

	return out
}
