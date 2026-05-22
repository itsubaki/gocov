package directory_test

import (
	"reflect"
	"testing"

	"github.com/itsubaki/gocov/render/directory"
)

func Test_pathParts(t *testing.T) {
	cases := []struct {
		path string
		want []string
	}{
		{
			path: "root",
			want: []string{"root"},
		},
		{
			path: "foo/bar/baz",
			want: []string{"foo", "bar", "baz"},
		},
	}

	for _, c := range cases {
		if got := directory.PathParts(c.path); !reflect.DeepEqual(got, c.want) {
			t.Errorf("pathParts(%v) = %v, want %v", c.path, got, c.want)
		}
	}
}
