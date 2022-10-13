package der

import (
	"testing"

	"github.com/gitchander/asn1/der/utils/random"
)

func TestUTCTime(t *testing.T) {
	r := random.NewRandNow()
	for i := 0; i < 1000; i++ {
		t1 := RandomUTCTime(r)
		data, err := encodeUTCTime(t1)
		if err != nil {
			t.Fatal(err)
		}
		t2, err := decodeUTCTime(data)
		if err != nil {
			t.Fatal(err)
		}
		//t.Log(t1, "|", t2)
		if !t1.Equal(t2) {
			t.Fatalf("(%v) != (%v)", t1, t2)
		}
	}
}

func TestYearExpandFor(t *testing.T) {
	for i := -50; i < 0; i++ {
		year := yearExpand(i)
		x := -1
		if year != x {
			t.Fatalf("%d != %d", year, x)
		}
	}
	for i := 0; i < 50; i++ {
		year := yearExpand(i)
		x := i + 2000
		if year != x {
			t.Fatalf("%d != %d", year, x)
		}
	}
	for i := 50; i < 100; i++ {
		year := yearExpand(i)
		x := i + 1900
		if year != x {
			t.Fatalf("%d != %d", year, x)
		}
	}
	for i := 100; i < 150; i++ {
		year := yearExpand(i)
		x := -1
		if year != x {
			t.Fatalf("%d != %d", year, x)
		}
	}
}

type intRange struct {
	min, max int
}

func TestYearExpand(t *testing.T) {
	samples := []struct {
		r intRange
		f func(i int) int
	}{
		{
			intRange{-50, 0},
			func(i int) int { return -1 },
		},
		{
			intRange{0, 50},
			func(i int) int { return i + 2000 },
		},
		{
			intRange{50, 100},
			func(i int) int { return i + 1900 },
		},
		{
			intRange{100, 150},
			func(i int) int { return -1 },
		},
	}
	for _, sample := range samples {
		for i := sample.r.min; i < sample.r.max; i++ {
			year := yearExpand(i)
			x := sample.f(i)
			//t.Logf("%d: %d, %d", i, year, x)
			if year != x {
				t.Fatalf("%d != %d", year, x)
			}
		}
	}
}

func TestYearCollapse(t *testing.T) {
	samples := []struct {
		r intRange
		f func(i int) int
	}{
		{
			intRange{1900, 1950},
			func(i int) int { return -1 },
		},
		{
			intRange{1950, 2000},
			func(i int) int { return i - 1900 },
		},
		{
			intRange{2000, 2050},
			func(i int) int { return i - 2000 },
		},
		{
			intRange{2050, 2100},
			func(i int) int { return -1 },
		},
	}
	for _, sample := range samples {
		for i := sample.r.min; i < sample.r.max; i++ {
			year := yearCollapse(i)
			x := sample.f(i)
			//t.Logf("%d: %d, %d", i, year, x)
			if year != x {
				t.Fatalf("%d != %d", year, x)
			}
		}
	}
}
