package expr

import (
	"errors"
	. "fmt"
	"github.com/luoruiyi/expr/node"
	"go/ast"
	"go/parser"
)

// int string
// > < >= <= && ||
// in_arr(1, []int{1,2,3,4}), ver_compare(x, ">", "10.1.1") with no nested

type LogicEngine struct {
	ruleAst ast.Expr
}

func NewEngine(expr string) (*LogicEngine, error) {
	engine := &LogicEngine{
	}
	result, err := engine.UpdateAst(expr)
	if err != nil || !result {
		return nil, err
	}

	return engine, nil
}

func (engine *LogicEngine) UpdateAst(expr string) (bool, error) {
	if engine == nil {
		panic("please init the engine first")
	}
	exprAst, err := parser.ParseExpr(expr)
	if err != nil {
		Println(err)
		return false, err
	}
	engine.ruleAst = exprAst
	return true, nil
}

func (engine *LogicEngine) RunRule(controlMap map[string]interface{}) (bool, error) {
	if engine == nil || engine.ruleAst == nil {
		return false, errors.New("rule expr is empty, please init it")
	}

	nodeMap := parseControlMap(controlMap)
	value := eval(nodeMap, engine.ruleAst)
	bValue, ok := value.(node.BoolNode)
	if !ok {
		return false, errors.New(value.GetTextValue())
	}
	return bValue.True, nil
}

func parseControlMap(controlMap map[string]interface{}) map[string]node.ValueNode {
	nodeMap := make(map[string]node.ValueNode, len(controlMap))
	for key, control := range controlMap {
		switch control.(type) {
		case int:
			node := node.NewIntNode(int64(control.(int)))
			nodeMap[key] = node
		case int64:
			node := node.NewIntNode(control.(int64))
			nodeMap[key] = node
		case string:
			node := node.NewStrNode(control.(string))
			nodeMap[key] = node
		}
	}
	return nodeMap
}

func eval(mem map[string]node.ValueNode, expr ast.Expr) (y node.ValueNode) {
	switch x := expr.(type) {
	case *ast.BasicLit:
		return node.Lit2ValueNode(x)
	case *ast.BinaryExpr:
		a := eval(mem, x.X)
		b := eval(mem, x.Y)
		op := x.Op

		switch a.GetType() {
		case node.TypeInt64:
			return BinaryIntExpr{}.Invoke(a, b, op)
		case node.TypeBool:
			return BinaryBoolExpr{}.Invoke(a, b, op)
		case node.TypeStr:
			return BinaryStrExpr{}.Invoke(a, b, op)
		case node.TypeBad:
			return node.NewBadNode("a:" + a.GetTextValue() + "b:" + b.GetTextValue())
		}
		return node.NewBadNode(Sprintf("%d op is not suppoort", op))
	case *ast.CallExpr:
		name := x.Fun.(*ast.Ident).Name
		return CallExpr{name, x.Args}.Invoke(mem)
	case *ast.ParenExpr:
		return eval(mem, x.X)
	case *ast.Ident:
		return mem[x.Name]
	default:
		return node.NewBadNode(Sprintf("%x type is not suppoort", x))
	}

	panic("internal error")
}
