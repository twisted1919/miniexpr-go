package miniexpr

import (
	"fmt"
	"strconv"
)

// Parser holds parser state.
type Parser struct {
	tokens  TokenList
	current int
}

// NewParser creates a new parser.
func NewParser(tokens TokenList) *Parser {
	return &Parser{
		tokens:  tokens,
		current: 0,
	}
}

// Parse will parse the token into a final expression ready to be interpreted.
// nolint: nolintlint, ireturn
func (p *Parser) Parse() (Expr, error) {
	p.current = 0

	expr, err := p.expression()
	if err != nil {
		return nil, err
	}

	if !p.isAtEnd() {
		return nil, NewUnexpectedTokenError(p.peek().Literal[0], p.peek().Pos)
	}

	return expr, nil
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Type == TokenEOF
}

// nolint: nolintlint, unparam
func (p *Parser) advance() Token {
	if !p.isAtEnd() {
		p.current++
	}

	return p.previous()
}

func (p *Parser) peek() Token {
	return p.tokens[p.current]
}

func (p *Parser) previous() Token {
	return p.tokens[p.current-1]
}

func (p *Parser) match(types ...TokenType) bool {
	for _, t := range types {
		if !p.isAtEnd() && p.peek().Type == t {
			p.advance()

			return true
		}
	}

	return false
}

// nolint: nolintlint, ireturn
func (p *Parser) expression() (Expr, error) {
	return p.term()
}

// nolint: nolintlint, ireturn
func (p *Parser) term() (Expr, error) {
	expr, err := p.factor()
	if err != nil {
		return nil, err
	}

	for p.match(TokenPlus, TokenMinus) {
		prev := p.previous()

		fact, err := p.factor()
		if err != nil {
			return nil, err
		}

		expr = BinaryExpr{
			Left:     expr,
			Operator: prev,
			Right:    fact,
		}
	}

	return expr, nil
}

// nolint: nolintlint, ireturn
func (p *Parser) factor() (Expr, error) {
	expr, err := p.unary()
	if err != nil {
		return nil, err
	}

	for p.match(TokenStar, TokenSlash, TokenLeftShift, TokenRightShift) {
		prev := p.previous()

		pow, err := p.unary()
		if err != nil {
			return nil, err
		}

		expr = BinaryExpr{
			Left:     expr,
			Operator: prev,
			Right:    pow,
		}
	}

	return expr, nil
}

// nolint: nolintlint, ireturn
func (p *Parser) unary() (Expr, error) {
	for p.match(TokenMinus, TokenPlus) {
		prev := p.previous()

		exp, err := p.unary()
		if err != nil {
			return nil, err
		}

		// nolint: nolintlint, staticcheck
		return UnaryExpr{
			Operator: prev,
			Right:    exp,
		}, nil
	}

	return p.pow()
}

// nolint: nolintlint, ireturn
func (p *Parser) pow() (Expr, error) {
	expr, err := p.primary()
	if err != nil {
		return nil, err
	}

	if p.match(TokenPow, TokenStarStar) {
		prev := p.previous()

		unary, err := p.pow()
		if err != nil {
			return nil, err
		}

		expr = BinaryExpr{
			Left:     expr,
			Operator: prev,
			Right:    unary,
		}
	}

	return expr, nil
}

// nolint: nolintlint, ireturn
func (p *Parser) primary() (Expr, error) {
	for p.match(TokenNumber) {
		num, err := strconv.ParseFloat(p.previous().Literal, 64)
		if err != nil {
			return nil, fmt.Errorf("parse float failed: %w", err)
		}

		// nolint: nolintlint, staticcheck
		return LiteralExpr{Value: num}, nil
	}

	for p.match(TokenLeftParen) {
		expr, err := p.expression()
		if err != nil {
			return nil, err
		}

		_ = p.advance()
		if p.previous().Type != TokenRightParen {
			return nil, NewUnexpectedTokenError(p.previous().Literal[0], p.previous().Pos)
		}

		// nolint: nolintlint, staticcheck
		return GroupingExpr{Expr: expr}, nil
	}

	// this is a stray paren
	if p.match(TokenRightParen) {
		return nil, NewUnexpectedTokenError(')', p.previous().Pos)
	}

	return nil, ErrUnknownExpression
}
