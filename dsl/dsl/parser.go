package dsl

import (
	"fmt"
	"strings"
)

// ****************** AST节点定义 ******************
type Node any

type Program struct {
	Statements []Node
}

type IfStmt struct {
	Condition Node
	Then      *BlockStmt
	ElseIf    []*ElseIfStmt
	Else      *BlockStmt
}

type ElseIfStmt struct {
	Condition Node
	Body      *BlockStmt
}

type ForStmt struct {
	Init      Node
	Condition Node
	Update    Node
	Body      *BlockStmt
}

type BlockStmt struct {
	Statements []Node
}

type ReturnStmt struct {
	Value Node
}

type AssignStmt struct {
	Target string
	Value  Node
}

type BinaryExpr struct {
	Left  Node
	Op    TokenType
	Right Node
}

type Identifier struct {
	Name string
}

type NumberLiteral struct {
	Value string
}

type StringLiteral struct {
	Value string
}

type BoolLiteral struct {
	Value string
}

// ****************** 解析器实现 ******************
type Parser struct {
	tokens []Token
	curPos int
}

func NewParser(tokens []Token) *Parser {
	return &Parser{
		tokens: tokens,
		curPos: 0,
	}
}

func (p *Parser) Parse() (*Program, error) {
	program := &Program{}
	for p.peek().Type != TokenEOF {
		stmt, err := p.parseStat()
		if err != nil {
			return nil, err
		}
		program.Statements = append(program.Statements, stmt)
	}
	return program, nil
}

func (p *Parser) parseStat() (Node, error) {
	switch p.peek().Type {
	case TokenIf:
		return p.parseIfStat()
	case TokenFor:
		return p.parseForStat()
	case TokenReturn:
		return p.parseReturnStat()
	case TokenIdentifier:
		return p.parseAssignStat()
	case TokenLBrace:
		return p.parseBlock()
	default:
		return nil, fmt.Errorf("unexpected token: %v", p.peek())
	}
}

// 解析if语句
func (p *Parser) parseIfStat() (*IfStmt, error) {
	p.consume() // 跳过if

	condition := p.parseCondition()

	thenBlock, err := p.parseBlock()
	if err != nil {
		return nil, err
	}

	var elseIfStmts []*ElseIfStmt
	for p.peek().Type == TokenElseIf {
		p.consume()
		condition := p.parseCondition()
		block, err := p.parseBlock()
		if err != nil {
			return nil, err
		}
		elseIfStmt := &ElseIfStmt{
			Condition: condition,
			Body:      block,
		}
		elseIfStmts = append(elseIfStmts, elseIfStmt)
	}

	var elseBlock *BlockStmt
	if p.peek().Type == TokenElse {
		p.consume()
		elseBlock, err = p.parseBlock()
		if err != nil {
			return nil, err
		}
	}

	return &IfStmt{
		Condition: condition,
		Then:      thenBlock,
		ElseIf:    elseIfStmts,
		Else:      elseBlock,
	}, nil
}

// 解析for循环
func (p *Parser) parseForStat() (*ForStmt, error) {
	p.consume() // 跳过for

	// 解析初始化语句
	init, err := p.parseAssignStat()
	if err != nil {
		return nil, err
	}

	// 解析条件
	condition := p.parseCondition()

	// 解析更新语句
	update, err := p.parseAssignStat()
	if err != nil {
		return nil, err
	}

	// 解析循环体
	body, err := p.parseBlock()
	if err != nil {
		return nil, err
	}

	return &ForStmt{
		Init:      init,
		Condition: condition,
		Update:    update,
		Body:      body,
	}, nil
}

// 解析赋值语句
func (p *Parser) parseAssignStat() (*AssignStmt, error) {
	target := p.consume()
	if target.Type != TokenIdentifier {
		return nil, fmt.Errorf("expected identifier, got %v", target)
	}

	if !p.expect(TokenAssign) {
		return nil, fmt.Errorf("expected =")
	}

	// 使用完整的表达式解析
	value := p.parseExpression()

	//// 添加分号检查（根据语法需求可选）
	//if p.peek().Type == TokenSemicolon {
	//	p.consume()
	//}

	return &AssignStmt{
		Target: target.Value,
		Value:  value,
	}, nil
}

// 解析返回语句
func (p *Parser) parseReturnStat() (*ReturnStmt, error) {
	p.consume() // 跳过return

	// 使用完整表达式解析
	value := p.parseExpression()

	//// 添加分号检查（根据语法需求可选）
	//if p.peek().Type == TokenSemicolon {
	//	p.consume()
	//}

	return &ReturnStmt{Value: value}, nil
}

// 解析代码块
func (p *Parser) parseBlock() (*BlockStmt, error) {
	if !p.expect(TokenLBrace) {
		return nil, fmt.Errorf("expected {")
	}

	block := &BlockStmt{}
	for p.peek().Type != TokenRBrace && p.peek().Type != TokenEOF {
		stmt, err := p.parseStat()
		if err != nil {
			return nil, err
		}
		block.Statements = append(block.Statements, stmt)
	}

	if !p.expect(TokenRBrace) {
		return nil, fmt.Errorf("expected }")
	}

	return block, nil
}

// 解析条件表达式
func (p *Parser) parseCondition() Node {
	return p.parseExpression()
}

// 表达式解析优先级（从低到高）
// parseExpression -> parseLogicalExpr -> parseComparisonExpr ->
// parseAdditiveExpr -> parseMultiplicativeExpr -> parsePrimary

// 解析表达式
func (p *Parser) parseExpression() Node {
	return p.parseLogicalExpr()
}

