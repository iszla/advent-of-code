package days

import (
	"strings"
)

type Day06 struct{}

func (*Day06) P01(input []string) int {
	var compat [][]string
	for i := range input {
		s := strings.Fields(input[i])
		compat = append(compat, s)
	}

	total := 0
	for x := range compat[0] {
		col := 0
		for i := range len(compat) - 1 {
			n := toInt(compat[i][x])
			switch compat[len(compat)-1][x] {
			case "*":
				if col == 0 {
					col = 1
				}
				col *= n
			case "+":
				col += n
			default:
				panic("Unknown operation " + string(compat[len(compat)-1][x]))
			}
		}
		total += col
	}

	return total
}

func (*Day06) P02(input []string) int {
	total := 0

	var values []string
	for i := len(input[0]) - 1; i >= 0; i-- {
		value := ""
		for x := 0; x < len(input)-1; x++ {
			value += string(input[x][i])
		}
		if strings.TrimRight(value, " ") == "" || i == 0 {
			op := input[len(input)-1][i+1]
			if i == 0 {
				op = input[len(input)-1][i]
				values = append(values, value)
			}
			switch op {
			case '*':
				sum := 1
				for x := range values {
					v := toInt(strings.ReplaceAll(values[x], " ", ""))
					sum *= v
				}
				total += sum
			case '+':
				sum := 0
				for x := range values {
					v := toInt(strings.ReplaceAll(values[x], " ", ""))
					sum += v
				}
				total += sum
			default:
				panic("Unknown operation " + string(op))
			}
			values = make([]string, 0)
			continue
		}
		values = append(values, value)
	}

	return total
}
