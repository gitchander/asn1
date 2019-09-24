package der

import (
	"fmt"
	"reflect"
)

/*

type Tag int

type Serializer interface {
	SerializeDER(params ...interface{}) (*Node, error)
}

type Deserializer interface {
	DeserializeDER(n *Node, params ...interface{}) error
}

ussage:
	SerializeDER()
	SerializeDER(der.Tag(0))
	SerializeDER(der.Tag(1))

*/

type Serializer interface {
	SerializeDER(params ...Parameter) (*Node, error)
}

type Deserializer interface {
	DeserializeDER(n *Node, params ...Parameter) error
}

var (
	typeSerializer   = reflect.TypeOf((*Serializer)(nil)).Elem()
	typeDeserializer = reflect.TypeOf((*Deserializer)(nil)).Elem()
)

// asn1.MarshalWithParams(val interface{}, params string) ([]byte, error)

func Marshal(v interface{}, params ...Parameter) ([]byte, error) {
	n, err := Serialize(v, params...)
	if err != nil {
		return nil, err
	}
	return EncodeNode(nil, n)
}

// asn1.UnmarshalWithParams(b []byte, val interface{}, params string) (rest []byte, err error)

func Unmarshal(data []byte, v interface{}, params ...Parameter) error {
	n := new(Node)
	rest, err := DecodeNode(data, n)
	if err != nil {
		return err
	}
	if len(rest) > 0 {
		return fmt.Errorf("extra data length %d", len(rest))
	}
	return Deserialize(v, n, params...)
}
