package utils

import (
	"strconv"
	"strings"
)

func StringToArrayInt(in string, sep string) []int {
	out := make([]int, 0)
	for _, v := range strings.Split(in, sep) {
		n, _ := strconv.Atoi(v)
		out = append(out, n)
	}
	return out
}