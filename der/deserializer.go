package der

import (
	"errors"
	"fmt"
	"reflect"
)

func Deserialize(v interface{}, n *Node) error {
	return valueDeserialize(reflect.ValueOf(v), n)
}

func valueDeserialize(v reflect.Value, n *Node) error {
	if v.Kind() != reflect.Ptr {
		return errors.New("value is not ptr")
	}
	fn := getDeserializeFunc(v.Type())
	return fn(v, n, -1)
}

type deserializeFunc func(v reflect.Value, n *Node, tag int) error

func getDeserializeFunc(t reflect.Type) deserializeFunc {

	if t.Implements(typeDeserializer) {
		return funcDeserialize
	}

	switch k := t.Kind(); k {

	case reflect.Bool:
		return boolDeserialize

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return intDeserialize

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return uintDeserialize

	case reflect.Float32:
		return float32Deserialize

	case reflect.Float64:
		return float64Deserialize

	case reflect.String:
		return stringDeserialize

	case reflect.Struct:
		return structDeserialize

	case reflect.Ptr:
		return newPtrDeserialize(t)

	case reflect.Array:
		return newArrayDeserialize(t)

	case reflect.Slice:
		return newSliceDeserialize(t)

	default:
		panic(fmt.Errorf("getDeserializeFunc: unsupported type %s", k))
	}

	return nil
}

func funcDeserialize(v reflect.Value, n *Node, tag int) error {
	d := v.Interface().(Deserializer)
	return d.DeserializeDER(n, tag)
}

func float32Deserialize(v reflect.Value, n *Node, tag int) error {

	panic("float32Deserialize is not completed")

	return nil
}

func float64Deserialize(v reflect.Value, n *Node, tag int) error {

	panic("float64Deserialize is not completed")

	return nil
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

		//valueMake(v)

		fn := getDeserializeFunc(v.Type())
		return fn(v, child, -1)
	}

	//valueMake(v)

	fn := getDeserializeFunc(v.Type())
	return fn(v, n, tag)
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

type arrayDeserializer struct {
	fn deserializeFunc
}

func newArrayDeserialize(t reflect.Type) deserializeFunc {
	d := arrayDeserializer{getDeserializeFunc(t.Elem())}
	return d.decode
}

func (p *arrayDeserializer) decode(v reflect.Value, n *Node, tag int) error {

	return nil
}

func newSliceDeserialize(t reflect.Type) deserializeFunc {

	if t.Elem().Kind() == reflect.Uint8 {
		return bytesDeserialize
	}

	return newArrayDeserialize(t)
}
