package node

type StrNode struct {
	TextValue string
	Value string
}

func (sNode StrNode) GetTextValue() string {
	return sNode.TextValue
}

func (sNode StrNode) GetType() Type {
	return TypeStr
}

func (sNode StrNode) GetValue() string{
	return sNode.Value
}

func NewStrNode(str string) ValueNode {
	return StrNode{TextValue:str, Value: str}
}