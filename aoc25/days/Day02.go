package days

import (
	"strings"
)

type Day02 struct{}

func (*Day02) P01(input []string) int {
	total := 0

	ids := strings.Split(input[0], ",")
	for i := range ids {
		s := strings.Split(ids[i], "-")
		s1 := toInt(s[0])
		s2 := toInt(s[1])
		for i := s1; i < s2+1; i++ {
			if !isValidID(i) {
				total += i
			}
		}
	}

	return total
}

func (*Day02) P02(input []string) int {
	total := 0

	ids := strings.Split(input[0], ",")
	for i := range ids {
		s := strings.Split(ids[i], "-")
		s1 := toInt(s[0])
		s2 := toInt(s[1])
		for i := s1; i < s2+1; i++ {
			if !isValidIDv2(i) {
				total += i
			}
		}
	}

	return total
}

func isValidID(id int) bool {
	s := toString(id)
	n := len(s)
	if n%2 != 0 {
		return true
	}
	h := n / 2
	if s[:h] == s[h:] {
		return false
	}
	return true
}

func isValidIDv2(id int) bool {
	s := toString(id)
	n := len(s)
	for i := 1; i <= n/2; i++ {
		if n%i != 0 {
			continue
		}
		t := s[:i]
		if t == s[i:i+i] {
			notValid := true
			for j := i + i; j < n; j += i {
				if t != s[j:j+i] {
					notValid = false
					break
				}
			}
			if notValid {
				return false
			}
		}
	}

	return true
}
