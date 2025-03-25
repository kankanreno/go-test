package dsl

import (
	"fmt"
	"github.com/spf13/cast"
	"strconv"
	"strings"
)

// 执行环境
type Environment struct {
	variables map[string]any
	parent    *Environment // 支持作用域链
}

func NewEnvironment(v map[string]any) *Environment {
	return &Environment{
		variables: v,
	}
}

// 执行器结构
type Executor struct {
	globalEnv *Environment
}

func NewExecutor() *Executor {
	return &Executor{}
}

// 设置环境变量
func (e *Executor) SetEnvironment(env *Environment) {
	e.globalEnv = env
}

// 主执行方法
func (e *Executor) Execute(nodes []Node) (any, error) {
	var lastResult any
	for _, node := range nodes {
		result, err := e.eval(node, e.globalEnv)
		if err != nil {
			return nil, err
		}
		if result != nil {
			lastResult = result
		}
	}
	return lastResult, nil
}

// 递归求值方法
// TODO: return及func的处理 http://showdoc.approj.com/web/#/p/92da9410a1a7020c7188e510ddf8e587
func (e *Executor) eval(node Node, env *Environment) (any, error) {
	switch n := node.(type) {
	// 字面量
	case *NumberLiteral:
		if strings.Contains(n.Value, ".") {
			f, _ := strconv.ParseFloat(n.Value, 64)
			return f, nil
		}
		i, _ := strconv.Atoi(n.Value)
		return i, nil
	case *StringLiteral:
		return n.Value, nil
	case *BoolLiteral:
		return n.Value, nil

	// 标识符
	case *Identifier:
		if val, ok := env.variables[n.Name]; ok {
			return val, nil
		}
		return nil, fmt.Errorf("undefined variable: %s", n.Name)

	// 二元表达式
	case *BinaryExpr:
		left, err := e.eval(n.Left, env)
		if err != nil {
			return nil, err
		}
		right, err := e.eval(n.Right, env)
		if err != nil {
			return nil, err
		}

		return e.evalBinaryOp(n.Op, left, right)

	// 赋值语句
	case *AssignStmt:
		val, err := e.eval(n.Value, env)
		if err != nil {
			return nil, err
		}
		env.variables[n.Target] = val
		return val, nil

	// 代码块
	case *BlockStmt:
		blockEnv := &Environment{
			parent:    env,
			variables: make(map[string]any),
		}
		var result any
		for _, stmt := range n.Statements {
			var err error
			result, err = e.eval(stmt, blockEnv)
			if err != nil {
				return nil, err
			}
		}
		return result, nil

	// 条件语句
	case *IfStmt:
		cond, err := e.eval(n.Condition, env)
		if err != nil {
			return nil, err
		}
		if cast.ToBool(cond) {
			return e.eval(n.Then, env)
		} else if n.ElseIf != nil {
			for _, stmt := range n.ElseIf {
				cond, err := e.eval(stmt.Condition, env)
				if err != nil {
					return nil, err
				}
				if cast.ToBool(cond) {
					return e.eval(stmt.Body, env)
				}
			}
			return e.eval(n.Else, env)
		} else if n.Else != nil {
			return e.eval(n.Else, env)
		}
		return nil, nil

	//// 函数调用（示例）
	//case *CallExpr:
	//	fn, err := e.eval(n.Func, env)
	//	if err != nil {
	//		return nil, err
	//	}
	//	// 此处可以添加内置函数处理

	// 处理返回值
	case *ReturnStmt:
		val, err := e.eval(n.Value, env)
		if err != nil {
			return nil, err
		}
		return val, nil
	}

	return nil, fmt.Errorf("unsupported node type: %T", node)
}

// 二元运算处理
func (e *Executor) evalBinaryOp(op TokenType, left, right any) (any, error) {
	switch op {
	case TokenAdd:
		if l, ok := left.(int); ok {
			if r, ok := right.(int); ok {
				return l + r, nil
			}
		}
		return cast.ToFloat64(left) + cast.ToFloat64(right), nil
	case TokenSub:
		return cast.ToFloat64(left) - cast.ToFloat64(right), nil
	case TokenMul:
		return cast.ToFloat64(left) * cast.ToFloat64(right), nil
	case TokenDiv:
		return cast.ToFloat64(left) / cast.ToFloat64(right), nil
	case TokenEQ:
		return equals(left, right), nil
	case TokenNEQ:
		return !equals(left, right), nil
	case TokenGT:
		return cast.ToFloat64(left) > cast.ToFloat64(right), nil
	case TokenLT:
		return cast.ToFloat64(left) < cast.ToFloat64(right), nil
	case TokenGE:
		return cast.ToFloat64(left) >= cast.ToFloat64(right), nil
	case TokenLE:
		return cast.ToFloat64(left) <= cast.ToFloat64(right), nil
	case TokenAnd:
		return cast.ToBool(left) && cast.ToBool(right), nil
	case TokenOr:
		return cast.ToBool(left) || cast.ToBool(right), nil
	default:
		return nil, fmt.Errorf("unknown operator: %v", op)
	}
}

func equals(a, b any) bool {
	switch b := b.(type) {
	case int:
		return cast.ToInt(a) == b
	case float64:
		return cast.ToFloat64(a) == b
	case string:
		return cast.ToString(a) == b
	case bool:
		return cast.ToBool(a) == b
	}
	return false
}
