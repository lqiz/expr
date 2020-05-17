package node

import "fmt"

type BoolNode struct {
	textValue string
	True      bool
}

func (bNode BoolNode) GetTextValue() string {
	return bNode.textValue
}

func (bNode BoolNode) GetValue() bool {
	return bNode.True
}

func (bNode BoolNode) GetType() Type {
	return TypeBool
}

func NewBoolNode(b bool) ValueNode {
	return BoolNode{
		True:      b,
		textValue: fmt.Sprintf("%t", b),
	}
}
