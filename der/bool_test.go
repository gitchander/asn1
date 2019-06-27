package der

import (
	"testing"
)

type boolSample struct {
	val  bool
	data []byte
}

var boolSamples = []boolSample{
	boolSample{false, []byte{0x01, 0x01, 0x00}},
	boolSample{true, []byte{0x01, 0x01, 0x01}},
	boolSample{true, []byte{0x01, 0x01, 0x0F}},
	boolSample{true, []byte{0x01, 0x01, 0xFF}},
}

func TestBool(t *testing.T) {

	var a, b bool

	for _, sample := range boolSamples {

		err := Unmarshal(sample.data, &a)
		if err != nil {
			t.Fatal(err.Error())
		}

		if a != sample.val {
			t.Fatalf("cmp1: %v != %v", a, sample.val)
		}

		data, err := Marshal(sample.val)
		if err != nil {
			t.Fatal(err.Error())
		}

		err = Unmarshal(data, &b)
		if err != nil {
			t.Fatal(err.Error())
		}

		if b != sample.val {
			t.Fatalf("cmp2: %v != %v", b, sample.val)
		}
	}
}
