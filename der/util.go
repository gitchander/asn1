package der

import (
	"io"
)

const (
	//sizeOfUint8  = 1
	//sizeOfUint16 = 2
	//sizeOfUint32 = 4
	sizeOfUint64 = 8
)

const (
	maxUint = ^uint(0)
	maxInt  = int(maxUint >> 1)
	minInt  = -maxInt - 1
)

func writeByte(w io.Writer, b byte) error {
	var bs [1]byte
	bs[0] = b
	_, err := w.Write(bs[:])
	if err != nil {
		return err
	}
	return nil
}

func readByte(r io.Reader) (byte, error) {
	var bs [1]byte
	_, err := io.ReadFull(r, bs[:])
	if err != nil {
		return 0, err
	}
	b := bs[0]
	return b, nil
}

func writeFull(w io.Writer, data []byte) (n int, err error) {
	return w.Write(data)
}

func readFull(r io.Reader, data []byte) (n int, err error) {
	return io.ReadFull(r, data)
}

// quo = x / y
// rem = x % y
func quoRem(x, y int) (quo, rem int) {
	quo = x / y
	rem = x % y
	return
}

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

// func encodeTwoDigits(buf *bytes.Buffer, val int) error {
// 	const (
// 		n    = 2
// 		base = 10
// 	)
// 	var bs [2]byte
// 	var digit int
// 	for i := n; i > 0; i-- {
// 		val, digit = quoRem(val, base)
// 		b, ok := digitToByte(digit)
// 		if !ok {
// 			return fmt.Errorf("invalid convert digit %d to byte", digit)
// 		}
// 		bs[i-1] = b
// 	}
// 	_, err := buf.Write(bs[:])
// 	return err
// }

// func decodeTwoDigits(bs []byte) (int, error) {
// 	const (
// 		n    = 2
// 		base = 10
// 	)
// 	if len(bs) < 2 {
// 		return 0, fmt.Errorf("decodeTwoDigits: insufficient data length, have:%d, want:%d", len(bs), 2)
// 	}
// 	var value int
// 	for i := 0; i < n; i++ {
// 		b := bs[i]
// 		digit, ok := byteToDigit(b)
// 		if !ok {
// 			return 0, fmt.Errorf("decodeTwoDigits: invalid convert byte %x to digit", b)
// 		}
// 		value = value*base + digit
// 	}
// 	return value, nil
// }

func cloneBytes(a []byte) []byte {
	b := make([]byte, len(a))
	copy(b, a)
	return b
}
