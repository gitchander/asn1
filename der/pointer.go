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

func (p *ptrSerializer) encode(v reflect.Value, tag int) (*Node, error) {

	if (v.Kind() == reflect.Ptr) && v.IsNil() {
		return nullSerialize(v, tag)
	}

	return p.fn(v.Elem(), tag)
}

type ptrDeserializer struct {
	fn deserializeFunc
}

func newPtrDeserialize(t reflect.Type) deserializeFunc {
	d := ptrDeserializer{getDeserializeFunc(t.Elem())}
	return d.decode
}

func (p *ptrDeserializer) decode(v reflect.Value, n *Node, tag int) error {

	if v.IsNil() {
		return fmt.Errorf("der: Decode(nil %s)", v.Type())
	}

	return ptrValueDeserialize(v.Elem(), n, tag, p.fn)
}

func ptrValueDeserialize(v reflect.Value, n *Node, tag int, fn deserializeFunc) error {

	// err := nullDeserialize(v, n, tag)
	// if err == nil {
	// 	return nil
	// }

	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			valueMake(v)
		}
	}

	return fn(v, n, tag)
}
