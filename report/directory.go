package report

type Directory struct {
	Name  string
	Stats Stats
}

func NewDirectory(files []*File) map[string]*Directory {
	dirs := make(map[string]*Directory)
	for _, f := range files {
		if dir, ok := dirs[f.Directory]; ok {
			dir.Stats.Merge(f.Stats)
			continue
		}

		dirs[f.Directory] = &Directory{
			Name:  f.Directory,
			Stats: f.Stats,
		}
	}

	return dirs
}
