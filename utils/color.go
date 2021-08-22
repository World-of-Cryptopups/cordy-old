package utils

import (
	"fmt"
	"strconv"
)

// ParseColor parses the string hex color.
func ParseColor(color string) int64 {
	c, err := strconv.ParseInt(color[1:], 16, 64)

	if err != nil {
		fmt.Println(err)
		return 0
	}

	return c
}