// 解析逻辑表达式（&& ||）
func (p *Parser) parseLogicalExpr() Node {
	node := p.parseComparisonExpr()

	for {
		switch p.peek().Type {
		case TokenAnd, TokenOr:
			op := p.consume().Type
			right := p.parseComparisonExpr()
			node = &BinaryExpr{
				Left:  node,
				Op:    op,
				Right: right,
			}
		default:
			return node
		}
	}
}

// 解析比较表达式
func (p *Parser) parseComparisonExpr() Node {
	node := p.parseAdditiveExpr()

	for {
		switch p.peek().Type {
		case TokenEQ, TokenNEQ, TokenGT, TokenLT, TokenGE, TokenLE:
			op := p.consume().Type
			right := p.parseAdditiveExpr()
			node = &BinaryExpr{
				Left:  node,
				Op:    op,
				Right: right,
			}
		default:
			return node
		}
	}
}

// 解析加减表达式
func (p *Parser) parseAdditiveExpr() Node {
	node := p.parseMultiplicativeExpr()

	for {
		switch p.peek().Type {
		case TokenAdd, TokenSub:
			op := p.consume().Type
			right := p.parseMultiplicativeExpr()
			node = &BinaryExpr{
				Left:  node,
				Op:    op,
				Right: right,
			}
		default:
			return node
		}
	}
}

// 解析乘除表达式
func (p *Parser) parseMultiplicativeExpr() Node {
	node := p.parsePrimary()

	for {
		switch p.peek().Type {
		case TokenMul, TokenDiv:
			op := p.consume().Type
			right := p.parsePrimary()
			node = &BinaryExpr{
				Left:  node,
				Op:    op,
				Right: right,
			}
		default:
			return node
		}
	}
}

// 解析基础元素
func (p *Parser) parsePrimary() Node {
	token := p.consume()
	switch token.Type {
	case TokenIdentifier:
		return &Identifier{Name: token.Value}
	case TokenNumber:
		return &NumberLiteral{Value: token.Value}
	case TokenString:
		return &StringLiteral{Value: token.Value}
	case TokenLParen:
		expr := p.parseExpression()
		if !p.expect(TokenRParen) {
			panic("missing )")
		}
		return expr
	default:
		panic(fmt.Sprintf("unexpected token: %v", token))
	}
}

// ****************** 辅助方法 ******************
func (p *Parser) peek() Token {
	if p.curPos >= len(p.tokens) {
		return Token{Type: TokenEOF}
	}
	return p.tokens[p.curPos]
}

func (p *Parser) consume() Token {
	if p.curPos >= len(p.tokens) {
		return Token{Type: TokenEOF}
	}
	t := p.tokens[p.curPos]
	p.curPos++
	return t
}

func (p *Parser) expect(tt TokenType) bool {
	if p.peek().Type == tt {
		p.consume()
		return true
	}
	return false
}

// 递归打印AST结构
func PrintAST(node any, depth int) {
	prefix := strings.Repeat("  ", depth)
	switch n := node.(type) {
	case *Program:
		fmt.Printf("%sProgram:\n", prefix)
		for _, stmt := range n.Statements {
			PrintAST(stmt, depth+1)
		}
	case *ForStmt:
		fmt.Printf("%sFor:\n", prefix)
		fmt.Printf("%s  Init:\n", prefix)
		PrintAST(n.Init, depth+2)
		fmt.Printf("%s  Condition:\n", prefix)
		PrintAST(n.Condition, depth+2)
		fmt.Printf("%s  Update:\n", prefix)
		PrintAST(n.Update, depth+2)
		fmt.Printf("%s  Body:\n", prefix)
		PrintAST(n.Body, depth+2)
	case *IfStmt:
		fmt.Printf("%sIf:\n", prefix)
		fmt.Printf("%s  Condition:\n", prefix)
		PrintAST(n.Condition, depth+2)
		fmt.Printf("%s  Then:\n", prefix)
		PrintAST(n.Then, depth+2)
		if n.ElseIf != nil {
			for _, stmt := range n.ElseIf {
				fmt.Printf("%s  ElseIf:\n", prefix)
				fmt.Printf("%s    Condition:\n", prefix)
				PrintAST(stmt.Condition, depth+3)
				fmt.Printf("%s    Body:\n", prefix)
				PrintAST(stmt.Body, depth+3)
			}
		}
		if n.Else != nil {
			fmt.Printf("%s  Else:\n", prefix)
			PrintAST(n.Else, depth+2)
		}
	case *BlockStmt:
		fmt.Printf("%sBlock:\n", prefix)
		for _, stmt := range n.Statements {
			PrintAST(stmt, depth+1)
		}
	case *AssignStmt:
		fmt.Printf("%sAssign: %s = \n", prefix, n.Target)
		PrintAST(n.Value, depth+1)
	case *BinaryExpr:
		fmt.Printf("%sBinaryExpr: %v\n", prefix, n.Op)
		PrintAST(n.Left, depth+1)
		PrintAST(n.Right, depth+1)
	case *Identifier:
		fmt.Printf("%sIdentifier: %s\n", prefix, n.Name)
	case *NumberLiteral:
		fmt.Printf("%sNumber: %s\n", prefix, n.Value)
	case *StringLiteral:
		fmt.Printf("%sString: %s\n", prefix, n.Value)
	case *ReturnStmt:
		fmt.Printf("%sReturn:\n", prefix)
		PrintAST(n.Value, depth+1)
	default:
		fmt.Printf("%sUnknown node: %T\n", prefix, node)
	}
}
