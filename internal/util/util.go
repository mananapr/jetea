package util

func StringPtr(s string) *string {
	return &s
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
