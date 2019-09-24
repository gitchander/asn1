package der

import (
	"errors"
	"reflect"
)

func nullSerialize(v reflect.Value, params ...Parameter) (*Node, error) {
	return NullSerialize(params...)
}

func nullDeserialize(v reflect.Value, n *Node, params ...Parameter) error {
	err := NullDeserialize(n, params...)
	if err != nil {
		return err
	}
	valueSetZero(v)
	return nil
}

func NullSerialize(params ...Parameter) (*Node, error) {

	class := CLASS_CONTEXT_SPECIFIC
	tag, ok := getTagByParams(params)
	if !ok {
		class = CLASS_UNIVERSAL
		tag = TAG_NULL
	}

	n := NewNode(class, tag)

	return n, nil
}

func NullDeserialize(n *Node, params ...Parameter) error {

	class := CLASS_CONTEXT_SPECIFIC
	tag, ok := getTagByParams(params)
	if !ok {
		class = CLASS_UNIVERSAL
		tag = TAG_NULL
	}

	err := CheckNode(n, class, tag)
	if err != nil {
		return err
	}

	if n.IsConstructed() {
		return errors.New("NullDeserialize: node is constructed")
	}

	return nil
}
