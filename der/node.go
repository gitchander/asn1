package der

import (
	"errors"
	"fmt"
	"time"
	"unicode/utf8"

	"github.com/gitchander/asn1/der/coda"
)

var (
	ErrNodeIsConstructed    = errors.New("node is constructed")
	ErrNodeIsNotConstructed = errors.New("node is not constructed")
)

//------------------------------------------------------------------------------
// golang asn1:
//
// type RawValue struct {
// 	Class, Tag int
// 	IsCompound bool
// 	Bytes      []byte
// 	FullBytes  []byte // includes the tag and length
// }
//------------------------------------------------------------------------------

type Node struct {
	class       int
	tag         int
	constructed bool // isCompound

	data  []byte  // Primitive:   (isCompound = false)
	nodes []*Node // Constructed: (isCompound = true)
}

func NewNode(class int, tag int) *Node {
	return &Node{
		class: class,
		tag:   tag,
	}
}

func CheckNode(n *Node, class int, tag int) error {
	if n.class != class {
		return fmt.Errorf("class: %d != %d", n.class, class)
	}
	if n.tag != tag {
		return fmt.Errorf("tag: %d != %d", n.tag, tag)
	}
	return nil
}

func (n *Node) GetTag() int {
	return n.tag
}

func (n *Node) getHeader() coda.Header {
	return coda.Header{
		Class:      n.class,
		Tag:        n.tag,
		IsCompound: n.constructed,
	}
}

func (n *Node) IsPrimitive() bool {
	return !(n.constructed)
}

func (n *Node) IsConstructed() bool {
	return (n.constructed)
}

func (n *Node) setHeader(h coda.Header) error {
	*n = Node{
		class:       h.Class,
		tag:         h.Tag,
		constructed: h.IsCompound,
	}
	return nil
}

func (n *Node) checkHeader(h coda.Header) error {
	k := n.getHeader()
	if !coda.EqualHeaders(k, h) {
		return errors.New("der: invalid header")
	}
	return nil
}

func EncodeNode(data []byte, n *Node) (rest []byte, err error) {

	header := n.getHeader()
	data, err = coda.EncodeHeader(data, &header)
	if err != nil {
		return nil, err
	}

	value, err := encodeValue(n)
	if err != nil {
		return nil, err
	}

	length := len(value)
	data, err = coda.EncodeLength(data, length)
	if err != nil {
		return nil, err
	}

	data = append(data, value...)
	return data, err
}

func DecodeNode(data []byte, n *Node) (rest []byte, err error) {

	var header coda.Header
	data, err = coda.DecodeHeader(data, &header)
	if err != nil {
		return nil, err
	}
	err = n.setHeader(header)
	if err != nil {
		return nil, err
	}

	var length int
	data, err = coda.DecodeLength(data, &length)
	if err != nil {
		return nil, err
	}
	if len(data) < length {
		return nil, errors.New("insufficient data length")
	}

	err = decodeValue(data[:length], n)
	if err != nil {
		return nil, err
	}

	rest = data[length:]

	return rest, nil
}

func encodeValue(n *Node) ([]byte, error) {
	if !n.constructed {
		return cloneBytes(n.data), nil
	}
	return encodeNodes(n.nodes)
}

func decodeValue(data []byte, n *Node) error {

	if !n.constructed {
		n.data = cloneBytes(data)
		return nil
	}

	ns, err := decodeNodes(data)
	if err != nil {
		return err
	}
	n.nodes = ns

	return nil
}

//----------------------------------------------------------------------------

func (n *Node) SetNodes(ns []*Node) {
	n.constructed = true
	n.nodes = ns
}

func (n *Node) GetNodes() ([]*Node, error) {
	if !n.constructed {
		return nil, ErrNodeIsNotConstructed
	}
	return n.nodes, nil
}

func (n *Node) SetBool(b bool) {
	n.constructed = false
	n.data = boolEncode(b)
}

func (n *Node) GetBool() (bool, error) {
	if n.constructed {
		return false, ErrNodeIsConstructed
	}
	return boolDecode(n.data)
}

func (n *Node) SetInt(i int64) {
	n.constructed = false
	n.data = intEncode(i)
}

func (n *Node) GetInt() (int64, error) {
	if n.constructed {
		return 0, ErrNodeIsConstructed
	}
	return intDecode(n.data)
}

func (n *Node) SetUint(u uint64) {
	n.constructed = false
	n.data = uintEncode(u)
}

func (n *Node) GetUint() (uint64, error) {
	if n.constructed {
		return 0, ErrNodeIsConstructed
	}
	return uintDecode(n.data)
}

func (n *Node) SetBytes(bs []byte) {
	n.constructed = false
	n.data = bs
}

func (n *Node) GetBytes() ([]byte, error) {
	if n.constructed {
		return nil, ErrNodeIsConstructed
	}
	return n.data, nil
}

func (n *Node) SetString(s string) {
	n.constructed = false
	n.data = []byte(s)
}

func (n *Node) GetString() (string, error) {
	if n.constructed {
		return "", ErrNodeIsConstructed
	}
	if !utf8.Valid(n.data) {
		return "", errors.New("invalid utf8 string")
		//return "", errors.New("data is not utf-8 string")
	}
	return string(n.data), nil
}

func (n *Node) SetUTCTime(t time.Time) error {
	data, err := encodeUTCTime(t)
	if err != nil {
		return err
	}
	n.constructed = false
	n.data = data
	return nil
}

func (n *Node) GetUTCTime() (time.Time, error) {
	if n.constructed {
		return time.Time{}, ErrNodeIsConstructed
	}
	return decodeUTCTime(n.data)
}
