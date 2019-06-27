package der

// Enumerated
func EnumSerialize(e int, tag int) (*Node, error) {

	class := CLASS_CONTEXT_SPECIFIC
	if tag < 0 {
		class = CLASS_UNIVERSAL
		tag = TAG_ENUMERATED
	}

	n := NewNode(class, tag)
	n.SetInt(int64(e))

	return n, nil
}

func EnumDeserialize(n *Node, tag int) (int, error) {

	class := CLASS_CONTEXT_SPECIFIC
	if tag < 0 {
		class = CLASS_UNIVERSAL
		tag = TAG_ENUMERATED
	}

	err := CheckNode(n, class, tag)
	if err != nil {
		return 0, err
	}

	i, err := n.GetInt()
	if err != nil {
		return 0, err
	}

	return int(i), nil
}
