package days

import "testing"

var d06Input = []string{
	"123 328  51 64 ",
	" 45 64  387 23 ",
	"  6 98  215 314",
	"*   +   *   +  ",
}

func TestD06P01(t *testing.T) {
	d := &Day06{}
	want := 4277556
	got := d.P01(d06Input)
	if want != got {
		t.Errorf("wanted %d got %d", want, got)
	}
}

func TestD06P02(t *testing.T) {
	d := &Day06{}
	want := 3263827
	got := d.P02(d06Input)
	if want != got {
		t.Errorf("wanted %d got %d", want, got)
	}
}
