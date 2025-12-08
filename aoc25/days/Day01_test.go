package days

import "testing"

var d01Input = []string{
	"L68",
	"L30",
	"R48",
	"L5",
	"R60",
	"L55",
	"L1",
	"L99",
	"R14",
	"L82",
}

func TestD01P01(t *testing.T) {
	d := &Day01{}
	want := 3
	got := d.P01(d01Input)
	if want != got {
		t.Errorf("wanted %d got %d", want, got)
	}
}

func TestD01P02(t *testing.T) {
	d := &Day01{}
	want := 6
	got := d.P02(d01Input)
	if want != got {
		t.Errorf("wanted %d got %d", want, got)
	}
}
