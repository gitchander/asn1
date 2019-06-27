package der

import (
	"bytes"
	"encoding/hex"
	"fmt"
)

func ConvertToString(n *Node) (string, error) {
	var buf bytes.Buffer
	err := nodeToString(n, &buf, 0)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func nodeToString(n *Node, buf *bytes.Buffer, indent int) error {

	indentBuff := make([]byte, indent)
	for i := 0; i < indent; i++ {
		indentBuff[i] = '\t'
	}

	_, err := buf.Write(indentBuff)
	if err != nil {
		return err
	}

	className := classShortName(n.class)
	s := fmt.Sprintf("%s(%d):", className, n.GetTag())
	if _, err = buf.WriteString(s); err != nil {
		return err
	}

	if !n.constructed {

		buf.WriteByte(' ')

		s = hex.EncodeToString(n.data)
		if _, err = buf.WriteString(s); err != nil {
			return err
		}

		buf.WriteByte('\n')

	} else {

		buf.WriteString(" {\n")

		for _, child := range n.nodes {
			if err = nodeToString(child, buf, indent+1); err != nil {
				return err
			}
		}

		buf.Write(indentBuff)
		buf.WriteString("}\n")
	}

	return nil
}
