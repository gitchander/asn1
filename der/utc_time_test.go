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

type intRange struct {
	min, max int
}

func TestYearExpand(t *testing.T) {
	samples := []struct {
		r intRange
		f func(i int) (int, bool)
	}{
		{
			intRange{-50, 0},
			func(i int) (int, bool) { return 0, false },
		},
		{
			intRange{0, 50},
			func(i int) (int, bool) { return i + 2000, true },
		},
		{
			intRange{50, 100},
			func(i int) (int, bool) { return i + 1900, true },
		},
		{
			intRange{100, 150},
			func(i int) (int, bool) { return 0, false },
		},
	}
	for _, sample := range samples {
		for i := sample.r.min; i < sample.r.max; i++ {
			var (
				haveYear, haveOK = yearExpand(i)
				wantYear, wantOK = sample.f(i)
			)
			if haveYear != wantYear {
				t.Fatalf("invalid %s: have %d, want %d", "year", haveYear, wantYear)
			}
			if haveOK != wantOK {
				t.Fatalf("invalid %s: have %t, want %t", "ok", haveOK, wantOK)
			}
		}
	}
}

func TestYearCollapse(t *testing.T) {
	samples := []struct {
		r intRange
		f func(i int) (int, bool)
	}{
		{
			intRange{1900, 1950},
			func(i int) (int, bool) { return 0, false },
		},
		{
			intRange{1950, 2000},
			func(i int) (int, bool) { return i - 1900, true },
		},
		{
			intRange{2000, 2050},
			func(i int) (int, bool) { return i - 2000, true },
		},
		{
			intRange{2050, 2100},
			func(i int) (int, bool) { return 0, false },
		},
	}
	for _, sample := range samples {
		for i := sample.r.min; i < sample.r.max; i++ {
			var (
				haveYear, haveOK = yearCollapse(i)
				wantYear, wantOK = sample.f(i)
			)
			if haveYear != wantYear {
				t.Fatalf("invalid %s: have %d, want %d", "year", haveYear, wantYear)
			}
			if haveOK != wantOK {
				t.Fatalf("invalid %s: have %t, want %t", "ok", haveOK, wantOK)
			}
		}
	}
}
