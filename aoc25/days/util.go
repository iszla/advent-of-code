package days

import (
	"strconv"
)

const INT_MAX_VALUE = int(^uint(0) >> 1)

type Day interface {
	P01(input []string) int
	P02(input []string) int
}

func toInt(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func toFloat(s string) float64 {
	v, _ := strconv.ParseFloat(s, 64)
	return v
}

func toString(i int) string {
	return strconv.Itoa(i)
}

func removeAtIndex[K comparable](s []K, i int) []K {
	return append(s[:i], s[i+1:]...)
}

func isEqual[K comparable](a, b []K) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
