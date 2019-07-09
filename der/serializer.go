package der

import (
	"fmt"
	"reflect"
)

func Serialize(v interface{}, tag int) (*Node, error) {
	return valueSerialize(reflect.ValueOf(v), tag)
}

func valueSerialize(v reflect.Value, tag int) (*Node, error) {
	fn := getSerializeFunc(v.Type())
	return fn(v, tag)
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
		return serializeUnsupportedType(t)
	}
}

func serializeUnsupportedType(t reflect.Type) serializeFunc {

	err := fmt.Errorf("der: serialize unsupported type %s", t.Kind())

	return func(v reflect.Value, tag int) (*Node, error) {
		return nil, err
	}
}

func funcSerialize(v reflect.Value, tag int) (*Node, error) {

	if (v.Kind() == reflect.Ptr) && v.IsNil() {
		return nullSerialize(v, tag)
	}

	s := v.Interface().(Serializer)
	return s.SerializeDER(tag)
}

func newSliceSerialize(t reflect.Type) serializeFunc {

	if t.Elem().Kind() == reflect.Uint8 {
		return bytesSerialize
	}

	return newArraySerialize(t)
}
