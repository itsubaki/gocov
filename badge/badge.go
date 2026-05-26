package badge

import "fmt"

func SVG(coverage float64) string {
	length := 260
	if int(coverage) == 100 {
		length = 320
	}

	return fmt.Sprintf(`
<svg xmlns="http://www.w3.org/2000/svg" width="100" height="20" role="img">
  <g shape-rendering="crispEdges">
    <rect width="61" height="20" fill="#555"/>
    <rect x="61" width="39" height="20" fill="#4b0"/>
  </g>
  <g fill="#fff" text-anchor="middle" font-family="Verdana,Geneva,DejaVu Sans,sans-serif" text-rendering="geometricPrecision" font-size="110">
    <text x="305" y="140" transform="scale(.1)" fill="#fff" textLength="510">coverage</text>
    <text x="805" y="140" transform="scale(.1)" fill="#fff" textLength="%d">%d%%</text>
  </g>
</svg>
`,
		length,
		int(coverage),
	)
}
