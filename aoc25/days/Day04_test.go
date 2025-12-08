package days

import (
	"os"
	"strings"
	"testing"
)

var d04Input = []string{
	"..@@.@@@@.",
	"@@@.@.@.@@",
	"@@@@@.@.@@",
	"@.@@@@..@.",
	"@@.@@@@.@@",
	".@@@@@@@.@",
	".@.@.@.@@@",
	"@.@@@.@@@@",
	".@@@@@@@@.",
	"@.@.@@@.@.",
}

func TestD04P01(t *testing.T) {
	d := &Day04{}
	want := 13
	got := d.P01(d04Input)
	if want != got {
		t.Errorf("wanted %d got %d", want, got)
	}
}

func TestD04P02(t *testing.T) {
	d := &Day04{}
	want := 43
	got := d.P02(d04Input)
	if want != got {
		t.Errorf("wanted %d got %d", want, got)
	}
}

func BenchmarkDay04P01(b *testing.B) {
	data, err := os.ReadFile("./Day04.txt")
	if err != nil {
		b.Fatalf("failed to read file: %s", err)
	}

	input := strings.Split(string(data), "\n")
	d := &Day04{}

	for b.Loop() {
		d.P01(input)
	}
}

func BenchmarkDay04P02(b *testing.B) {
	data, err := os.ReadFile("./Day04.txt")
	if err != nil {
		b.Fatalf("failed to read file: %s", err)
	}

	input := strings.Split(string(data), "\n")
	d := &Day04{}

	for b.Loop() {
		d.P02(input)
	}
}
