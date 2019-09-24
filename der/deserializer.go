package der

import (
	"errors"
	"fmt"
	"reflect"
)

func Deserialize(v interface{}, n *Node, params ...Parameter) error {
	return valueDeserialize(reflect.ValueOf(v), n, params...)
}

func valueDeserialize(v reflect.Value, n *Node, params ...Parameter) error {
	if v.Kind() != reflect.Ptr {
		return errors.New("der deserialize: value is not ptr")
	}
	fn := getDeserializeFunc(v.Type())
	return fn(v, n, params...)
}

type deserializeFunc func(v reflect.Value, n *Node, params ...Parameter) error

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
		return deserializeUnsupportedType(t)
	}
}

func deserializeUnsupportedType(t reflect.Type) deserializeFunc {

	err := fmt.Errorf("der: deserialize unsupported type %s", t.Kind())

	return func(v reflect.Value, n *Node, params ...Parameter) error {
		return err
	}
}

func funcDeserialize(v reflect.Value, n *Node, params ...Parameter) error {
	d := v.Interface().(Deserializer)
	return d.DeserializeDER(n, params...)
}

func newSliceDeserialize(t reflect.Type) deserializeFunc {

	if t.Elem().Kind() == reflect.Uint8 {
		return bytesDeserialize
	}

	return newArrayDeserialize(t)
}
