package der

import (
	"fmt"
	"strconv"
)

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

func charToDigit(char byte) (digit int, ok bool) {
	if ('0' <= char) && (char <= '9') {
		return int(char - '0'), true
	}
	return 0, false
}

func digitToChar(digit int) (char byte, ok bool) {
	if (0 <= digit) && (digit <= 9) {
		return byte(digit + '0'), true
	}
	return 0, false
}

func cloneBytes(a []byte) []byte {
	b := make([]byte, len(a))
	copy(b, a)
	return b
}

func parseInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func newInt(a int) *int {
	return &a
}

func not(a bool) bool {
	return !a
}

func checkHaveWantInt(name string, have, want int) error {
	if have != want {
		return fmt.Errorf("invalid (%s): have %d, want %d", name, have, want)
	}
	return nil
}
