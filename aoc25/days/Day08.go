package days

import (
	"fmt"
	"math"
	"slices"
	"sort"
	"strings"
)

type Day08 struct{}

func (*Day08) P01(input []string) int {
	var points []Point
	for i := range input {
		s := strings.Split(input[i], ",")
		points = append(points, Point{
			X: toFloat(s[0]),
			Y: toFloat(s[1]),
			Z: toFloat(s[2]),
		})
	}

	distances := []Junction{}
	for x := range points {
		minDist := math.MaxFloat64
		var p Point
		for y := range points {
			if x == y {
				continue
			}
			dist := points[x].Distance(points[y])
			if dist < minDist {
				minDist = dist
				p = points[y]
			}
		}
		np := []Point{points[x], p}
		ax := false
		for d := range distances {
			if containsAll(distances[d].points, np) {
				ax = true
				break
			}
		}
		if !ax {
			distances = append(distances, Junction{distance: minDist, points: np})
		}
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].distance < distances[j].distance
	})

	circuts := [][]Point{}
	for i := range 10 {
		p := distances[i]
		exists := false
		for i := range circuts {
			if slices.Contains(circuts[i], p.points[0]) && slices.Contains(circuts[i], p.points[1]) {
				exists = true
				break
			}
			if slices.Contains(circuts[i], p.points[0]) && !anyContains(circuts, p.points[1]) {
				circuts[i] = append(circuts[i], p.points[1])
				exists = true
				break
			}
			if slices.Contains(circuts[i], p.points[1]) && !anyContains(circuts, p.points[0]) {
				circuts[i] = append(circuts[i], p.points[0])
				exists = true
				break
			}
		}
		if exists {
			continue
		}
		circuts = append(circuts, p.points)

	}
	slices.SortFunc(circuts, func(a, b []Point) int { return len(b) - len(a) })
	total := 1
	for i := range 3 {
		total *= len(circuts[i])
	}
	fmt.Println(circuts)

	return total
}

func (*Day08) P02(input []string) int {
	return 0
}

type Point struct {
	X float64
	Y float64
	Z float64
}

func (p *Point) Distance(o Point) float64 {
	x := p.X - o.X
	y := p.Y - o.Y
	z := p.Z - o.Z
	return math.Sqrt(x*x + y*y + z*z)
}

type Junction struct {
	distance float64
	points   []Point
}

func anyContains(s [][]Point, p Point) bool {
	for i := range s {
		if slices.Contains(s[i], p) {
			return true
		}
	}
	return false
}

func containsAll(s []Point, p []Point) bool {
	for _, x := range p {
		if !slices.Contains(s, x) {
			return false
		}
	}
	return true
}
