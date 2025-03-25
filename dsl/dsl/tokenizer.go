package dsl

import "unicode"

// ****************** Token类型定义 ******************
type TokenType int

const (
	TokenEOF        TokenType = iota
	TokenErr        TokenType = iota
	TokenIdentifier           // 标识符
	TokenNumber               // 数字
	TokenString               // 字符串
	TokenTrue                 // true
	TokenFalse                // false
	TokenAnd                  // &&
	TokenOr                   // ||
	TokenEQ                   // ==
	TokenNEQ                  // !=
	TokenGT                   // >
	TokenLT                   // <
	TokenGE                   // >=
	TokenLE                   // <=
	TokenAssign               // =
	TokenAdd                  // +
	TokenSub                  // -
	TokenMul                  // *
	TokenDiv                  // /
	TokenLParen               // (
	TokenRParen               // )
	TokenLBrace               // {
	TokenRBrace               // }
	TokenSemicolon            // ;
	TokenComma                // ,
	TokenIf                   // if
	TokenElseIf               // if
	TokenElse                 // else
	TokenFor                  // for
	TokenReturn               // return
	TokenBreak                // break
	TokenContinue             // continue
)

type Token struct {
	Type  TokenType
	Value string
	Pos   int // 在源码中的位置（可选）
}

// ****************** 分词器实现 ******************
// 分词器结构
type Tokenizer struct {
	input   []rune // 输入字符流
	curPos  int    // 当前读取位置
	nextPos int    // 下一个读取位置
	curChar rune   // 当前字符
}

func NewTokenizer(input string) *Tokenizer {
	t := &Tokenizer{
		input: []rune(input),
	}
	t.readChar()
	return t
}

func (t *Tokenizer) Tokenize() []Token {
	var tokens []Token
	for t.curChar != 0 {
		t.skipWhitespace()

		var tok Token
		switch t.curChar {
		case '=':
			if t.peekChar() == '=' {
				t.readChar()
				tok = Token{Type: TokenEQ, Value: "==", Pos: t.curPos}
			} else {
				tok = Token{Type: TokenAssign, Value: "=", Pos: t.curPos}
			}
		case '!':
			if t.peekChar() == '=' {
				t.readChar()
				tok = Token{Type: TokenNEQ, Value: "!=", Pos: t.curPos}
			}
		case '>':
			if t.peekChar() == '=' {
				t.readChar()
				tok = Token{Type: TokenGE, Value: ">=", Pos: t.curPos}
			} else {
				tok = Token{Type: TokenGT, Value: ">", Pos: t.curPos}
			}
		case '<':
			if t.peekChar() == '=' {
				t.readChar()
				tok = Token{Type: TokenLE, Value: "<=", Pos: t.curPos}
			} else {
				tok = Token{Type: TokenLT, Value: "<", Pos: t.curPos}
			}
		case '&':
			if t.peekChar() == '&' {
				t.readChar()
				tok = Token{Type: TokenAnd, Value: "&&", Pos: t.curPos}
			}
		case '|':
			if t.peekChar() == '|' {
				t.readChar()
				tok = Token{Type: TokenOr, Value: "||", Pos: t.curPos}
			}
		case '+':
			tok = Token{Type: TokenAdd, Value: "+", Pos: t.curPos}
		case '-':
			tok = Token{Type: TokenSub, Value: "-", Pos: t.curPos}
		case '*':
			tok = Token{Type: TokenMul, Value: "*", Pos: t.curPos}
		case '/':
			tok = Token{Type: TokenDiv, Value: "/", Pos: t.curPos}
		case '(':
			tok = Token{Type: TokenLParen, Value: "(", Pos: t.curPos}
		case ')':
			tok = Token{Type: TokenRParen, Value: ")", Pos: t.curPos}
		case '{':
			tok = Token{Type: TokenLBrace, Value: "{", Pos: t.curPos}
		case '}':
			tok = Token{Type: TokenRBrace, Value: "}", Pos: t.curPos}
		case ';':
			tok = Token{Type: TokenSemicolon, Value: ";", Pos: t.curPos}
		case ',':
			tok = Token{Type: TokenComma, Value: ",", Pos: t.curPos}
		case 0:
			tok = Token{Type: TokenEOF, Value: "", Pos: t.curPos}
		case '"':
			tok = Token{Type: TokenString, Value: t.readString(), Pos: t.curPos}
		default:
			if isLetter(t.curChar) {
				ident := t.readIdentifier()
				switch ident {
				case "if":
					tok = Token{Type: TokenIf, Value: ident, Pos: t.curPos}
				case "elif":
					tok = Token{Type: TokenElseIf, Value: ident, Pos: t.curPos}
				case "else":
					tok = Token{Type: TokenElse, Value: ident, Pos: t.curPos}
				case "for":
					tok = Token{Type: TokenFor, Value: ident, Pos: t.curPos}
				case "return":
					tok = Token{Type: TokenReturn, Value: ident, Pos: t.curPos}
				case "true":
					tok = Token{Type: TokenTrue, Value: ident, Pos: t.curPos}
				case "false":
					tok = Token{Type: TokenFalse, Value: ident, Pos: t.curPos}
				default:
					tok = Token{Type: TokenIdentifier, Value: ident, Pos: t.curPos}
				}
			} else if isDigit(t.curChar) {
				tok = Token{Type: TokenNumber, Value: t.readNumber(), Pos: t.curPos}
			} else {
				tok = Token{Type: TokenErr, Value: string(t.curChar), Pos: t.curPos}
			}
		}
		tokens = append(tokens, tok)
		t.readChar()
	}
	return tokens
}

// ****************** 辅助方法 ******************
func (t *Tokenizer) readChar() {
	if t.nextPos >= len(t.input) {
		t.curChar = 0
	} else {
		t.curChar = t.input[t.nextPos]
	}
	t.curPos = t.nextPos
	t.nextPos++
}

func (t *Tokenizer) peekChar() rune {
	if t.nextPos >= len(t.input) {
		return 0
	}
	return t.input[t.nextPos]
}

func (t *Tokenizer) skipWhitespace() {
	for t.curChar == ' ' || t.curChar == '\t' || t.curChar == '\n' || t.curChar == '\r' {
		t.readChar()
	}
}

func (t *Tokenizer) readIdentifier() string {
	start := t.curPos
	for isLetter(t.curChar) && (isLetter(t.peekChar()) || isDigit(t.peekChar())) {
		t.readChar()
	}
	return string(t.input[start : t.curPos+1])
}

func (t *Tokenizer) readNumber() string {
	start := t.curPos
	for (isDigit(t.curChar) && (isDigit(t.peekChar()) || t.peekChar() == '.')) || (t.curChar == '.' && isDigit(t.peekChar())) {
		t.readChar()
	}
	return string(t.input[start : t.curPos+1])
}

func (t *Tokenizer) readString() string {
	start := t.curPos
	t.readChar() // 跳过起始的"
	for t.curChar != '"' && t.curChar != 0 {
		t.readChar()
	}
	return string(t.input[start+1 : t.curPos]) // 排除两个引号
}

func isLetter(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_'
}

func isDigit(ch rune) bool {
	return unicode.IsDigit(ch)
}
