package profile_test

import (
	"testing"

	"github.com/itsubaki/gocov/profile"
)

func TestParsePoint(t *testing.T) {
	cases := []struct {
		s         string
		line, col int
		hasErr    bool
	}{
		{"1.2", 1, 2, false},
		{"10.20", 10, 20, false},
		{"a.b.c", 0, 0, true},
		{"a.1", 0, 0, true},
		{"1.b", 0, 0, true},
	}

	for _, c := range cases {
		line, col, err := profile.ParsePoint(c.s)
		if err != nil {
			if !c.hasErr {
				t.Errorf("unexpected error: %v", err)
			}

			continue
		}

		if line != c.line || col != c.col {
			t.Errorf("got=(%d, %d), want=(%d, %d)", line, col, c.line, c.col)
		}
	}
}
