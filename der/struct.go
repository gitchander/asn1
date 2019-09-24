package der

import (
	"fmt"
	"reflect"
)

func structSerialize(v reflect.Value, params ...Parameter) (*Node, error) {

	tinfo, err := getTypeInfo(v.Type())
	if err != nil {
		return nil, err
	}

	count := v.NumField()
	nodes := make([]*Node, 0, count)
	for i := 0; i < count; i++ {
		child, err := structFieldSerialize(v.Field(i), &(tinfo.fields[i]))
		if err != nil {
			return nil, fmt.Errorf("%s >> %s", v.Type(), err)
		}
		if child != nil {
			nodes = append(nodes, child)
		}
	}

	n := NewConstructed(params...)
	n.SetNodes(nodes)
	return n, nil
}

func structFieldSerialize(v reflect.Value, finfo *fieldInfo) (*Node, error) {

	if finfo.tag == nil {
		return nil, fmt.Errorf("field %s hasn't tag", fieldInfoToString(finfo))
	}
	tag := *(finfo.tag)

	if (v.Kind() == reflect.Ptr) && (v.IsNil()) {
		if finfo.optional {
			return nil, nil
		}
		return nil, fmt.Errorf("field %s value is nil, field must be optional", fieldInfoToString(finfo))
	}

	if finfo.explicit {

		fn := getSerializeFunc(v.Type())
		child, err := fn(v)
		if err != nil {
			return nil, err
		}
		nodes := []*Node{child}

		n := NewConstructed(Tag(tag))
		n.SetNodes(nodes)
		return n, nil
	}

	fn := getSerializeFunc(v.Type())
	return fn(v, Tag(tag))
}

func structDeserialize(v reflect.Value, n *Node, params ...Parameter) error {

	tinfo, err := getTypeInfo(v.Type())
	if err != nil {
		return err
	}

	err = CheckConstructed(n, params...)
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
			return fmt.Errorf("%s >> %s", v.Type(), err)
		}
	}

	return nil
}

func structFieldDeserialize(nodes []*Node, v reflect.Value, finfo *fieldInfo) error {

	if finfo.tag == nil {
		return fmt.Errorf("field %s hasn't tag", fieldInfoToString(finfo))
	}
	tag := *(finfo.tag)

	n := NodeByTag(nodes, tag)
	if n == nil {
		if finfo.optional {
			valueSetZero(v)
			return nil
		}
		return fmt.Errorf("field %s value is nil, field must be optional", fieldInfoToString(finfo))
	}

	if finfo.explicit {

		err := CheckConstructed(n, Tag(tag))
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
		return fn(v, child)
	}

	fn := getDeserializeFunc(v.Type())
	return fn(v, n, Tag(tag))
}
