package helper

import (
	"fmt"
	"strings"
)

func SliceToString(arr []interface{}) string {
	var sb strings.Builder
	sb.WriteString("[")

	for i, v := range arr {
		if v == nil {
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
