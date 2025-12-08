package days

type Day01 struct{}

func (*Day01) P01(input []string) int {
	dial := 50
	total := 0

	for i := range input {
		row := []rune(input[i])
		dir := row[0]
		s := toInt(string(row[1:]))
		switch dir {
		case 'R':
			dial, _ = turnDial(dial, s)
		case 'L':
			dial, _ = turnDial(dial, -s)
		default:
			panic("Unknown direction " + string(dir))
		}
		if dial == 0 {
			total++
		}
	}

	return total
}

func (*Day01) P02(input []string) int {
	dial := 50
	total := 0

	for i := range input {
		row := []rune(input[i])
		dir := row[0]
		s := toInt(string(row[1:]))
		var t int
		switch dir {
		case 'R':
			dial, t = turnDial(dial, s)
		case 'L':
			dial, t = turnDial(dial, -s)
		default:
			panic("Unknown direction " + string(dir))
		}
		total += t
	}

	return total
}

func turnDial(dial int, steps int) (int, int) {
	total := 0
	if steps < 0 {
		for i := steps; i < 0; i++ {
			dial--
			if dial == 0 {
				total++
			}
			if dial == -1 {
				dial = 99
			}
		}
	} else {
		for i := steps; i > 0; i-- {
			dial++
			if dial == 100 {
				dial = 0
			}
			if dial == 0 {
				total++
			}
		}
	}
	return dial, total
}
