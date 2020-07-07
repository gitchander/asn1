package der

import (
	"encoding/binary"
	"math"
	"reflect"
)

var byteOrder = binary.BigEndian

func intBytesCrop(data []byte) []byte {

	if size := len(data); size > 0 {

		sign := data[0] & 0x80

		var b byte
		if sign != 0 {
			b = 0xFF
		}

		pos := 0
		for pos+1 < size {

			if data[pos] != b {
				break
			}

			if (data[pos+1] & 0x80) != sign {
				break
			}

			pos++
		}

		data = data[pos:]
	}

	return data
}

func intBytesComplete(data []byte, n int) []byte {

	if size := len(data); size < n {

		newData := make([]byte, n)

		var b byte
		if (data[0] & 0x80) != 0 {
			b = 0xFF
		}

		pos := 0
		for pos+size < n {
			newData[pos] = b
			pos++
		}

		copy(newData[pos:], data)
		data = newData
	}

	return data
}

func intEncode(x int64) []byte {
	data := make([]byte, sizeOfUint64)
	byteOrder.PutUint64(data, uint64(x))
	return intBytesCrop(data)
}

func uintEncode(x uint64) []byte {
	data := make([]byte, sizeOfUint64+1)
	data[0] = 0
	byteOrder.PutUint64(data[1:], x)
	return intBytesCrop(data)
}

func intDecode(data []byte) (int64, error) {
	completedData := intBytesComplete(data, sizeOfUint64)
	if len(completedData) == sizeOfUint64 {
		return int64(byteOrder.Uint64(completedData)), nil
	}
	return 0, ErrorUnmarshalBytes{data, reflect.Int}
}

func uintDecode(data []byte) (uint64, error) {
	completedData := intBytesComplete(data, sizeOfUint64+1)
	if len(completedData) == sizeOfUint64+1 {
		if completedData[0] == 0 {
			return byteOrder.Uint64(completedData[1:]), nil
		}
	}
	return 0, ErrorUnmarshalBytes{data, reflect.Uint}
}

func uintSerialize(v reflect.Value, params ...Parameter) (*Node, error) {
	return UintSerialize(v.Uint(), params...)
}

func intSerialize(v reflect.Value, params ...Parameter) (*Node, error) {
	return IntSerialize(v.Int(), params...)
}

func uintDeserialize(v reflect.Value, n *Node, params ...Parameter) error {

	x, err := UintDeserialize(n, params...)
	if err != nil {
		return err
	}

	switch k := v.Kind(); k {
	case reflect.Uint:
		if x > uint64(maxUint) {
			return ErrorUnmarshalUint{x, k}
		}
	case reflect.Uint8:
		if x > math.MaxUint8 {
			return ErrorUnmarshalUint{x, k}
		}
	case reflect.Uint16:
		if x > math.MaxUint16 {
			return ErrorUnmarshalUint{x, k}
		}
	case reflect.Uint32:
		if x > math.MaxUint32 {
			return ErrorUnmarshalUint{x, k}
		}
	}

	v.SetUint(x)

	return nil
}

func intDeserialize(v reflect.Value, n *Node, params ...Parameter) error {

	x, err := IntDeserialize(n, params...)
	if err != nil {
		return err
	}

	switch k := v.Kind(); k {
	case reflect.Int:
		if (int64(minInt) > x) || (x > int64(maxInt)) {
			return ErrorUnmarshalInt{x, k}
		}
	case reflect.Int8:
		if (math.MinInt8 > x) || (x > math.MaxInt8) {
			return ErrorUnmarshalInt{x, k}
		}
	case reflect.Int16:
		if (math.MinInt16 > x) || (x > math.MaxInt16) {
			return ErrorUnmarshalInt{x, k}
		}
	case reflect.Int32:
		if (math.MinInt32 > x) || (x > math.MaxInt32) {
			return ErrorUnmarshalInt{x, k}
		}
	}

	v.SetInt(x)

	return nil
}

func IntSerialize(x int64, params ...Parameter) (*Node, error) {

	class := CLASS_CONTEXT_SPECIFIC
	tag, ok := getTagByParams(params)
	if !ok {
		class = CLASS_UNIVERSAL
		tag = TAG_INTEGER
	}

	n := NewNode(class, tag)
	n.SetInt(x)

	return n, nil
}

func IntDeserialize(n *Node, params ...Parameter) (int64, error) {

	class := CLASS_CONTEXT_SPECIFIC
	tag, ok := getTagByParams(params)
	if !ok {
		class = CLASS_UNIVERSAL
		tag = TAG_INTEGER
	}

	err := CheckNode(n, class, tag)
	if err != nil {
		return 0, err
	}

	return n.GetInt()
}

func UintSerialize(x uint64, params ...Parameter) (*Node, error) {

	class := CLASS_CONTEXT_SPECIFIC
	tag, ok := getTagByParams(params)
	if !ok {
		class = CLASS_UNIVERSAL
		tag = TAG_INTEGER
	}

	n := NewNode(class, tag)
	n.SetUint(x)

	return n, nil
}

func UintDeserialize(n *Node, params ...Parameter) (uint64, error) {

	class := CLASS_CONTEXT_SPECIFIC
	tag, ok := getTagByParams(params)
	if !ok {
		class = CLASS_UNIVERSAL
		tag = TAG_INTEGER
	}

	err := CheckNode(n, class, tag)
	if err != nil {
		return 0, err
	}

	return n.GetUint()
}
