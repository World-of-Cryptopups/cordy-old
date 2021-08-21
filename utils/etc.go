package utils

import "strconv"

// convert to int directly, ignore errors
func ConvertInt(s string) int {
	x, _ := strconv.Atoi(s)
	return x
}
