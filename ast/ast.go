package ast

import "monkey/token"

// Node - a base interface ofr AST nodes.
type Node interface {
	TokenLiteral() string
}

// Statement - a Node interface for Statements.
type Statement interface {
	Node
	statementNode()
}

// Expression - a Node interface for Expressions
type Expression interface {
	Node
	expressionNode()
}

// Program - The root node of every AST produced by the parser.
type Program struct {
	Statements []Statement
}

// TokenLiteral - returns the root node of the AST (Program.Statements[0])
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// LetStatement - an object that implements the Statement interface
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {} // dummy for compiler hint

// TokenLiteral - returns the Literal token from the LetStatement.
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// Identifier - object that implements the Expression interface
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {} // dummy for compiler hint

// TokenLiteral - returns the Literal for the IDENT token.
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
