package der

import (
	"errors"
	"reflect"
)

func nullSerialize(v reflect.Value, tag int) (*Node, error) {
	return _nullSerialize(tag)
}

func nullDeserialize(v reflect.Value, n *Node, tag int) error {
	err := _nullDeserialize(n, tag)
	if err != nil {
		return err
	}
	valueSetZero(v)
	return nil
}

func _nullSerialize(tag int) (*Node, error) {

	class := CLASS_CONTEXT_SPECIFIC
	if tag < 0 {
		class = CLASS_UNIVERSAL
		tag = TAG_NULL
	}

	n := NewNode(class, tag)

	return n, nil
}

func _nullDeserialize(n *Node, tag int) error {

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
