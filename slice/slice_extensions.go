package slice

import (
	"strconv"
	"strings"
)

func Contains(s []int32, e int32) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func StringContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func Unique(s []int32) []int32 {
	keys := make(map[int32]bool)
	list := []int32{}
	for _, entry := range s {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func Join(s []int32, separator string) string {
	b := make([]string, len(s))
	for i, v := range s {
		b[i] = strconv.FormatInt(int64(v), 10)
	}
	return strings.Join(b, separator)
}
