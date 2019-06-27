package der

import (
	"errors"
	"reflect"

	"github.com/gitchander/asn1/der/coda"
)

func Serialize(v interface{}) (*Node, error) {
	return valueSerialize(reflect.ValueOf(v))
}

func valueSerialize(v reflect.Value) (*Node, error) {
	fn := getSerializeFunc(v.Type())
	return fn(v)
}

type serializeFunc func(v reflect.Value) (*Node, error)

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
	}

	return nil
}

func funcSerialize(v reflect.Value) (*Node, error) {

	if (v.Kind() == reflect.Ptr) && v.IsNil() {
		return nullSerialize(v)
	}

	s := v.Interface().(Serializer)
	return s.SerializeDER(-1)
}

func nullSerialize(v reflect.Value) (*Node, error) {

	h := coda.Header{
		Class:      CLASS_UNIVERSAL,
		Tag:        TAG_NULL,
		IsCompound: false,
	}

	n := new(Node)
	n.setHeader(h)

	return n, nil
}

func float32Serialize(v reflect.Value) (*Node, error) {

	return nil, nil
}

func float64Serialize(v reflect.Value) (*Node, error) {

	return nil, nil
}

func stringSerialize(v reflect.Value) (*Node, error) {

	h := coda.Header{
		Class:      CLASS_UNIVERSAL,
		Tag:        TAG_UTF8_STRING,
		IsCompound: false,
	}

	n := new(Node)
	n.setHeader(h)

	n.data = []byte(v.String())

	return n, nil
}

func bytesSerialize(v reflect.Value) (*Node, error) {

	h := coda.Header{
		Class:      CLASS_UNIVERSAL,
		Tag:        TAG_OCTET_STRING,
		IsCompound: false,
	}

	n := new(Node)
	n.setHeader(h)

	n.data = cloneBytes(v.Bytes())

	return n, nil
}

func structSerialize(v reflect.Value) (*Node, error) {

	tinfo, err := getTypeInfo(v.Type())
	if err != nil {
		return nil, err
	}

	n := NewConstructed(-1)
	for i := 0; i < v.NumField(); i++ {
		err := structFieldSerialize(n, v.Field(i), &(tinfo.fields[i]))
		if err != nil {
			return nil, err
		}
	}
	return n, nil
}

func structFieldSerialize(n *Node, v reflect.Value, finfo *fieldInfo) error {

	if (v.Kind() == reflect.Ptr) && (v.IsNil()) {
		if finfo.optional {
			return nil
		} else {
			return errors.New("Serializer is nil")
		}
	}

	if finfo.tag == nil {
		return nil
	}

	tag := *(finfo.tag)

	cs := NewConstructed(tag)

	fn := getSerializeFunc(v.Type())
	child, err := fn(v)
	if err != nil {
		return err
	}

	cs.nodes = []*Node{child}
	n.nodes = append(n.nodes, cs)

	return nil
}

type ptrSerializer struct {
	fn serializeFunc
}

func newPtrSerialize(t reflect.Type) serializeFunc {
	s := ptrSerializer{getSerializeFunc(t.Elem())}
	return s.encode
}

func (p *ptrSerializer) encode(v reflect.Value) (*Node, error) {

	if (v.Kind() == reflect.Ptr) && v.IsNil() {
		return nullSerialize(v)
	}

	return p.fn(v.Elem())
}

type arraySerializer struct {
	fn serializeFunc
}

func newArraySerialize(t reflect.Type) serializeFunc {
	s := arraySerializer{getSerializeFunc(t.Elem())}
	return s.encode
}

func (p *arraySerializer) encode(v reflect.Value) (*Node, error) {
	if (v.Kind() == reflect.Ptr) && v.IsNil() {
		return nullSerialize(v)
	}
	n := NewConstructed(-1)
	k := v.Len()
	for i := 0; i < k; i++ {

		child, err := p.fn(v.Index(i))
		if err != nil {
			return nil, err
		}

		n.nodes = append(n.nodes, child)
	}
	return n, nil
}

func newSliceSerialize(t reflect.Type) serializeFunc {

	if t.Elem().Kind() == reflect.Uint8 {
		return bytesSerialize
	}

	return newArraySerialize(t)
}
