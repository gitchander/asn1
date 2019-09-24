package der

import (
	"reflect"
)

func bytesSerialize(v reflect.Value, params ...Parameter) (*Node, error) {
	return BytesSerialize(v.Bytes(), params...)
}

func bytesDeserialize(v reflect.Value, n *Node, params ...Parameter) error {
	bs, err := BytesDeserialize(n, params...)
	if err != nil {
		return err
	}
	v.SetBytes(bs)
	return nil
}

func BytesSerialize(bs []byte, params ...Parameter) (*Node, error) {

	class := CLASS_CONTEXT_SPECIFIC
	tag, ok := GetTagByParams(params)
	if !ok {
		class = CLASS_UNIVERSAL
		tag = TAG_OCTET_STRING
	}

	n := NewNode(class, tag)
	n.SetBytes(bs)

	return n, nil
}

func BytesDeserialize(n *Node, params ...Parameter) ([]byte, error) {

	class := CLASS_CONTEXT_SPECIFIC
	tag, ok := GetTagByParams(params)
	if !ok {
		class = CLASS_UNIVERSAL
		tag = TAG_OCTET_STRING
	}

	err := CheckNode(n, class, tag)
	if err != nil {
		return nil, err
	}

	return n.GetBytes()
}
