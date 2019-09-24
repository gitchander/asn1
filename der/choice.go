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

func ChoiceSerializeDER(c Choice, params ...Parameter) (*Node, error) {

	tag, ok := getTagByParams(params)
	if !ok {
		tag, err := c.GetTag()
		if err != nil {
			return nil, err
		}
		return valueSerialize(reflect.ValueOf(c.Value()), Tag(tag))
	}

	n := NewNode(CLASS_CONTEXT_SPECIFIC, tag)

	childTag, err := c.GetTag()
	if err != nil {
		return nil, err
	}

	child, err := valueSerialize(reflect.ValueOf(c.Value()), Tag(childTag))
	if err != nil {
		return nil, err
	}

	nodes := []*Node{child}
	n.SetNodes(nodes)

	return n, nil
}

func ChoiceDeserializeDER(c Choice, n *Node, params ...Parameter) error {

	tag, ok := getTagByParams(params)
	if !ok {
		tag := n.GetTag()
		err := c.SetTag(tag)
		if err != nil {
			return err
		}
		return valueDeserialize(reflect.ValueOf(c.Value()), n, Tag(tag))
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

	return valueDeserialize(reflect.ValueOf(c.Value()), child, Tag(childTag))
}
