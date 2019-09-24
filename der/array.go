package der

import (
	"reflect"
)

type arraySerializer struct {
	fn serializeFunc
}

func newArraySerialize(t reflect.Type) serializeFunc {
	s := arraySerializer{getSerializeFunc(t.Elem())}
	return s.encode
}

func (p *arraySerializer) encode(v reflect.Value, params ...Parameter) (*Node, error) {
	nodes := make([]*Node, v.Len())
	for i := range nodes {
		child, err := p.fn(v.Index(i))
		if err != nil {
			return nil, err
		}
		nodes[i] = child
	}
	n := NewConstructed(params...)
	n.SetNodes(nodes)
	return n, nil
}

type arrayDeserializer struct {
	fn deserializeFunc
}

func newArrayDeserialize(t reflect.Type) deserializeFunc {
	d := arrayDeserializer{getDeserializeFunc(t.Elem())}
	return d.decode
}

func (p *arrayDeserializer) decode(v reflect.Value, n *Node, params ...Parameter) error {

	err := CheckConstructed(n, params...)
	if err != nil {
		return err
	}
	nodes, err := n.GetNodes()
	if err != nil {
		return err
	}

	count := len(nodes)
	slice := reflect.MakeSlice(v.Type(), count, count)
	v.Set(slice)

	for i, child := range nodes {
		err = p.fn(v.Index(i), child)
		if err != nil {
			return err
		}
	}

	return nil
}
