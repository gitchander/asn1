package der

import (
	"reflect"
)

func stringSerialize(v reflect.Value, params ...Parameter) (*Node, error) {
	return StringSerialize(v.String(), params...)
}

func stringDeserialize(v reflect.Value, n *Node, params ...Parameter) error {
	s, err := StringDeserialize(n, params...)
	if err != nil {
		return err
	}
	v.SetString(s)
	return nil
}

func StringSerialize(s string, params ...Parameter) (*Node, error) {

	class := CLASS_CONTEXT_SPECIFIC
	tag, ok := GetTagByParams(params)
	if !ok {
		class = CLASS_UNIVERSAL
		tag = TAG_UTF8_STRING
	}

	n := NewNode(class, tag)
	n.SetString(s)

	return n, nil
}

func StringDeserialize(n *Node, params ...Parameter) (string, error) {

	class := CLASS_CONTEXT_SPECIFIC
	tag, ok := GetTagByParams(params)
	if !ok {
		class = CLASS_UNIVERSAL
		tag = TAG_UTF8_STRING
	}

	err := CheckNode(n, class, tag)
	if err != nil {
		return "", err
	}

	return n.GetString()
}
