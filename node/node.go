package node

import (
	"fmt"
	"go/ast"
	"go/token"
	"strconv"
)

type Type int

const (
	TypeInt64 Type = iota
	TypeStr
	TypeBool
	TypeBad
)

type ValueNode interface {
	GetType() Type
	GetTextValue() string
	//SetValue(string)
}

func Lit2ValueNode(lit *ast.BasicLit) ValueNode {
	switch lit.Kind {
	case token.INT:
		value, err := strconv.ParseInt(lit.Value, 10, 64)
		if err != nil {
			return NewBadNode(err.Error())
		}
		return NewIntNode(value)
	case token.STRING:
		value, err := strconv.Unquote(lit.Value)
		if err != nil {
			return NewBadNode(err.Error())
		}
		return NewStrNode(value)
	}

	return NewBadNode(fmt.Sprintf("%s is not support type", lit.Kind))
}
