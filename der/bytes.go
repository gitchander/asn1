package der

func BytesSerialize(bs []byte, tag int) (*Node, error) {

	class := CLASS_CONTEXT_SPECIFIC
	if tag < 0 {
		class = CLASS_UNIVERSAL
		tag = TAG_OCTET_STRING
	}

	n := NewNode(class, tag)
	n.SetBytes(bs)

	return n, nil
}

func BytesDeserialize(n *Node, tag int) ([]byte, error) {

	class := CLASS_CONTEXT_SPECIFIC
	if tag < 0 {
		class = CLASS_UNIVERSAL
		tag = TAG_OCTET_STRING
	}

	err := CheckNode(n, class, tag)
	if err != nil {
		return nil, err
	}

	return n.GetBytes()
}
