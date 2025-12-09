package days

import "testing"

var d08Input = []string{
	"162,817,812",
	"57,618,57",
	"906,360,560",
	"592,479,940",
	"352,342,300",
	"466,668,158",
	"542,29,236",
	"431,825,988",
	"739,650,466",
	"52,470,668",
	"216,146,977",
	"819,987,18",
	"117,168,530",
	"805,96,715",
	"346,949,466",
	"970,615,88",
	"941,993,340",
	"862,61,35",
	"984,92,344",
	"425,690,689",
}

func TestD08P01(t *testing.T) {
	d := &Day08{}
	want := 40
	got := d.P01(d08Input)
	if want != got {
		t.Errorf("wanted %d got %d", want, got)
	}
}

func TestD08P02(t *testing.T) {
	d := &Day08{}
	want := 40
	got := d.P02(d08Input)
	if want != got {
		t.Errorf("wanted %d got %d", want, got)
	}
}
