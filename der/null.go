package der

import (
	"errors"
	"reflect"
)

func nullSerialize(v reflect.Value, tag int) (*Node, error) {
	return NullSerialize(tag)
}

func nullDeserialize(v reflect.Value, n *Node, tag int) error {
	err := NullDeserialize(n, tag)
	if err != nil {
		return err
	}
	valueSetZero(v)
	return nil
}

func NullSerialize(tag int) (*Node, error) {

	class := CLASS_CONTEXT_SPECIFIC
	if tag < 0 {
		class = CLASS_UNIVERSAL
		tag = TAG_NULL
	}

	n := NewNode(class, tag)

	return n, nil
}

func NullDeserialize(n *Node, tag int) error {

	class := CLASS_CONTEXT_SPECIFIC
	if tag < 0 {
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
