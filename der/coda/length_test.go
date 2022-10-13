package coda

import (
	"bytes"
	"encoding/hex"
	"io"
	"math/rand"
	"testing"

	"github.com/gitchander/asn1/der/utils/random"
)

func randLength(r *rand.Rand) int {
	return int(r.Int() >> uint(r.Intn(63)))
}

func TestLengthSamples(t *testing.T) {
	samples := []struct {
		v int
		s string
	}{
		{0, "00"},
		{1, "01"},
		{15, "0F"},
		{16, "10"},
		{127, "7F"},
		{128, "8180"},
		{1<<8 - 1, "81FF"},
		{1 << 8, "820100"},
		{1<<16 - 1, "82FFFF"},
		{1 << 16, "83010000"},
		{1<<24 - 1, "83FFFFFF"},
		{1 << 24, "8401000000"},
		{1<<32 - 1, "84FFFFFFFF"},
		{1 << 32, "850100000000"},
		{1<<40 - 1, "85FFFFFFFFFF"},
		{1 << 40, "86010000000000"},
		{1<<48 - 1, "86FFFFFFFFFFFF"},
		{1 << 48, "8701000000000000"},
		{1<<56 - 1, "87FFFFFFFFFFFFFF"},
		{1 << 56, "880100000000000000"},
		//		{1<<64 - 1, "88FFFFFFFFFFFFFFFF"},
		//		{1<<64 + 0, "89010000000000000000"},
	}
	for _, sample := range samples {
		d1, err := EncodeLength(nil, sample.v)
		if err != nil {
			t.Fatal(err)
		}
		d2, err := hex.DecodeString(sample.s)
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(d1, d2) {
			t.Fatalf("data: [%X] != [%X]", d1, d2)
		}
	}
}

func TestLengthRand(t *testing.T) {
	r := random.NewRandNow()
	for i := 0; i < 1000; i++ {
		a := randLength(r)
		data, err := EncodeLength(nil, a)
		if err != nil {
			t.Fatal(err)
		}
		var b int
		_, err = DecodeLength(data, &b)
		if err != nil {
			t.Fatal(err)
		}
		if a != b {
			t.Fatalf("length %d != %d", a, b)
		}
	}
}

func TestLengthMergeSplit(t *testing.T) {

	r := random.NewRandNow()

	as := make([]int, 100)

	var (
		data []byte
		err  error
	)

	// Merge:
	for i := range as {
		a := randLength(r)
		data, err = EncodeLength(data, a)
		if err != nil {
			t.Fatal(err)
		}
		as[i] = a
	}

	// Split:
	bs := make([]int, 0, len(as))
	for {
		var b int
		data, err = DecodeLength(data, &b)
		if err != nil {
			if err == io.EOF {
				break
			}
			t.Fatal(err)
		}
		bs = append(bs, b)
	}

	if len(bs) != len(as) {
		t.Fatalf("not equal length: %d != %d", len(bs), len(as))
	}

	for i := range as {
		var (
			a = as[i]
			b = bs[i]
		)
		//t.Logf("%d : %d", a, b)
		if a != b {
			t.Fatalf("length: %d != %d", a, b)
		}
	}
}
