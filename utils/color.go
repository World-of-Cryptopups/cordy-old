package utils

import "strconv"

// ParseColor parses the string hex color.
func ParseColor(color string) (int64, error) {
	return strconv.ParseInt(color, 16, 64)
}
