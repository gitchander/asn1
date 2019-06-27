package der

import (
	"errors"
	"fmt"
	"reflect"
)

func Serialize(v interface{}) (*Node, error) {
	return valueSerialize(reflect.ValueOf(v))
}

func valueSerialize(v reflect.Value) (*Node, error) {
	fn := getSerializeFunc(v.Type())
	return fn(v, -1)
}

type serializeFunc func(v reflect.Value, tag int) (*Node, error)

func getSerializeFunc(t reflect.Type) serializeFunc {

	if t.Implements(typeSerializer) {
		return funcSerialize
	}

	switch k := t.Kind(); k {

	case reflect.Bool:
		return boolSerialize

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return intSerialize

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return uintSerialize

	case reflect.Float32:
		return float32Serialize

	case reflect.Float64:
		return float64Serialize

	case reflect.String:
		return stringSerialize

	case reflect.Struct:
		return structSerialize

	case reflect.Ptr:
		return newPtrSerialize(t)

	case reflect.Array:
		return newArraySerialize(t)

	case reflect.Slice:
		return newSliceSerialize(t)

	default:
		panic(fmt.Errorf("getSerializeFunc: unsupported type %s", k))
	}

	return nil
}

func funcSerialize(v reflect.Value, tag int) (*Node, error) {

	if (v.Kind() == reflect.Ptr) && v.IsNil() {
		return nullSerialize(v, tag)
	}

	s := v.Interface().(Serializer)
	return s.SerializeDER(tag)
}

func float32Serialize(v reflect.Value, tag int) (*Node, error) {

	panic("float32Serialize is not completed")

	return nil, nil
}

func float64Serialize(v reflect.Value, tag int) (*Node, error) {

	panic("float64Serialize is not completed")

	return nil, nil
}

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

type ptrSerializer struct {
	fn serializeFunc
}

func newPtrSerialize(t reflect.Type) serializeFunc {
	s := ptrSerializer{getSerializeFunc(t.Elem())}
	return s.encode
}

func (p *ptrSerializer) encode(v reflect.Value, tag int) (*Node, error) {

	if (v.Kind() == reflect.Ptr) && v.IsNil() {
		return nullSerialize(v, tag)
	}

	return p.fn(v.Elem(), tag)
}

type arraySerializer struct {
	fn serializeFunc
}

func newArraySerialize(t reflect.Type) serializeFunc {
	s := arraySerializer{getSerializeFunc(t.Elem())}
	return s.encode
}

func (p *arraySerializer) encode(v reflect.Value, tag int) (*Node, error) {
	if (v.Kind() == reflect.Ptr) && v.IsNil() {
		return nullSerialize(v, tag)
	}
	nodes := make([]*Node, v.Len())
	for i := range nodes {
		child, err := p.fn(v.Index(i), -1)
		if err != nil {
			return nil, err
		}
		nodes[i] = child
	}
	n := NewConstructed(tag)
	n.SetNodes(nodes)
	return n, nil
}

func newSliceSerialize(t reflect.Type) serializeFunc {

	if t.Elem().Kind() == reflect.Uint8 {
		return bytesSerialize
	}

	return newArraySerialize(t)
}
