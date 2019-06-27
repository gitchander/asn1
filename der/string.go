package der

import (
	"errors"
	"reflect"
	"unicode/utf8"
)

func stringSerialize(v reflect.Value, tag int) (*Node, error) {
	return StringSerialize(v.String(), tag)
}

func stringDeserialize(v reflect.Value, n *Node, tag int) error {
	s, err := StringDeserialize(n, tag)
	if err != nil {
		return err
	}
	v.SetString(s)
	return nil
}

func StringSerialize(s string, tag int) (*Node, error) {

	class := CLASS_CONTEXT_SPECIFIC
	if tag < 0 {
		class = CLASS_UNIVERSAL
		tag = TAG_UTF8_STRING
	}

	n := NewNode(class, tag)
	n.SetString(s)

	return n, nil
}

func StringDeserialize(n *Node, tag int) (string, error) {

	class := CLASS_CONTEXT_SPECIFIC
	if tag < 0 {
		class = CLASS_UNIVERSAL
		tag = TAG_UTF8_STRING
	}

	err := CheckNode(n, class, tag)
	if err != nil {
		return "", err
	}

	s, err := n.GetString()
	if err != nil {
		return "", err
	}

	if !utf8.ValidString(s) {
		return "", errors.New("invalid UTF-8 string")
	}

	return s, nil
}
