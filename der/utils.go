package der

const (
	// bytesPerUint8  = 1
	// bytesPerUint16 = 2
	// bytesPerUint32 = 4
	bytesPerUint64 = 8
)

const (
	maxUint = ^uint(0)
	maxInt  = int(maxUint >> 1)
	minInt  = -maxInt - 1
)

func byteIsDigit(b byte) bool {
	return ('0' <= b) && (b <= '9')
}

func byteToDigit(b byte) (digit int, ok bool) {
	if byteIsDigit(b) {
		digit = int(b - '0')
		return digit, true
	}
	return 0, false
}

func digitToByte(digit int) (b byte, ok bool) {
	if (0 <= digit) && (digit <= 9) {
		b = byte('0' + digit)
		return b, true
	}
	return 0, false
}

func cloneBytes(a []byte) []byte {
	b := make([]byte, len(a))
	copy(b, a)
	return b
}
