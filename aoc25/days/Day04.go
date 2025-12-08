package days

type Day04 struct{}

var directions = [8][2]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func (*Day04) P01(input []string) int {
	total := 0
	for x := 0; x < len(input); x++ {
		for y := 0; y < len(input[0]); y++ {
			if input[x][y] == '@' && countNeighbours(x, y, input) < 4 {
				total++
			}
		}
	}
	return total
}

func (*Day04) P02(input []string) int {
	total := 0
	removed := true
	for removed {
		removed = false
		for x := 0; x < len(input); x++ {
			for y := 0; y < len(input[0]); y++ {
				if input[x][y] != '@' {
					continue
				}
				if countNeighbours(x, y, input) < 4 {
					input[x] = input[x][:y] + "." + input[x][y+1:]
					total++
					removed = true
				}
			}
		}
	}
	return total
}

func countNeighbours(x int, y int, shelf []string) int {
	maxx := len(shelf)
	maxy := len(shelf[0])

	total := 0
	for _, d := range directions {
		nx := x + d[0]
		ny := y + d[1]
		if nx >= 0 && nx < maxx && ny >= 0 && ny < maxy {
			if shelf[nx][ny] == '@' {
				total++
				if total >= 4 {
					return total
				}
			}
		}
	}

	return total
}
