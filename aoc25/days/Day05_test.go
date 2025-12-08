package days

import "testing"

var d05Input = []string{
	"3-5",
	"10-14",
	"16-20",
	"12-18",
	"",
	"1",
	"5",
	"8",
	"11",
	"17",
	"32",
}

func TestD05P01(t *testing.T) {
	d := &Day05{}
	want := 3
	got := d.P01(d05Input)
	if want != got {
		t.Errorf("wanted %d got %d", want, got)
	}
}

func TestD05P02(t *testing.T) {
	d := &Day05{}
	want := 14
	got := d.P02(d05Input)
	if want != got {
		t.Errorf("wanted %d got %d", want, got)
	}
}
