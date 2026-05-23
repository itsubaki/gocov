package render

import (
	"fmt"
	"html/template"
	"io"
	"math"
	"time"

	"github.com/itsubaki/gocov/render/coverage"
	"github.com/itsubaki/gocov/render/directory"
	"github.com/itsubaki/gocov/report"
)

// HTML renders the coverage report as HTML.
func HTML(w io.Writer, rep *report.Report) error {
	tmpl, err := template.New("report").Funcs(template.FuncMap{
		"directories":   directory.New,
		"sharePct":      coverage.SharePercent,
		"pct":           coverage.Percent,
		"coverageColor": coverage.Color,
		"generatedAt":   generatedAt,
		"stylePct":      stylePct,
		"lineClass":     lineClass,
	}).Parse(reportTemplate)
	if err != nil {
		return err
	}

	return tmpl.Execute(w, rep)
}

func generatedAt(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	return t.Format("2006-01-02 15:04:05 MST")
}

func stylePct(v float64) string {
	if math.IsNaN(v) || math.IsInf(v, 0) {
		return "0%"
	}

	return fmt.Sprintf("%.4f%%", min(max(v, 0), 100))
}

func covColor(v float64) template.CSS {
	return template.CSS(coverage.Color(v))
}

func lineClass(state string) string {
	return "line-" + state
}
