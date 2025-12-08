package days

type Day03 struct{}

func (*Day03) P01(input []string) int {
	total := 0
	for i := range input {
		first, j := findJolt(input[i], 0)
		second, _ := findJolt(input[i], j+1)
		jolt := toInt(first + second)
		total += jolt
	}
	return total
}

func (*Day03) P02(input []string) int {
	total := 0
	for i := range input {
		joltage := ""
		j := 0
		skipped := 0
		for len(joltage) < 12 {
			jolt, k := findNextJolt(input[i], j, skipped)
			joltage = joltage + jolt
			skipped += k - j
			j = k + 1
		}
		b := toInt(joltage)
		total += b
	}
	return total
}

func findJolt(bank string, start int) (string, int) {
	r := []rune(bank)
	end := len(r) - 1
	if start > 0 {
		end = len(r)
	}
	maxv := 0
	maxi := 0
	for i := start; i < end; i++ {
		d := toInt(string(r[i]))
		if d > maxv {
			maxv = d
			maxi = i
		}
	}

	return toString(maxv), maxi
}

func findNextJolt(bank string, start int, skipped int) (string, int) {
	r := []rune(bank)
	end := len(r) - 12
	maxv := 0
	maxi := 0
	for i := start; i <= (start + end - skipped); i++ {
		d := toInt(string(r[i]))
		if d > maxv {
			maxv = d
			maxi = i
		}
	}

	return toString(maxv), maxi
}
