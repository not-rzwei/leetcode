package helper

import (
	"fmt"
	"strings"
)

func SliceToString[T any](arr []T) string {
	var sb strings.Builder
	sb.WriteString("[")

	for i, v := range arr {
		if isNil(v) {
			sb.WriteString("nil")
		} else {
			sb.WriteString(fmt.Sprintf("%v", v))
		}

		if i < len(arr)-1 {
			sb.WriteString(",")
		}
	}

	sb.WriteString("]")
	return sb.String()
}

func isNil(v any) bool {
	return v == nil
}
