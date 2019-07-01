package der

import (
	"testing"
)

func TestBoolSamples(t *testing.T) {

	var samples = []struct {
		val  bool
		data []byte
	}{
		{false, []byte{0x01, 0x01, 0x00}},
		{true, []byte{0x01, 0x01, 0x01}},
		{true, []byte{0x01, 0x01, 0x0F}},
		{true, []byte{0x01, 0x01, 0xFF}},
	}

	var a, b bool

	for _, sample := range samples {

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
