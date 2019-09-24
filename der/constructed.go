package der

import (
	"io"
)

func NodeByTag(ns []*Node, tag int) *Node {
	for _, n := range ns {
		if n.tag == tag {
			return n
		}
	}
	return nil
}

func NewConstructed(params ...Parameter) (n *Node) {

	class := CLASS_CONTEXT_SPECIFIC
	tag, ok := getTagByParams(params)
	if !ok {
		class = CLASS_UNIVERSAL
		tag = TAG_SEQUENCE
	}

	return &Node{
		class:       class,
		tag:         tag,
		constructed: true,
	}
}

func CheckConstructed(n *Node, params ...Parameter) error {

	if !n.constructed {
		return ErrNodeIsNotConstructed
	}

	tag, ok := getTagByParams(params)
	if !ok {
		return CheckNode(n, CLASS_UNIVERSAL, TAG_SEQUENCE)
	}

	return CheckNode(n, CLASS_CONTEXT_SPECIFIC, tag)
}

func childSerialize(n *Node, s Serializer, tag int) error {
	child, err := s.SerializeDER(Tag(tag))
	if err != nil {
		return err
	}
	if child != nil {
		n.nodes = append(n.nodes, child)
	}
	return nil
}

func childDeserialize(n *Node, d Deserializer, tag int) error {
	child := NodeByTag(n.nodes, tag)
	// child can be nil for an optional value
	return d.DeserializeDER(child, Tag(tag))
}

func encodeNodes(ns []*Node) (data []byte, err error) {
	for _, n := range ns {
		if n == nil {
			continue
		}
		data, err = EncodeNode(data, n)
		if err != nil {
			return nil, err
		}
	}
	return data, nil
}

func decodeNodes(data []byte) (ns []*Node, err error) {
	for {
		child := new(Node)
		data, err = DecodeNode(data, child)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		ns = append(ns, child)
	}
	return ns, nil
}
