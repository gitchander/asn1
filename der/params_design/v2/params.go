package v2

import (
	"fmt"
	"log"
)

type Node struct{}

type Serializer interface {
	SerializeDER(parameters ...Parameter) (*Node, error)
}

type Deserializer interface {
	DeserializeDER(n *Node, parameters ...Parameter) error
}

type Parameter interface {
	isParameter()
}

type Name string
type Tag int
type Optional struct{}
type Explicit struct{}

func (Name) isParameter()     {}
func (Tag) isParameter()      {}
func (Optional) isParameter() {}
func (Explicit) isParameter() {}

func TagByParameters(parameters ...Parameter) (tag Tag, ok bool) {
	for _, p := range parameters {
		if tag, ok := p.(Tag); ok {
			return tag, true
		}
	}
	return 0, false
}

//var DefaultOptions = Options{
//	//Tag: -1,
//	//Tag: optInt{},
//}

//func Name(name string) Option {
//	return func(o *Options) error {
//		o.Name = name
//		return nil
//	}
//}

//func Tag(tag int) Option {
//	return func(o *Options) error {
//		o.Tag = makeOptInt(tag)
//		return nil
//	}
//}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func TestParams(parameters ...Parameter) {
	for _, p := range parameters {
		fmt.Printf("%T, %v\n", p, p)
	}
}
