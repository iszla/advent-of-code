package days

import (
	"slices"
	"sort"
	"strings"
)

type Day05 struct{}

func (*Day05) P01(input []string) int {
	r, i := process(input)

	total := 0
	for x := range i {
		if isFresh(i[x], r) {
			total++
		}
	}
	return total
}

func (*Day05) P02(input []string) int {
	r, _ := process(input)

	sort.Slice(r, func(i, j int) bool {
		return r[i].min > r[j].min
	})

	for i := len(r) - 1; i > 0; i-- {
		if r[i].max >= r[i-1].min {
			if r[i-1].max > r[i].max {
				r[i].max = r[i-1].max
			}
			r = removeAtIndex(r, i-1)
		}
	}

	total := 0
	for i := 0; i < len(r); i++ {
		total += r[i].max - r[i].min + 1
	}

	return total
}

func process(input []string) (ranges []Range, ingredients []int) {
	d := slices.Index(input, "")
	ids := input[:d]
	for i := range ids {
		s := strings.Split(ids[i], "-")
		ranges = append(ranges, Range{toInt(s[0]), toInt(s[1])})
	}

	items := input[d:]
	for i := range items {
		ingredients = append(ingredients, toInt(items[i]))
	}

	return ranges, ingredients
}

func isFresh(item int, ranges []Range) bool {
	for i := range ranges {
		if ranges[i].min <= item && ranges[i].max >= item {
			return true
		}
	}
	return false
}

type Range struct {
	min int
	max int
}
