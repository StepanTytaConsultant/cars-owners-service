package urlvals

import "strings"

func getKey(prefix, tag string) string {
	return strings.TrimRight(
		strings.TrimLeft(
			strings.TrimPrefix(tag, prefix),
			"["),
		"]")
}
