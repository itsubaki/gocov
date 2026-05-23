package coverage

import (
	"fmt"
	"math"
)

// Percent formats a float64 as a percentage string with one decimal place.
func Percent(v float64) string {
	if math.IsNaN(v) || math.IsInf(v, 0) {
		return "0.0%"
	}

	return fmt.Sprintf("%.1f%%", v)
}

// SharePercent returns the percentage of part over total as a string with one decimal place.
func SharePercent(part, total int) string {
	if total <= 0 || part <= 0 {
		return "0.0%"
	}

	return fmt.Sprintf("%.1f%%", float64(part)*100/float64(total))
}
