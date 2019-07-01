package der

import (
	"errors"
	"fmt"
	"reflect"
)

func structSerialize(v reflect.Value, tag int) (*Node, error) {

	tinfo, err := getTypeInfo(v.Type())
	if err != nil {
		return nil, err
	}

	count := v.NumField()
	nodes := make([]*Node, 0, count)
	for i := 0; i < count; i++ {
		child, err := structFieldSerialize(v.Field(i), &(tinfo.fields[i]))
		if err != nil {
			return nil, err
		}
		if child != nil {
			nodes = append(nodes, child)
		}
	}

	n := NewConstructed(tag)
	n.SetNodes(nodes)
	return n, nil
}

func structFieldSerialize(v reflect.Value, finfo *fieldInfo) (*Node, error) {

	if (v.Kind() == reflect.Ptr) && (v.IsNil()) {
		if finfo.optional {
			return nil, nil
		} else {
			return nil, errors.New("structFieldSerialize: serialize nil value")
		}
	}

	if finfo.tag == nil {
		return nil, errors.New("struct field tag is not exist")
	}
	tag := *(finfo.tag)

	if finfo.explicit {

		fn := getSerializeFunc(v.Type())
		child, err := fn(v, -1)
		if err != nil {
			return nil, err
		}
		nodes := []*Node{child}

		n := NewConstructed(tag)
		n.SetNodes(nodes)
		return n, nil
	}

	fn := getSerializeFunc(v.Type())
	return fn(v, tag)
}

func structDeserialize(v reflect.Value, n *Node, tag int) error {

	tinfo, err := getTypeInfo(v.Type())
	if err != nil {
		return err
	}

	err = CheckConstructed(n, tag)
	if err != nil {
		return err
	}

	nodes, err := n.GetNodes()
	if err != nil {
		return err
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.CanAddr() {
			field = field.Addr()
		}
		err := structFieldDeserialize(nodes, field, &(tinfo.fields[i]))
		if err != nil {
			return fmt.Errorf("der: field no. %d (for %v) deserialize error: %s", i, field.Type(), err)
		}
	}

	return nil
}

func structFieldDeserialize(nodes []*Node, v reflect.Value, finfo *fieldInfo) error {

	if finfo.tag == nil {
		return errors.New("struct field tag is not exist")
	}
	tag := *(finfo.tag)

	n := NodeByTag(nodes, tag)
	if n == nil {
		if finfo.optional {
			valueSetZero(v)
			return nil
		}
		return errors.New("structFieldDeserialize: deserialize nil value")
	}

	if finfo.explicit {

		err := CheckConstructed(n, tag)
		if err != nil {
			return err
		}
		cs, err := n.GetNodes()
		if err != nil {
			return err
		}
		if len(cs) != 1 {
			return fmt.Errorf("CS child number %d", len(cs))
		}
		child := cs[0]

		fn := getDeserializeFunc(v.Type())
		return fn(v, child, -1)
	}

	fn := getDeserializeFunc(v.Type())
	return fn(v, n, tag)
}
