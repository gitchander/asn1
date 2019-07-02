package der

import (
	"errors"
	"reflect"
)

type Choice interface {
	GetTag() (tag int, err error)
	SetTag(tag int) error

	Value() interface{}
}

func ChoiceSerializeDER(c Choice, tag int) (*Node, error) {
	if tag < 0 {
		tag, err := c.GetTag()
		if err != nil {
			return nil, err
		}
		return valueSerialize(reflect.ValueOf(c.Value()), tag)
	}

	n := NewNode(CLASS_CONTEXT_SPECIFIC, tag)

	childTag, err := c.GetTag()
	if err != nil {
		return nil, err
	}

	child, err := valueSerialize(reflect.ValueOf(c.Value()), childTag)
	if err != nil {
		return nil, err
	}

	nodes := []*Node{child}
	n.SetNodes(nodes)

	return n, nil
}

func ChoiceDeserializeDER(c Choice, n *Node, tag int) error {

	if tag < 0 {
		tag := n.GetTag()
		err := c.SetTag(tag)
		if err != nil {
			return err
		}
		return valueDeserialize(reflect.ValueOf(c.Value()), n, tag)
	}

	err := CheckNode(n, CLASS_CONTEXT_SPECIFIC, tag)
	if err != nil {
		return err
	}

	nodes, err := n.GetNodes()
	if err != nil {
		return err
	}
	if len(nodes) == 0 {
		return errors.New("nodes length = 0")
	}
	child := nodes[0]

	childTag := child.GetTag()
	if err = c.SetTag(childTag); err != nil {
		return err
	}

	return valueDeserialize(reflect.ValueOf(c.Value()), child, childTag)
}
