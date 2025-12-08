package days

import (
	"testing"
)

var d02Input = []string{"11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528," +
	"446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"}

func TestD02P01(t *testing.T) {
	d := &Day02{}
	want := 1227775554
	got := d.P01(d02Input)
	if want != got {
		t.Errorf("wanted %d got %d", want, got)
	}
}

func TestD02P02(t *testing.T) {
	d := &Day02{}
	want := 4174379265
	got := d.P02(d02Input)
	if want != got {
		t.Errorf("wanted %d got %d", want, got)
	}
}

func BenchmarkIsValidIDv2_Repetitive(b *testing.B) {
	// 12121212 -> repeats "12"
	id := 12121212
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = isValidIDv2(id)
	}
}

func BenchmarkIsValidIDv2_NonRepetitive(b *testing.B) {
	// 123456789012 -> non-repetitive
	id := 123456789012
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = isValidIDv2(id)
	}
}
