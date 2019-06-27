package der

import (
	"errors"
	"fmt"
	"reflect"
	"unicode/utf8"

	"github.com/gitchander/asn1/der/coda"
)

func Deserialize(v interface{}, n *Node) error {
	return valueDeserialize(reflect.ValueOf(v), n)
}

func valueDeserialize(v reflect.Value, n *Node) error {

	if v.Kind() != reflect.Ptr {
		return errors.New("value is not ptr")
	}

	fn := getDeserializeFunc(v.Type())
	return fn(v, n)
}

type deserializeFunc func(reflect.Value, *Node) error

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
	}

	return nil
}

func funcDeserialize(v reflect.Value, n *Node) error {
	d := v.Interface().(Deserializer)
	return d.DeserializeDER(n, -1)
}

func float32Deserialize(v reflect.Value, n *Node) error {

	return nil
}

func float64Deserialize(v reflect.Value, n *Node) error {

	return nil
}

func stringDeserialize(v reflect.Value, n *Node) error {

	h := coda.Header{
		Class:      CLASS_UNIVERSAL,
		Tag:        TAG_UTF8_STRING,
		IsCompound: false,
	}

	err := n.checkHeader(h)
	if err != nil {
		return err
	}

	data := n.data
	if !utf8.Valid(data) {
		return errors.New("asn1: invalid UTF-8 string")
	}
	v.SetString(string(data))

	return nil
}

func bytesDeserialize(v reflect.Value, n *Node) error {

	h := coda.Header{
		Class:      CLASS_UNIVERSAL,
		Tag:        TAG_OCTET_STRING,
		IsCompound: false,
	}

	err := n.checkHeader(h)
	if err != nil {
		return err
	}
	v.SetBytes(cloneBytes(n.data))
	return nil
}

func structDeserialize(v reflect.Value, n *Node) error {

	tinfo, err := getTypeInfo(v.Type())
	if err != nil {
		return err
	}

	err = CheckConstructed(n, -1)
	if err != nil {
		return err
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		err := structFieldDeserialize(n.nodes, field, &(tinfo.fields[i]))
		if err != nil {
			return fmt.Errorf("der: field no.%d (for type %v) deserialize error: %s", i, field.Type(), err)
		}
	}

	return nil
}

func structFieldDeserialize(nodes []*Node, v reflect.Value, finfo *fieldInfo) error {

	if finfo.tag == nil {
		return errors.New("tag is nil")
	}

	tag := *(finfo.tag)

	cs := NodeByTag(nodes, tag)
	if cs == nil {
		if finfo.optional {
			valueSetZero(v)
			return nil
		}
		return errors.New("Deserializer is nil")
	}

	err := CheckConstructed(cs, tag)
	if err != nil {
		return err
	}

	child := cs.nodes[0]

	valueMake(v)

	fn := getDeserializeFunc(v.Type())
	return fn(v, child)
}

type ptrDeserializer struct {
	fn deserializeFunc
}

func newPtrDeserialize(t reflect.Type) deserializeFunc {
	d := ptrDeserializer{getDeserializeFunc(t.Elem())}
	return d.decode
}

func (p *ptrDeserializer) decode(v reflect.Value, n *Node) error {

	if v.IsNil() {
		return fmt.Errorf("der: Decode(nil %s)", v.Type())
	}

	return ptrValueDeserialize(v.Elem(), n, p.fn)
}

func ptrValueDeserialize(v reflect.Value, n *Node, fn deserializeFunc) error {

	h := coda.Header{
		Class:      CLASS_UNIVERSAL,
		Tag:        TAG_NULL,
		IsCompound: false,
	}

	err := n.checkHeader(h)
	if err == nil {
		valueSetZero(v)
		return nil
	}

	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			valueMake(v)
		}
	}

	return fn(v, n)
}

type arrayDeserializer struct {
	fn deserializeFunc
}

func newArrayDeserialize(t reflect.Type) deserializeFunc {
	d := arrayDeserializer{getDeserializeFunc(t.Elem())}
	return d.decode
}

func (p *arrayDeserializer) decode(v reflect.Value, n *Node) error {

	return nil
}

func newSliceDeserialize(t reflect.Type) deserializeFunc {

	if t.Elem().Kind() == reflect.Uint8 {
		return bytesDeserialize
	}

	return newArrayDeserialize(t)
}
