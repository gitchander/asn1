package der

import (
	"fmt"
	"reflect"
)

type ptrSerializer struct {
	fn serializeFunc
}

func newPtrSerialize(t reflect.Type) serializeFunc {
	s := ptrSerializer{getSerializeFunc(t.Elem())}
	return s.encode
}

func (p *ptrSerializer) encode(v reflect.Value, params ...Parameter) (*Node, error) {

	if (v.Kind() == reflect.Ptr) && v.IsNil() {
		return nullSerialize(v, params...)
	}

	return p.fn(v.Elem(), params...)
}

type ptrDeserializer struct {
	fn deserializeFunc
}

func newPtrDeserialize(t reflect.Type) deserializeFunc {
	d := ptrDeserializer{getDeserializeFunc(t.Elem())}
	return d.decode
}

func (p *ptrDeserializer) decode(v reflect.Value, n *Node, params ...Parameter) error {

	if v.IsNil() {
		return fmt.Errorf("der: Decode(nil %s)", v.Type())
	}

	return ptrValueDeserialize(p.fn, v.Elem(), n, params...)
}

func ptrValueDeserialize(fn deserializeFunc, v reflect.Value, n *Node, params ...Parameter) error {

	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			makeValue(v)
		}
	}

	return fn(v, n, params...)
}
