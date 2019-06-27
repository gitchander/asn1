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
	SerializeDER(tag int) (*Node, error)
}

type Deserializer interface {
	DeserializeDER(n *Node, tag int) error
}

var (
	typeSerializer   = reflect.TypeOf((*Serializer)(nil)).Elem()
	typeDeserializer = reflect.TypeOf((*Deserializer)(nil)).Elem()
)

func Marshal(v interface{}) ([]byte, error) {
	n, err := Serialize(v)
	if err != nil {
		return nil, err
	}
	return EncodeNode(nil, n)
}

func Unmarshal(data []byte, v interface{}) error {
	n := new(Node)
	rest, err := DecodeNode(data, n)
	if err != nil {
		return err
	}
	if len(rest) > 0 {
		return fmt.Errorf("extra data length %d", len(rest))
	}
	return Deserialize(v, n)
}
