package der

// Enumerated

func EnumSerialize(e int, params ...Parameter) (*Node, error) {

	class := CLASS_CONTEXT_SPECIFIC
	tag, ok := getTagByParams(params)
	if !ok {
		class = CLASS_UNIVERSAL
		tag = TAG_ENUMERATED
	}

	n := NewNode(class, tag)
	n.SetInt(int64(e))

	return n, nil
}

func EnumDeserialize(n *Node, params ...Parameter) (int, error) {

	class := CLASS_CONTEXT_SPECIFIC
	tag, ok := getTagByParams(params)
	if !ok {
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
