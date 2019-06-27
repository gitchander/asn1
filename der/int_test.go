package der

import (
	"bytes"
	"encoding/hex"
	"errors"
	"math/rand"
	"testing"

	"github.com/gitchander/asn1/der/random"
)

func byteIsHex(b byte) bool {

	if (b >= '0') && (b <= '9') {
		return true
	}

	if (b >= 'a') && (b <= 'f') {
		return true
	}

	if (b >= 'A') && (b <= 'F') {
		return true
	}

	return false
}

func onlyHex(s string) string {

	data := []byte(s)

	var res []byte
	for _, b := range data {
		if byteIsHex(b) {
			res = append(res, b)
		}
	}

	return string(res)
}

type int64Sample struct {
	val  int64
	data []byte
}

func newInt64Sample(v int64, s string) *int64Sample {
	s = onlyHex(s)
	data, err := hex.DecodeString(s)
	if err != nil {
		panic(err.Error())
	}
	return &int64Sample{v, data}
}

var int64Samples = []*int64Sample{
	newInt64Sample(0, "02 01 00"),
	newInt64Sample(1, "02 01 01"),
	newInt64Sample(-1, "02 01 FF"),
	newInt64Sample(2, "02 01 02"),
	newInt64Sample(-2, "02 01 FE"),
	newInt64Sample(3, "02 01 03"),
	newInt64Sample(-3, "02 01 FD"),
	newInt64Sample(4, "02 01 04"),
	newInt64Sample(-4, "02 01 FC"),
	newInt64Sample(5, "02 01 05"),
	newInt64Sample(-5, "02 01 FB"),
	newInt64Sample(7, "02 01 07"),
	newInt64Sample(-7, "02 01 F9"),
	newInt64Sample(8, "02 01 08"),
	newInt64Sample(-8, "02 01 F8"),
	newInt64Sample(9, "02 01 09"),
	newInt64Sample(-9, "02 01 F7"),
	newInt64Sample(15, "02 01 0F"),
	newInt64Sample(-15, "02 01 F1"),
	newInt64Sample(16, "02 01 10"),
	newInt64Sample(-16, "02 01 F0"),
	newInt64Sample(17, "02 01 11"),
	newInt64Sample(-17, "02 01 EF"),
	newInt64Sample(31, "02 01 1F"),
	newInt64Sample(-31, "02 01 E1"),
	newInt64Sample(32, "02 01 20"),
	newInt64Sample(-32, "02 01 E0"),
	newInt64Sample(33, "02 01 21"),
	newInt64Sample(-33, "02 01 DF"),
	newInt64Sample(63, "02 01 3F"),
	newInt64Sample(-63, "02 01 C1"),
	newInt64Sample(64, "02 01 40"),
	newInt64Sample(-64, "02 01 C0"),
	newInt64Sample(65, "02 01 41"),
	newInt64Sample(-65, "02 01 BF"),
	newInt64Sample(127, "02 01 7F"),
	newInt64Sample(-127, "02 01 81"),
	newInt64Sample(128, "02 02 00 80"),
	newInt64Sample(-128, "02 01 80"),
	newInt64Sample(129, "02 02 00 81"),
	newInt64Sample(-129, "02 02 FF 7F"),
	newInt64Sample(255, "02 02 00 FF"),
	newInt64Sample(-255, "02 02 FF 01"),
	newInt64Sample(256, "02 02 01 00"),
	newInt64Sample(-256, "02 02 FF 00"),
	newInt64Sample(257, "02 02 01 01"),
	newInt64Sample(-257, "02 02 FE FF"),
	newInt64Sample(511, "02 02 01 FF"),
	newInt64Sample(-511, "02 02 FE 01"),
	newInt64Sample(512, "02 02 02 00"),
	newInt64Sample(-512, "02 02 FE 00"),
	newInt64Sample(513, "02 02 02 01"),
	newInt64Sample(-513, "02 02 FD FF"),
	newInt64Sample(1023, "02 02 03 FF"),
	newInt64Sample(-1023, "02 02 FC 01"),
	newInt64Sample(1024, "02 02 04 00"),
	newInt64Sample(-1024, "02 02 FC 00"),
	newInt64Sample(1025, "02 02 04 01"),
	newInt64Sample(-1025, "02 02 FB FF"),
	newInt64Sample(2047, "02 02 07 FF"),
	newInt64Sample(-2047, "02 02 F8 01"),
	newInt64Sample(2048, "02 02 08 00"),
	newInt64Sample(-2048, "02 02 F8 00"),
	newInt64Sample(2049, "02 02 08 01"),
	newInt64Sample(-2049, "02 02 F7 FF"),
	newInt64Sample(4095, "02 02 0F FF"),
	newInt64Sample(-4095, "02 02 F0 01"),
	newInt64Sample(4096, "02 02 10 00"),
	newInt64Sample(-4096, "02 02 F0 00"),
	newInt64Sample(4097, "02 02 10 01"),
	newInt64Sample(-4097, "02 02 EF FF"),
	newInt64Sample(8191, "02 02 1F FF"),
	newInt64Sample(-8191, "02 02 E0 01"),
	newInt64Sample(8192, "02 02 20 00"),
	newInt64Sample(-8192, "02 02 E0 00"),
	newInt64Sample(8193, "02 02 20 01"),
	newInt64Sample(-8193, "02 02 DF FF"),
	newInt64Sample(16383, "02 02 3F FF"),
	newInt64Sample(-16383, "02 02 C0 01"),
	newInt64Sample(16384, "02 02 40 00"),
	newInt64Sample(-16384, "02 02 C0 00"),
	newInt64Sample(16385, "02 02 40 01"),
	newInt64Sample(-16385, "02 02 BF FF"),
	newInt64Sample(32767, "02 02 7F FF"),
	newInt64Sample(-32767, "02 02 80 01"),
	newInt64Sample(32768, "02 03 00 80 00"),
	newInt64Sample(-32768, "02 02 80 00"),
	newInt64Sample(32769, "02 03 00 80 01"),
	newInt64Sample(-32769, "02 03 FF 7F FF"),
	newInt64Sample(65535, "02 03 00 FF FF"),
	newInt64Sample(-65535, "02 03 FF 00 01"),
	newInt64Sample(65536, "02 03 01 00 00"),
	newInt64Sample(-65536, "02 03 FF 00 00"),
	newInt64Sample(65537, "02 03 01 00 01"),
	newInt64Sample(-65537, "02 03 FE FF FF"),
	newInt64Sample(131071, "02 03 01 FF FF"),
	newInt64Sample(-131071, "02 03 FE 00 01"),
	newInt64Sample(131072, "02 03 02 00 00"),
	newInt64Sample(-131072, "02 03 FE 00 00"),
	newInt64Sample(131073, "02 03 02 00 01"),
	newInt64Sample(-131073, "02 03 FD FF FF"),
	newInt64Sample(262143, "02 03 03 FF FF"),
	newInt64Sample(-262143, "02 03 FC 00 01"),
	newInt64Sample(262144, "02 03 04 00 00"),
	newInt64Sample(-262144, "02 03 FC 00 00"),
	newInt64Sample(262145, "02 03 04 00 01"),
	newInt64Sample(-262145, "02 03 FB FF FF"),
	newInt64Sample(524287, "02 03 07 FF FF"),
	newInt64Sample(-524287, "02 03 F8 00 01"),
	newInt64Sample(524288, "02 03 08 00 00"),
	newInt64Sample(-524288, "02 03 F8 00 00"),
	newInt64Sample(524289, "02 03 08 00 01"),
	newInt64Sample(-524289, "02 03 F7 FF FF"),
	newInt64Sample(1048575, "02 03 0F FF FF"),
	newInt64Sample(-1048575, "02 03 F0 00 01"),
	newInt64Sample(1048576, "02 03 10 00 00"),
	newInt64Sample(-1048576, "02 03 F0 00 00"),
	newInt64Sample(1048577, "02 03 10 00 01"),
	newInt64Sample(-1048577, "02 03 EF FF FF"),
	newInt64Sample(2097151, "02 03 1F FF FF"),
	newInt64Sample(-2097151, "02 03 E0 00 01"),
	newInt64Sample(2097152, "02 03 20 00 00"),
	newInt64Sample(-2097152, "02 03 E0 00 00"),
	newInt64Sample(2097153, "02 03 20 00 01"),
	newInt64Sample(-2097153, "02 03 DF FF FF"),
	newInt64Sample(4194303, "02 03 3F FF FF"),
	newInt64Sample(-4194303, "02 03 C0 00 01"),
	newInt64Sample(4194304, "02 03 40 00 00"),
	newInt64Sample(-4194304, "02 03 C0 00 00"),
	newInt64Sample(4194305, "02 03 40 00 01"),
	newInt64Sample(-4194305, "02 03 BF FF FF"),
	newInt64Sample(8388607, "02 03 7F FF FF"),
	newInt64Sample(-8388607, "02 03 80 00 01"),
	newInt64Sample(8388608, "02 04 00 80 00 00"),
	newInt64Sample(-8388608, "02 03 80 00 00"),
	newInt64Sample(8388609, "02 04 00 80 00 01"),
	newInt64Sample(-8388609, "02 04 FF 7F FF FF"),
	newInt64Sample(16777215, "02 04 00 FF FF FF"),
	newInt64Sample(-16777215, "02 04 FF 00 00 01"),
	newInt64Sample(16777216, "02 04 01 00 00 00"),
	newInt64Sample(-16777216, "02 04 FF 00 00 00"),
	newInt64Sample(16777217, "02 04 01 00 00 01"),
	newInt64Sample(-16777217, "02 04 FE FF FF FF"),
	newInt64Sample(33554431, "02 04 01 FF FF FF"),
	newInt64Sample(-33554431, "02 04 FE 00 00 01"),
	newInt64Sample(33554432, "02 04 02 00 00 00"),
	newInt64Sample(-33554432, "02 04 FE 00 00 00"),
	newInt64Sample(33554433, "02 04 02 00 00 01"),
	newInt64Sample(-33554433, "02 04 FD FF FF FF"),
	newInt64Sample(67108863, "02 04 03 FF FF FF"),
	newInt64Sample(-67108863, "02 04 FC 00 00 01"),
	newInt64Sample(67108864, "02 04 04 00 00 00"),
	newInt64Sample(-67108864, "02 04 FC 00 00 00"),
	newInt64Sample(67108865, "02 04 04 00 00 01"),
	newInt64Sample(-67108865, "02 04 FB FF FF FF"),
	newInt64Sample(134217727, "02 04 07 FF FF FF"),
	newInt64Sample(-134217727, "02 04 F8 00 00 01"),
	newInt64Sample(134217728, "02 04 08 00 00 00"),
	newInt64Sample(-134217728, "02 04 F8 00 00 00"),
	newInt64Sample(134217729, "02 04 08 00 00 01"),
	newInt64Sample(-134217729, "02 04 F7 FF FF FF"),
	newInt64Sample(268435455, "02 04 0F FF FF FF"),
	newInt64Sample(-268435455, "02 04 F0 00 00 01"),
	newInt64Sample(268435456, "02 04 10 00 00 00"),
	newInt64Sample(-268435456, "02 04 F0 00 00 00"),
	newInt64Sample(268435457, "02 04 10 00 00 01"),
	newInt64Sample(-268435457, "02 04 EF FF FF FF"),
	newInt64Sample(536870911, "02 04 1F FF FF FF"),
	newInt64Sample(-536870911, "02 04 E0 00 00 01"),
	newInt64Sample(536870912, "02 04 20 00 00 00"),
	newInt64Sample(-536870912, "02 04 E0 00 00 00"),
	newInt64Sample(536870913, "02 04 20 00 00 01"),
	newInt64Sample(-536870913, "02 04 DF FF FF FF"),
	newInt64Sample(1073741823, "02 04 3F FF FF FF"),
	newInt64Sample(-1073741823, "02 04 C0 00 00 01"),
	newInt64Sample(1073741824, "02 04 40 00 00 00"),
	newInt64Sample(-1073741824, "02 04 C0 00 00 00"),
	newInt64Sample(1073741825, "02 04 40 00 00 01"),
	newInt64Sample(-1073741825, "02 04 BF FF FF FF"),
	newInt64Sample(2147483647, "02 04 7F FF FF FF"),
	newInt64Sample(-2147483647, "02 04 80 00 00 01"),
	newInt64Sample(2147483648, "02 05 00 80 00 00 00"),
	newInt64Sample(-2147483648, "02 04 80 00 00 00"),
	newInt64Sample(2147483649, "02 05 00 80 00 00 01"),
	newInt64Sample(-2147483649, "02 05 FF 7F FF FF FF"),
	newInt64Sample(4294967295, "02 05 00 FF FF FF FF"),
	newInt64Sample(-4294967295, "02 05 FF 00 00 00 01"),
	newInt64Sample(4294967296, "02 05 01 00 00 00 00"),
	newInt64Sample(-4294967296, "02 05 FF 00 00 00 00"),
	newInt64Sample(4294967297, "02 05 01 00 00 00 01"),
	newInt64Sample(-4294967297, "02 05 FE FF FF FF FF"),
	newInt64Sample(8589934591, "02 05 01 FF FF FF FF"),
	newInt64Sample(-8589934591, "02 05 FE 00 00 00 01"),
	newInt64Sample(8589934592, "02 05 02 00 00 00 00"),
	newInt64Sample(-8589934592, "02 05 FE 00 00 00 00"),
	newInt64Sample(8589934593, "02 05 02 00 00 00 01"),
	newInt64Sample(-8589934593, "02 05 FD FF FF FF FF"),
	newInt64Sample(17179869183, "02 05 03 FF FF FF FF"),
	newInt64Sample(-17179869183, "02 05 FC 00 00 00 01"),
	newInt64Sample(17179869184, "02 05 04 00 00 00 00"),
	newInt64Sample(-17179869184, "02 05 FC 00 00 00 00"),
	newInt64Sample(17179869185, "02 05 04 00 00 00 01"),
	newInt64Sample(-17179869185, "02 05 FB FF FF FF FF"),
	newInt64Sample(34359738367, "02 05 07 FF FF FF FF"),
	newInt64Sample(-34359738367, "02 05 F8 00 00 00 01"),
	newInt64Sample(34359738368, "02 05 08 00 00 00 00"),
	newInt64Sample(-34359738368, "02 05 F8 00 00 00 00"),
	newInt64Sample(34359738369, "02 05 08 00 00 00 01"),
	newInt64Sample(-34359738369, "02 05 F7 FF FF FF FF"),
	newInt64Sample(68719476735, "02 05 0F FF FF FF FF"),
	newInt64Sample(-68719476735, "02 05 F0 00 00 00 01"),
	newInt64Sample(68719476736, "02 05 10 00 00 00 00"),
	newInt64Sample(-68719476736, "02 05 F0 00 00 00 00"),
	newInt64Sample(68719476737, "02 05 10 00 00 00 01"),
	newInt64Sample(-68719476737, "02 05 EF FF FF FF FF"),
	newInt64Sample(137438953471, "02 05 1F FF FF FF FF"),
	newInt64Sample(-137438953471, "02 05 E0 00 00 00 01"),
	newInt64Sample(137438953472, "02 05 20 00 00 00 00"),
	newInt64Sample(-137438953472, "02 05 E0 00 00 00 00"),
	newInt64Sample(137438953473, "02 05 20 00 00 00 01"),
	newInt64Sample(-137438953473, "02 05 DF FF FF FF FF"),
	newInt64Sample(274877906943, "02 05 3F FF FF FF FF"),
	newInt64Sample(-274877906943, "02 05 C0 00 00 00 01"),
	newInt64Sample(274877906944, "02 05 40 00 00 00 00"),
	newInt64Sample(-274877906944, "02 05 C0 00 00 00 00"),
	newInt64Sample(274877906945, "02 05 40 00 00 00 01"),
	newInt64Sample(-274877906945, "02 05 BF FF FF FF FF"),
	newInt64Sample(549755813887, "02 05 7F FF FF FF FF"),
	newInt64Sample(-549755813887, "02 05 80 00 00 00 01"),
	newInt64Sample(549755813888, "02 06 00 80 00 00 00 00"),
	newInt64Sample(-549755813888, "02 05 80 00 00 00 00"),
	newInt64Sample(549755813889, "02 06 00 80 00 00 00 01"),
	newInt64Sample(-549755813889, "02 06 FF 7F FF FF FF FF"),
	newInt64Sample(1099511627775, "02 06 00 FF FF FF FF FF"),
	newInt64Sample(-1099511627775, "02 06 FF 00 00 00 00 01"),
	newInt64Sample(1099511627776, "02 06 01 00 00 00 00 00"),
	newInt64Sample(-1099511627776, "02 06 FF 00 00 00 00 00"),
	newInt64Sample(1099511627777, "02 06 01 00 00 00 00 01"),
	newInt64Sample(-1099511627777, "02 06 FE FF FF FF FF FF"),
	newInt64Sample(2199023255551, "02 06 01 FF FF FF FF FF"),
	newInt64Sample(-2199023255551, "02 06 FE 00 00 00 00 01"),
	newInt64Sample(2199023255552, "02 06 02 00 00 00 00 00"),
	newInt64Sample(-2199023255552, "02 06 FE 00 00 00 00 00"),
	newInt64Sample(2199023255553, "02 06 02 00 00 00 00 01"),
	newInt64Sample(-2199023255553, "02 06 FD FF FF FF FF FF"),
	newInt64Sample(4398046511103, "02 06 03 FF FF FF FF FF"),
	newInt64Sample(-4398046511103, "02 06 FC 00 00 00 00 01"),
	newInt64Sample(4398046511104, "02 06 04 00 00 00 00 00"),
	newInt64Sample(-4398046511104, "02 06 FC 00 00 00 00 00"),
	newInt64Sample(4398046511105, "02 06 04 00 00 00 00 01"),
	newInt64Sample(-4398046511105, "02 06 FB FF FF FF FF FF"),
	newInt64Sample(8796093022207, "02 06 07 FF FF FF FF FF"),
	newInt64Sample(-8796093022207, "02 06 F8 00 00 00 00 01"),
	newInt64Sample(8796093022208, "02 06 08 00 00 00 00 00"),
	newInt64Sample(-8796093022208, "02 06 F8 00 00 00 00 00"),
	newInt64Sample(8796093022209, "02 06 08 00 00 00 00 01"),
	newInt64Sample(-8796093022209, "02 06 F7 FF FF FF FF FF"),
	newInt64Sample(17592186044415, "02 06 0F FF FF FF FF FF"),
	newInt64Sample(-17592186044415, "02 06 F0 00 00 00 00 01"),
	newInt64Sample(17592186044416, "02 06 10 00 00 00 00 00"),
	newInt64Sample(-17592186044416, "02 06 F0 00 00 00 00 00"),
	newInt64Sample(17592186044417, "02 06 10 00 00 00 00 01"),
	newInt64Sample(-17592186044417, "02 06 EF FF FF FF FF FF"),
	newInt64Sample(35184372088831, "02 06 1F FF FF FF FF FF"),
	newInt64Sample(-35184372088831, "02 06 E0 00 00 00 00 01"),
	newInt64Sample(35184372088832, "02 06 20 00 00 00 00 00"),
	newInt64Sample(-35184372088832, "02 06 E0 00 00 00 00 00"),
	newInt64Sample(35184372088833, "02 06 20 00 00 00 00 01"),
	newInt64Sample(-35184372088833, "02 06 DF FF FF FF FF FF"),
	newInt64Sample(70368744177663, "02 06 3F FF FF FF FF FF"),
	newInt64Sample(-70368744177663, "02 06 C0 00 00 00 00 01"),
	newInt64Sample(70368744177664, "02 06 40 00 00 00 00 00"),
	newInt64Sample(-70368744177664, "02 06 C0 00 00 00 00 00"),
	newInt64Sample(70368744177665, "02 06 40 00 00 00 00 01"),
	newInt64Sample(-70368744177665, "02 06 BF FF FF FF FF FF"),
	newInt64Sample(140737488355327, "02 06 7F FF FF FF FF FF"),
	newInt64Sample(-140737488355327, "02 06 80 00 00 00 00 01"),
	newInt64Sample(140737488355328, "02 07 00 80 00 00 00 00 00"),
	newInt64Sample(-140737488355328, "02 06 80 00 00 00 00 00"),
	newInt64Sample(140737488355329, "02 07 00 80 00 00 00 00 01"),
	newInt64Sample(-140737488355329, "02 07 FF 7F FF FF FF FF FF"),
	newInt64Sample(281474976710655, "02 07 00 FF FF FF FF FF FF"),
	newInt64Sample(-281474976710655, "02 07 FF 00 00 00 00 00 01"),
	newInt64Sample(281474976710656, "02 07 01 00 00 00 00 00 00"),
	newInt64Sample(-281474976710656, "02 07 FF 00 00 00 00 00 00"),
	newInt64Sample(281474976710657, "02 07 01 00 00 00 00 00 01"),
	newInt64Sample(-281474976710657, "02 07 FE FF FF FF FF FF FF"),
	newInt64Sample(562949953421311, "02 07 01 FF FF FF FF FF FF"),
	newInt64Sample(-562949953421311, "02 07 FE 00 00 00 00 00 01"),
	newInt64Sample(562949953421312, "02 07 02 00 00 00 00 00 00"),
	newInt64Sample(-562949953421312, "02 07 FE 00 00 00 00 00 00"),
	newInt64Sample(562949953421313, "02 07 02 00 00 00 00 00 01"),
	newInt64Sample(-562949953421313, "02 07 FD FF FF FF FF FF FF"),
	newInt64Sample(1125899906842623, "02 07 03 FF FF FF FF FF FF"),
	newInt64Sample(-1125899906842623, "02 07 FC 00 00 00 00 00 01"),
	newInt64Sample(1125899906842624, "02 07 04 00 00 00 00 00 00"),
	newInt64Sample(-1125899906842624, "02 07 FC 00 00 00 00 00 00"),
	newInt64Sample(1125899906842625, "02 07 04 00 00 00 00 00 01"),
	newInt64Sample(-1125899906842625, "02 07 FB FF FF FF FF FF FF"),
	newInt64Sample(2251799813685247, "02 07 07 FF FF FF FF FF FF"),
	newInt64Sample(-2251799813685247, "02 07 F8 00 00 00 00 00 01"),
	newInt64Sample(2251799813685248, "02 07 08 00 00 00 00 00 00"),
	newInt64Sample(-2251799813685248, "02 07 F8 00 00 00 00 00 00"),
	newInt64Sample(2251799813685249, "02 07 08 00 00 00 00 00 01"),
	newInt64Sample(-2251799813685249, "02 07 F7 FF FF FF FF FF FF"),
	newInt64Sample(4503599627370495, "02 07 0F FF FF FF FF FF FF"),
	newInt64Sample(-4503599627370495, "02 07 F0 00 00 00 00 00 01"),
	newInt64Sample(4503599627370496, "02 07 10 00 00 00 00 00 00"),
	newInt64Sample(-4503599627370496, "02 07 F0 00 00 00 00 00 00"),
	newInt64Sample(4503599627370497, "02 07 10 00 00 00 00 00 01"),
	newInt64Sample(-4503599627370497, "02 07 EF FF FF FF FF FF FF"),
	newInt64Sample(9007199254740991, "02 07 1F FF FF FF FF FF FF"),
	newInt64Sample(-9007199254740991, "02 07 E0 00 00 00 00 00 01"),
	newInt64Sample(9007199254740992, "02 07 20 00 00 00 00 00 00"),
	newInt64Sample(-9007199254740992, "02 07 E0 00 00 00 00 00 00"),
	newInt64Sample(9007199254740993, "02 07 20 00 00 00 00 00 01"),
	newInt64Sample(-9007199254740993, "02 07 DF FF FF FF FF FF FF"),
	newInt64Sample(18014398509481983, "02 07 3F FF FF FF FF FF FF"),
	newInt64Sample(-18014398509481983, "02 07 C0 00 00 00 00 00 01"),
	newInt64Sample(18014398509481984, "02 07 40 00 00 00 00 00 00"),
	newInt64Sample(-18014398509481984, "02 07 C0 00 00 00 00 00 00"),
	newInt64Sample(18014398509481985, "02 07 40 00 00 00 00 00 01"),
	newInt64Sample(-18014398509481985, "02 07 BF FF FF FF FF FF FF"),
	newInt64Sample(36028797018963967, "02 07 7F FF FF FF FF FF FF"),
	newInt64Sample(-36028797018963967, "02 07 80 00 00 00 00 00 01"),
	newInt64Sample(36028797018963968, "02 08 00 80 00 00 00 00 00 00"),
	newInt64Sample(-36028797018963968, "02 07 80 00 00 00 00 00 00"),
	newInt64Sample(36028797018963969, "02 08 00 80 00 00 00 00 00 01"),
	newInt64Sample(-36028797018963969, "02 08 FF 7F FF FF FF FF FF FF"),
	newInt64Sample(72057594037927935, "02 08 00 FF FF FF FF FF FF FF"),
	newInt64Sample(-72057594037927935, "02 08 FF 00 00 00 00 00 00 01"),
	newInt64Sample(72057594037927936, "02 08 01 00 00 00 00 00 00 00"),
	newInt64Sample(-72057594037927936, "02 08 FF 00 00 00 00 00 00 00"),
	newInt64Sample(72057594037927937, "02 08 01 00 00 00 00 00 00 01"),
	newInt64Sample(-72057594037927937, "02 08 FE FF FF FF FF FF FF FF"),
	newInt64Sample(144115188075855871, "02 08 01 FF FF FF FF FF FF FF"),
	newInt64Sample(-144115188075855871, "02 08 FE 00 00 00 00 00 00 01"),
	newInt64Sample(144115188075855872, "02 08 02 00 00 00 00 00 00 00"),
	newInt64Sample(-144115188075855872, "02 08 FE 00 00 00 00 00 00 00"),
	newInt64Sample(144115188075855873, "02 08 02 00 00 00 00 00 00 01"),
	newInt64Sample(-144115188075855873, "02 08 FD FF FF FF FF FF FF FF"),
	newInt64Sample(288230376151711743, "02 08 03 FF FF FF FF FF FF FF"),
	newInt64Sample(-288230376151711743, "02 08 FC 00 00 00 00 00 00 01"),
	newInt64Sample(288230376151711744, "02 08 04 00 00 00 00 00 00 00"),
	newInt64Sample(-288230376151711744, "02 08 FC 00 00 00 00 00 00 00"),
	newInt64Sample(288230376151711745, "02 08 04 00 00 00 00 00 00 01"),
	newInt64Sample(-288230376151711745, "02 08 FB FF FF FF FF FF FF FF"),
	newInt64Sample(576460752303423487, "02 08 07 FF FF FF FF FF FF FF"),
	newInt64Sample(-576460752303423487, "02 08 F8 00 00 00 00 00 00 01"),
	newInt64Sample(576460752303423488, "02 08 08 00 00 00 00 00 00 00"),
	newInt64Sample(-576460752303423488, "02 08 F8 00 00 00 00 00 00 00"),
	newInt64Sample(576460752303423489, "02 08 08 00 00 00 00 00 00 01"),
	newInt64Sample(-576460752303423489, "02 08 F7 FF FF FF FF FF FF FF"),
	newInt64Sample(1152921504606846975, "02 08 0F FF FF FF FF FF FF FF"),
	newInt64Sample(-1152921504606846975, "02 08 F0 00 00 00 00 00 00 01"),
	newInt64Sample(1152921504606846976, "02 08 10 00 00 00 00 00 00 00"),
	newInt64Sample(-1152921504606846976, "02 08 F0 00 00 00 00 00 00 00"),
	newInt64Sample(1152921504606846977, "02 08 10 00 00 00 00 00 00 01"),
	newInt64Sample(-1152921504606846977, "02 08 EF FF FF FF FF FF FF FF"),
	newInt64Sample(2305843009213693951, "02 08 1F FF FF FF FF FF FF FF"),
	newInt64Sample(-2305843009213693951, "02 08 E0 00 00 00 00 00 00 01"),
	newInt64Sample(2305843009213693952, "02 08 20 00 00 00 00 00 00 00"),
	newInt64Sample(-2305843009213693952, "02 08 E0 00 00 00 00 00 00 00"),
	newInt64Sample(2305843009213693953, "02 08 20 00 00 00 00 00 00 01"),
	newInt64Sample(-2305843009213693953, "02 08 DF FF FF FF FF FF FF FF"),
	newInt64Sample(4611686018427387903, "02 08 3F FF FF FF FF FF FF FF"),
	newInt64Sample(-4611686018427387903, "02 08 C0 00 00 00 00 00 00 01"),
	newInt64Sample(4611686018427387904, "02 08 40 00 00 00 00 00 00 00"),
	newInt64Sample(-4611686018427387904, "02 08 C0 00 00 00 00 00 00 00"),
	newInt64Sample(4611686018427387905, "02 08 40 00 00 00 00 00 00 01"),
	newInt64Sample(-4611686018427387905, "02 08 BF FF FF FF FF FF FF FF"),
	newInt64Sample(9223372036854775807, "02 08 7F FF FF FF FF FF FF FF"),
	newInt64Sample(-9223372036854775807, "02 08 80 00 00 00 00 00 00 01"),
	newInt64Sample(-9223372036854775808, "02 08 80 00 00 00 00 00 00 00"),
}

