package days

type Day07 struct{}

func (*Day07) P01(input []string) int {
	var tree [][]rune
	for i := range input {
		tree = append(tree, []rune(input[i]))
	}
	total := 0
	for i := 0; i < len(tree); i++ {
		for j := range tree[i] {
			switch tree[i][j] {
			case '^':
				if tree[i-1][j] == '|' {
					total++
					tree[i][j-1] = '|'
					tree[i][j+1] = '|'
				}
			case 'S':
				tree[i+1][j] = '|'
			default:
				if i > 0 && tree[i-1][j] == '|' {
					tree[i][j] = '|'
				}
			}
		}
	}

	return total
}

func (*Day07) P02(input []string) int {
	var tree [][]rune
	for i := range input {
		tree = append(tree, []rune(input[i]))
	}
	visited := make([]int, len(tree[0]))
	for i := 0; i < len(tree); i++ {
		for j := range tree[i] {
			switch tree[i][j] {
			case '^':
				if tree[i-1][j] == '|' {
					tree[i][j-1] = '|'
					tree[i][j+1] = '|'
					if visited[j] != 0 {
						v := visited[j]
						visited[j-1] += v
						visited[j+1] += v
						visited[j] = 0
					}
				}
			case 'S':
				tree[i+1][j] = '|'
				visited[j] = 1
			default:
				if i > 0 && tree[i-1][j] == '|' {
					tree[i][j] = '|'
				}
			}
		}
	}
	total := 0
	for _, count := range visited {
		total += count
	}

	return total
}
