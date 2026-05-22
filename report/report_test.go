package report_test

import (
	"strings"
	"testing"

	"github.com/itsubaki/gocov/report"
)

func Test_modulePath(t *testing.T) {
	cases := []struct {
		gomod string
		want  string
	}{
		{
			gomod: `module example.com/myapp
go 1.22
`,
			want: "example.com/myapp",
		},
		{
			gomod: `go 1.22`,
			want:  "",
		},
		{
			gomod: `

module   example.com/myapp   

go 1.22
`,
			want: "example.com/myapp",
		},
		{
			gomod: "",
			want:  "",
		},
	}

	for _, c := range cases {
		got := report.ModulePath(strings.NewReader(c.gomod))
		if got != c.want {
			t.Fatalf("got %q, want %q", got, c.want)
		}
	}
}
