package days

import "testing"

var d07Input = []string{
	".......S.......",
	"...............",
	".......^.......",
	"...............",
	"......^.^......",
	"...............",
	".....^.^.^.....",
	"...............",
	"....^.^...^....",
	"...............",
	"...^.^...^.^...",
	"...............",
	"..^...^.....^..",
	"...............",
	".^.^.^.^.^...^.",
	"...............",
}

func TestD07P01(t *testing.T) {
	d := &Day07{}
	want := 21
	got := d.P01(d07Input)
	if want != got {
		t.Errorf("wanted %d got %d", want, got)
	}
}

func TestD07P02(t *testing.T) {
	d := &Day07{}
	want := 40
	got := d.P02(d07Input)
	if want != got {
		t.Errorf("wanted %d got %d", want, got)
	}
}
