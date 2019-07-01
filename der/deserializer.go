package der

import (
	"errors"
	"fmt"
	"reflect"
)

func Deserialize(v interface{}, n *Node) error {
	return valueDeserialize(reflect.ValueOf(v), n, -1)
}

func DeserializeTag(v interface{}, n *Node, tag int) error {
	return valueDeserialize(reflect.ValueOf(v), n, tag)
}

func valueDeserialize(v reflect.Value, n *Node, tag int) error {
	if v.Kind() != reflect.Ptr {
		return errors.New("value is not ptr")
	}
	fn := getDeserializeFunc(v.Type())
	return fn(v, n, tag)
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

func newSliceDeserialize(t reflect.Type) deserializeFunc {

	if t.Elem().Kind() == reflect.Uint8 {
		return bytesDeserialize
	}

	return newArrayDeserialize(t)
}
