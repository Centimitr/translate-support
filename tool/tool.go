package tool

import (
	"strings"
)

func StringSliceToJson(arr []string) string {
	return "[\"" + strings.Join(arr, "\",\"") + "\"]"
}

func StringSliceRemove(slice []string, index int) []string {
	var r = slice[:0]
	var rlen = 0
	if len(slice) <= index {
		// log error
	} else {
		for i, item := range slice {
			if i != index {
				r = append(r, item)
				rlen++
			}
		}
	}
	return r[:rlen]
}
