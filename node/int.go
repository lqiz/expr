package node

import (
	"fmt"
)

type IntNode struct {
	Value     int64
	TextValue string
}

func (iNode IntNode) GetTextValue() string {
	return iNode.TextValue
}

//func (iNode IntNode) GetValue() int64 {
//	return iNode.Value
//}

func (iNode IntNode) GetType() Type {
	return TypeInt64
}

//func (iNode IntNode) SetValue(str string) {
//	iNode.TextValue = str
//}

func NewIntNode(value int64) ValueNode {
	textValue := fmt.Sprintf("%d", value)
	return IntNode{Value: value, TextValue: textValue}
}

