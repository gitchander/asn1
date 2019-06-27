package der

func StringSerialize(s string, tag int) (*Node, error) {

	class := CLASS_CONTEXT_SPECIFIC
	if tag < 0 {
		class = CLASS_UNIVERSAL
		tag = TAG_UTF8_STRING
	}

	n := NewNode(class, tag)
	n.SetString(s)

	return n, nil
}

func StringDeserialize(n *Node, tag int) (string, error) {

	class := CLASS_CONTEXT_SPECIFIC
	if tag < 0 {
		class = CLASS_UNIVERSAL
		tag = TAG_UTF8_STRING
	}

	err := CheckNode(n, class, tag)
	if err != nil {
		return "", err
	}

	return n.GetString()
}
