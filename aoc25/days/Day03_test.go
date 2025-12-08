package days

import "testing"

var d03Input = []string{
	"987654321111111",
	"811111111111119",
	"234234234234278",
	"818181911112111",
}

func TestD32P01(t *testing.T) {
	d := &Day03{}
	want := 357
	got := d.P01(d03Input)
	if want != got {
		t.Errorf("wanted %d got %d", want, got)
	}
}

func TestD32P02(t *testing.T) {
	d := &Day03{}
	want := 3121910778619
	got := d.P02(d03Input)
	if want != got {
		t.Errorf("wanted %d got %d", want, got)
	}
}
