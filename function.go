package expr

import (
	"fmt"
	"github.com/lqiz/expr/node"
	"go/ast"
	"go/token"
	"strings"
)

type BinaryBoolExpr struct{}

type BinaryStrExpr struct{}

type BinaryIntExpr struct{}

type CallExpr struct {
	fn   string // one of "in_array", "ver_compare"
	args []ast.Expr
}

func (b BinaryBoolExpr) Invoke(x, y node.ValueNode, op token.Token) node.ValueNode {
	xb, xok := x.(node.BoolNode)
	yb, yok := y.(node.BoolNode)

	if !xok || !yok {
		return node.NewBadNode(x.GetTextValue() + y.GetTextValue())
	}

	switch op {
	case token.LAND:
		return node.NewBoolNode(xb.True && yb.True)
	case token.LOR:
		return node.NewBoolNode(xb.True || yb.True)
	}
	return node.NewBadNode(fmt.Sprintf("unsupported binary operator: %s", op.String()))
}

func (b BinaryStrExpr) Invoke(x, y node.ValueNode, op token.Token) node.ValueNode {
	xs, xok := x.(node.StrNode)
	ys, yok := y.(node.StrNode)

	if !xok || !yok {
		return node.NewBadNode("x: " + x.GetTextValue() + "y: " + y.GetTextValue())
	}

	switch op {
	case token.EQL: // ==
		return node.NewBoolNode(strings.Compare(xs.GetValue(), ys.GetValue()) == 0)
	case token.LSS: // <
		return node.NewBoolNode(strings.Compare(xs.GetValue(), ys.GetValue()) == -1)
	case token.GTR: // >
		return node.NewBoolNode(strings.Compare(xs.GetValue(), ys.GetValue()) == +1)
	case token.GEQ: // >=
		return node.NewBoolNode(strings.Compare(xs.GetValue(), ys.GetValue()) >= 0)
	case token.LEQ: // <=
		return node.NewBoolNode(strings.Compare(xs.GetValue(), ys.GetValue()) <= 0)
	}
	return node.NewBadNode(fmt.Sprintf("unsupported binary operator: %s", op.String()))
}

func (b BinaryIntExpr) Invoke(x, y node.ValueNode, op token.Token) node.ValueNode {
	xs, xok := x.(node.IntNode)
	ys, yok := y.(node.IntNode)

	if !xok || !yok {
		return node.NewBadNode(x.GetTextValue() + y.GetTextValue())
	}

	switch op {
	case token.EQL: // ==
		return node.BoolNode{
			True: xs.Value == ys.Value,
		}
	case token.LSS: // <
		return node.BoolNode{
			True: xs.Value < ys.Value,
		}
	case token.GTR: // >
		return node.NewBoolNode(xs.Value > ys.Value)
	case token.GEQ: // >=
		return node.NewBoolNode(xs.Value >= ys.Value)
	case token.LEQ: // <=
		return node.NewBoolNode(xs.Value <= ys.Value)
	}
	return node.NewBadNode(fmt.Sprintf("unsupported binary operator: %s", op.String()))
}

func (c CallExpr) Invoke(mem map[string]node.ValueNode) node.ValueNode {
	switch c.fn {
	case "in_array":
		parm := eval(mem, c.args[0])
		if parm.GetType() == node.TypeBad {
			return parm
		}
		vRange, ok := c.args[1].(*ast.CompositeLit)
		if !ok {
			return node.NewBadNode("func in_array 2ed params is not a composite lit")
		}
		eltNodes := make([]node.ValueNode, 0, len(vRange.Elts))
		for _, p := range vRange.Elts {
			elt := eval(mem, p)
			eltNodes = append(eltNodes, elt)
		}

		has := false
		for _, v := range eltNodes {
			if v.GetType() == parm.GetType() && v.GetTextValue() == parm.GetTextValue() {
				has = true
			}
		}

		return node.NewBoolNode(has)
	case "ver_compare":
		if len(c.args) != 3 {
			return node.NewBadNode("func ver_compare doesn't have enough params")
		}

		args := make([]string, 0, 3)
		for _, v := range c.args {
			arg := eval(mem, v)
			if arg.GetType() != node.TypeStr {
				return node.NewBadNode("func ver_compare params type error")
			}
			args = append(args, arg.GetTextValue())
		}

		ret := VersionCompare(args[0], args[1], args[2])
		return node.NewBoolNode(ret)
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}
