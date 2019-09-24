package der

import (
	"reflect"
)

func boolEncode(x bool) []byte {
	data := []byte{0}
	if x {
		data[0] = 0xFF
	}
	return data
}

func boolDecode(data []byte) (bool, error) {
	if len(data) != 1 {
		return false, ErrorUnmarshalBytes{data, reflect.Bool}
	}
	return (data[0] != 0), nil
}

func boolSerialize(v reflect.Value, params ...Parameter) (*Node, error) {
	return BoolSerialize(v.Bool(), params...)
}

func boolDeserialize(v reflect.Value, n *Node, params ...Parameter) error {
	b, err := BoolDeserialize(n, params...)
	if err != nil {
		return err
	}
	v.SetBool(b)
	return nil
}

func BoolSerialize(b bool, params ...Parameter) (*Node, error) {

	class := CLASS_CONTEXT_SPECIFIC
	tag, ok := getTagByParams(params)
	if !ok {
		class = CLASS_UNIVERSAL
		tag = TAG_BOOLEAN
	}

	n := NewNode(class, tag)
	n.SetBool(b)

	return n, nil
}

func BoolDeserialize(n *Node, params ...Parameter) (bool, error) {

	class := CLASS_CONTEXT_SPECIFIC
	tag, ok := getTagByParams(params)
	if !ok {
		class = CLASS_UNIVERSAL
		tag = TAG_BOOLEAN
	}

	err := CheckNode(n, class, tag)
	if err != nil {
		return false, err
	}

	return n.GetBool()
}