func TestInt64Samples(t *testing.T) {

	for _, sample := range int64Samples {

		data, err := Marshal(sample.val)
		if err != nil {
			t.Fatal(err.Error())
		}

		if bytes.Compare(data, sample.data) != 0 {
			t.Fatal(errors.New("wrong compare data"))
		}

		var x int64
		err = Unmarshal(data, &x)
		if err != nil {
			t.Fatal(err.Error())
		}

		if sample.val != x {
			t.Fatal(errors.New("wrong unmarshal data"))
		}
	}
}

type uint64Sample struct {
	val  uint64
	data []byte
}

func newUint64Sample(v uint64, s string) *uint64Sample {
	s = onlyHex(s)
	data, err := hex.DecodeString(s)
	if err != nil {
		panic(err.Error())
	}
	return &uint64Sample{v, data}
}

var uint64Samples = []*uint64Sample{
	newUint64Sample(0, "02 01 00"),
	newUint64Sample(1, "02 01 01"),
	newUint64Sample(2, "02 01 02"),
	newUint64Sample(3, "02 01 03"),
	newUint64Sample(4, "02 01 04"),
	newUint64Sample(5, "02 01 05"),
	newUint64Sample(7, "02 01 07"),
	newUint64Sample(8, "02 01 08"),
	newUint64Sample(9, "02 01 09"),
	newUint64Sample(15, "02 01 0F"),
	newUint64Sample(16, "02 01 10"),
	newUint64Sample(17, "02 01 11"),
	newUint64Sample(31, "02 01 1F"),
	newUint64Sample(32, "02 01 20"),
	newUint64Sample(33, "02 01 21"),
	newUint64Sample(63, "02 01 3F"),
	newUint64Sample(64, "02 01 40"),
	newUint64Sample(65, "02 01 41"),
	newUint64Sample(127, "02 01 7F"),
	newUint64Sample(128, "02 02 00 80"),
	newUint64Sample(129, "02 02 00 81"),
	newUint64Sample(255, "02 02 00 FF"),
	newUint64Sample(256, "02 02 01 00"),
	newUint64Sample(257, "02 02 01 01"),
	newUint64Sample(511, "02 02 01 FF"),
	newUint64Sample(512, "02 02 02 00"),
	newUint64Sample(513, "02 02 02 01"),
	newUint64Sample(1023, "02 02 03 FF"),
	newUint64Sample(1024, "02 02 04 00"),
	newUint64Sample(1025, "02 02 04 01"),
	newUint64Sample(2047, "02 02 07 FF"),
	newUint64Sample(2048, "02 02 08 00"),
	newUint64Sample(2049, "02 02 08 01"),
	newUint64Sample(4095, "02 02 0F FF"),
	newUint64Sample(4096, "02 02 10 00"),
	newUint64Sample(4097, "02 02 10 01"),
	newUint64Sample(8191, "02 02 1F FF"),
	newUint64Sample(8192, "02 02 20 00"),
	newUint64Sample(8193, "02 02 20 01"),
	newUint64Sample(16383, "02 02 3F FF"),
	newUint64Sample(16384, "02 02 40 00"),
	newUint64Sample(16385, "02 02 40 01"),
	newUint64Sample(32767, "02 02 7F FF"),
	newUint64Sample(32768, "02 03 00 80 00"),
	newUint64Sample(32769, "02 03 00 80 01"),
	newUint64Sample(65535, "02 03 00 FF FF"),
	newUint64Sample(65536, "02 03 01 00 00"),
	newUint64Sample(65537, "02 03 01 00 01"),
	newUint64Sample(131071, "02 03 01 FF FF"),
	newUint64Sample(131072, "02 03 02 00 00"),
	newUint64Sample(131073, "02 03 02 00 01"),
	newUint64Sample(262143, "02 03 03 FF FF"),
	newUint64Sample(262144, "02 03 04 00 00"),
	newUint64Sample(262145, "02 03 04 00 01"),
	newUint64Sample(524287, "02 03 07 FF FF"),
	newUint64Sample(524288, "02 03 08 00 00"),
	newUint64Sample(524289, "02 03 08 00 01"),
	newUint64Sample(1048575, "02 03 0F FF FF"),
	newUint64Sample(1048576, "02 03 10 00 00"),
	newUint64Sample(1048577, "02 03 10 00 01"),
	newUint64Sample(2097151, "02 03 1F FF FF"),
	newUint64Sample(2097152, "02 03 20 00 00"),
	newUint64Sample(2097153, "02 03 20 00 01"),
	newUint64Sample(4194303, "02 03 3F FF FF"),
	newUint64Sample(4194304, "02 03 40 00 00"),
	newUint64Sample(4194305, "02 03 40 00 01"),
	newUint64Sample(8388607, "02 03 7F FF FF"),
	newUint64Sample(8388608, "02 04 00 80 00 00"),
	newUint64Sample(8388609, "02 04 00 80 00 01"),
	newUint64Sample(16777215, "02 04 00 FF FF FF"),
	newUint64Sample(16777216, "02 04 01 00 00 00"),
	newUint64Sample(16777217, "02 04 01 00 00 01"),
	newUint64Sample(33554431, "02 04 01 FF FF FF"),
	newUint64Sample(33554432, "02 04 02 00 00 00"),
	newUint64Sample(33554433, "02 04 02 00 00 01"),
	newUint64Sample(67108863, "02 04 03 FF FF FF"),
	newUint64Sample(67108864, "02 04 04 00 00 00"),
	newUint64Sample(67108865, "02 04 04 00 00 01"),
	newUint64Sample(134217727, "02 04 07 FF FF FF"),
	newUint64Sample(134217728, "02 04 08 00 00 00"),
	newUint64Sample(134217729, "02 04 08 00 00 01"),
	newUint64Sample(268435455, "02 04 0F FF FF FF"),
	newUint64Sample(268435456, "02 04 10 00 00 00"),
	newUint64Sample(268435457, "02 04 10 00 00 01"),
	newUint64Sample(536870911, "02 04 1F FF FF FF"),
	newUint64Sample(536870912, "02 04 20 00 00 00"),
	newUint64Sample(536870913, "02 04 20 00 00 01"),
	newUint64Sample(1073741823, "02 04 3F FF FF FF"),
	newUint64Sample(1073741824, "02 04 40 00 00 00"),
	newUint64Sample(1073741825, "02 04 40 00 00 01"),
	newUint64Sample(2147483647, "02 04 7F FF FF FF"),
	newUint64Sample(2147483648, "02 05 00 80 00 00 00"),
	newUint64Sample(2147483649, "02 05 00 80 00 00 01"),
	newUint64Sample(4294967295, "02 05 00 FF FF FF FF"),
	newUint64Sample(4294967296, "02 05 01 00 00 00 00"),
	newUint64Sample(4294967297, "02 05 01 00 00 00 01"),
	newUint64Sample(8589934591, "02 05 01 FF FF FF FF"),
	newUint64Sample(8589934592, "02 05 02 00 00 00 00"),
	newUint64Sample(8589934593, "02 05 02 00 00 00 01"),
	newUint64Sample(17179869183, "02 05 03 FF FF FF FF"),
	newUint64Sample(17179869184, "02 05 04 00 00 00 00"),
	newUint64Sample(17179869185, "02 05 04 00 00 00 01"),
	newUint64Sample(34359738367, "02 05 07 FF FF FF FF"),
	newUint64Sample(34359738368, "02 05 08 00 00 00 00"),
	newUint64Sample(34359738369, "02 05 08 00 00 00 01"),
	newUint64Sample(68719476735, "02 05 0F FF FF FF FF"),
	newUint64Sample(68719476736, "02 05 10 00 00 00 00"),
	newUint64Sample(68719476737, "02 05 10 00 00 00 01"),
	newUint64Sample(137438953471, "02 05 1F FF FF FF FF"),
	newUint64Sample(137438953472, "02 05 20 00 00 00 00"),
	newUint64Sample(137438953473, "02 05 20 00 00 00 01"),
	newUint64Sample(274877906943, "02 05 3F FF FF FF FF"),
	newUint64Sample(274877906944, "02 05 40 00 00 00 00"),
	newUint64Sample(274877906945, "02 05 40 00 00 00 01"),
	newUint64Sample(549755813887, "02 05 7F FF FF FF FF"),
	newUint64Sample(549755813888, "02 06 00 80 00 00 00 00"),
	newUint64Sample(549755813889, "02 06 00 80 00 00 00 01"),
	newUint64Sample(1099511627775, "02 06 00 FF FF FF FF FF"),
	newUint64Sample(1099511627776, "02 06 01 00 00 00 00 00"),
	newUint64Sample(1099511627777, "02 06 01 00 00 00 00 01"),
	newUint64Sample(2199023255551, "02 06 01 FF FF FF FF FF"),
	newUint64Sample(2199023255552, "02 06 02 00 00 00 00 00"),
	newUint64Sample(2199023255553, "02 06 02 00 00 00 00 01"),
	newUint64Sample(4398046511103, "02 06 03 FF FF FF FF FF"),
	newUint64Sample(4398046511104, "02 06 04 00 00 00 00 00"),
	newUint64Sample(4398046511105, "02 06 04 00 00 00 00 01"),
	newUint64Sample(8796093022207, "02 06 07 FF FF FF FF FF"),
	newUint64Sample(8796093022208, "02 06 08 00 00 00 00 00"),
	newUint64Sample(8796093022209, "02 06 08 00 00 00 00 01"),
	newUint64Sample(17592186044415, "02 06 0F FF FF FF FF FF"),
	newUint64Sample(17592186044416, "02 06 10 00 00 00 00 00"),
	newUint64Sample(17592186044417, "02 06 10 00 00 00 00 01"),
	newUint64Sample(35184372088831, "02 06 1F FF FF FF FF FF"),
	newUint64Sample(35184372088832, "02 06 20 00 00 00 00 00"),
	newUint64Sample(35184372088833, "02 06 20 00 00 00 00 01"),
	newUint64Sample(70368744177663, "02 06 3F FF FF FF FF FF"),
	newUint64Sample(70368744177664, "02 06 40 00 00 00 00 00"),
	newUint64Sample(70368744177665, "02 06 40 00 00 00 00 01"),
	newUint64Sample(140737488355327, "02 06 7F FF FF FF FF FF"),
	newUint64Sample(140737488355328, "02 07 00 80 00 00 00 00 00"),
	newUint64Sample(140737488355329, "02 07 00 80 00 00 00 00 01"),
	newUint64Sample(281474976710655, "02 07 00 FF FF FF FF FF FF"),
	newUint64Sample(281474976710656, "02 07 01 00 00 00 00 00 00"),
	newUint64Sample(281474976710657, "02 07 01 00 00 00 00 00 01"),
	newUint64Sample(562949953421311, "02 07 01 FF FF FF FF FF FF"),
	newUint64Sample(562949953421312, "02 07 02 00 00 00 00 00 00"),
	newUint64Sample(562949953421313, "02 07 02 00 00 00 00 00 01"),
	newUint64Sample(1125899906842623, "02 07 03 FF FF FF FF FF FF"),
	newUint64Sample(1125899906842624, "02 07 04 00 00 00 00 00 00"),
	newUint64Sample(1125899906842625, "02 07 04 00 00 00 00 00 01"),
	newUint64Sample(2251799813685247, "02 07 07 FF FF FF FF FF FF"),
	newUint64Sample(2251799813685248, "02 07 08 00 00 00 00 00 00"),
	newUint64Sample(2251799813685249, "02 07 08 00 00 00 00 00 01"),
	newUint64Sample(4503599627370495, "02 07 0F FF FF FF FF FF FF"),
	newUint64Sample(4503599627370496, "02 07 10 00 00 00 00 00 00"),
	newUint64Sample(4503599627370497, "02 07 10 00 00 00 00 00 01"),
	newUint64Sample(9007199254740991, "02 07 1F FF FF FF FF FF FF"),
	newUint64Sample(9007199254740992, "02 07 20 00 00 00 00 00 00"),
	newUint64Sample(9007199254740993, "02 07 20 00 00 00 00 00 01"),
	newUint64Sample(18014398509481983, "02 07 3F FF FF FF FF FF FF"),
	newUint64Sample(18014398509481984, "02 07 40 00 00 00 00 00 00"),
	newUint64Sample(18014398509481985, "02 07 40 00 00 00 00 00 01"),
	newUint64Sample(36028797018963967, "02 07 7F FF FF FF FF FF FF"),
	newUint64Sample(36028797018963968, "02 08 00 80 00 00 00 00 00 00"),
	newUint64Sample(36028797018963969, "02 08 00 80 00 00 00 00 00 01"),
	newUint64Sample(72057594037927935, "02 08 00 FF FF FF FF FF FF FF"),
	newUint64Sample(72057594037927936, "02 08 01 00 00 00 00 00 00 00"),
	newUint64Sample(72057594037927937, "02 08 01 00 00 00 00 00 00 01"),
	newUint64Sample(144115188075855871, "02 08 01 FF FF FF FF FF FF FF"),
	newUint64Sample(144115188075855872, "02 08 02 00 00 00 00 00 00 00"),
	newUint64Sample(144115188075855873, "02 08 02 00 00 00 00 00 00 01"),
	newUint64Sample(288230376151711743, "02 08 03 FF FF FF FF FF FF FF"),
	newUint64Sample(288230376151711744, "02 08 04 00 00 00 00 00 00 00"),
	newUint64Sample(288230376151711745, "02 08 04 00 00 00 00 00 00 01"),
	newUint64Sample(576460752303423487, "02 08 07 FF FF FF FF FF FF FF"),
	newUint64Sample(576460752303423488, "02 08 08 00 00 00 00 00 00 00"),
	newUint64Sample(576460752303423489, "02 08 08 00 00 00 00 00 00 01"),
	newUint64Sample(1152921504606846975, "02 08 0F FF FF FF FF FF FF FF"),
	newUint64Sample(1152921504606846976, "02 08 10 00 00 00 00 00 00 00"),
	newUint64Sample(1152921504606846977, "02 08 10 00 00 00 00 00 00 01"),
	newUint64Sample(2305843009213693951, "02 08 1F FF FF FF FF FF FF FF"),
	newUint64Sample(2305843009213693952, "02 08 20 00 00 00 00 00 00 00"),
	newUint64Sample(2305843009213693953, "02 08 20 00 00 00 00 00 00 01"),
	newUint64Sample(4611686018427387903, "02 08 3F FF FF FF FF FF FF FF"),
	newUint64Sample(4611686018427387904, "02 08 40 00 00 00 00 00 00 00"),
	newUint64Sample(4611686018427387905, "02 08 40 00 00 00 00 00 00 01"),
	newUint64Sample(9223372036854775807, "02 08 7F FF FF FF FF FF FF FF"),
	newUint64Sample(9223372036854775808, "02 09 00 80 00 00 00 00 00 00 00"),
	newUint64Sample(9223372036854775809, "02 09 00 80 00 00 00 00 00 00 01"),
	newUint64Sample(18446744073709551615, "02 09 00 FF FF FF FF FF FF FF FF"),
}

func TestUint64Samples(t *testing.T) {

	for _, sample := range uint64Samples {

		data, err := Marshal(sample.val)
		if err != nil {
			t.Fatal(err.Error())
		}

		if bytes.Compare(data, sample.data) != 0 {
			t.Fatal(errors.New("wrong compare data"))
		}

		var x uint64
		err = Unmarshal(data, &x)
		if err != nil {
			t.Fatal(err.Error())
		}

		if sample.val != x {
			t.Fatal(errors.New("wrong unmarshal data"))
		}
	}
}

func randInt64(r *rand.Rand) int64 {
	a := (r.Int63() >> uint(r.Intn(62)))
	if (r.Int() & 1) == 0 {
		a = -a
	}
	return a
}

func TestInt64Marshal(t *testing.T) {

	var a, b int64
	r := random.NewRandNow()

	for i := 0; i < 10000; i++ {

		a = randInt64(r)

		data, err := Marshal(a)
		if err != nil {
			t.Fatal(err.Error())
		}

		err = Unmarshal(data, &b)
		if err != nil {
			t.Fatal(err.Error())
		}

		if a != b {
			t.Fatalf("%d != %d", a, b)
		}
	}
}
