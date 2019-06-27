package coda

import (
	"bytes"
	"encoding/hex"
	"io"
	"math/rand"
	"testing"

	"github.com/gitchander/asn1/der/random"
)

func TestTagSamples(t *testing.T) {

	f1 := func(tag int) Header {
		return Header{
			Class:      CLASS_UNIVERSAL,
			Tag:        tag,
			IsCompound: false,
		}
	}

	f2 := func(tag int) Header {
		return Header{
			Class:      CLASS_CONTEXT_SPECIFIC,
			Tag:        tag,
			IsCompound: false,
		}
	}

	samples := []sample{
		{f1(0), "00"},
		{f1(1), "01"},
		{f1(30), "1E"},
		{f1(31), "1F1F"},
		{f1(127), "1F7F"},
		{f1(128), "1F8100"},
		{f1(16383), "1FFF7F"},
		{f1(16384), "1F818000"},
		{f1(2097151), "1FFFFF7F"},
		{f1(2097152), "1F81808000"},
		{f1(268435455), "1FFFFFFF7F"},
		{f1(268435456), "1F8180808000"},
		{f1(34359738367), "1FFFFFFFFF7F"},
		{f1(34359738368), "1F818080808000"},

		{f2(0), "80"},
		{f2(1), "81"},
		{f2(30), "9E"},
		{f2(31), "9F1F"},
		{f2(127), "9F7F"},
		{f2(128), "9F8100"},
		{f2(16383), "9FFF7F"},
		{f2(16384), "9F818000"},
		{f2(2097151), "9FFFFF7F"},
		{f2(2097152), "9F81808000"},
		{f2(268435455), "9FFFFFFF7F"},
		{f2(268435456), "9F8180808000"},
		{f2(34359738367), "9FFFFFFFFF7F"},
		{f2(34359738368), "9F818080808000"},
	}

	testSamplesEncode(t, samples)
	testSamplesDecode(t, samples)
}

type sample struct {
	h Header
	s string
}

func testSamplesEncode(t *testing.T, samples []sample) {
	for _, sample := range samples {
		d1, err := EncodeHeader(nil, &(sample.h))
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

func testSamplesDecode(t *testing.T, samples []sample) {
	for _, sample := range samples {
		h1 := sample.h
		data, err := hex.DecodeString(sample.s)
		if err != nil {
			t.Fatal(err)
		}
		var h2 Header
		_, err = DecodeHeader(data, &h2)
		if err != nil {
			t.Fatal(err)
		}
		if !EqualHeaders(h1, h2) {
			t.Fatalf("header: %+v != %+v", h1, h2)
		}
	}
}

func TestTagRand(t *testing.T) {
	r := random.NewRandNow()
	for i := 0; i < 10000; i++ {
		var a Header
		randHeader(r, &a)
		//t.Logf("%+v", a)
		data, err := EncodeHeader(nil, &a)
		if err != nil {
			t.Fatal(err)
		}
		var b Header
		_, err = DecodeHeader(data, &b)
		if err != nil {
			t.Fatal(err)
		}
		if !EqualHeaders(a, b) {
			t.Fatalf("%+v != %+v", a, b)
		}
	}
}

func randHeader(r *rand.Rand, h *Header) {
	switch r.Intn(4) {
	case 0:
		h.Class = CLASS_UNIVERSAL
	case 1:
		h.Class = CLASS_APPLICATION
	case 2:
		h.Class = CLASS_CONTEXT_SPECIFIC
	case 3:
		h.Class = CLASS_PRIVATE
	}
	h.Tag = int(r.Int31() >> uint(r.Intn(31)))
	h.IsCompound = ((r.Int() & 1) == 1)
}

func TestHeaderMergeSplit(t *testing.T) {

	r := random.NewRandNow()

	as := make([]Header, 100)

	var (
		data []byte
		err  error
	)

	// Merge:
	for i := range as {
		var a Header
		randHeader(r, &a)
		data, err = EncodeHeader(data, &a)
		if err != nil {
			t.Fatal(err)
		}
		as[i] = a
	}

	//t.Logf("data len: %d, data: %X", len(data), data)

	// Split:
	bs := make([]Header, 0, len(as))
	for {
		var b Header
		data, err = DecodeHeader(data, &b)
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
		//t.Logf("%+v : %+v", a, b)
		if !EqualHeaders(a, b) {
			t.Fatalf("%+v != %+v", a, b)
		}
	}
}
