package miniexpr

// TokenType is the numeric representation of the token type.
type TokenType int

const (
	// TokenPlus represents the addition operator: "+".
	TokenPlus TokenType = iota

	// TokenMinus represents the subtraction operator: "-".
	TokenMinus

	// TokenStar represents the multiplication operator: "*".
	TokenStar

	// TokenStarStar represents the exponentiation operator: "**".
	// (the alternative form of power, equivalent to TokenPow).
	TokenStarStar

	// TokenPow represents the exponentiation operator: "^".
	TokenPow

	// TokenSlash represents the division operator: "/".
	TokenSlash

	// TokenLeftShift represents the left-shift operator: "<<".
	TokenLeftShift

	// TokenRightShift represents the right-shift operator: ">>".
	TokenRightShift

	// TokenLeftParen represents the left parenthesis: "(".
	TokenLeftParen

	// TokenRightParen represents the right parenthesis: ")".
	TokenRightParen

	// TokenNumber represents a numeric literal (integer or decimal).
	TokenNumber

	// TokenEOF signals the end of input.
	TokenEOF
)

// Token represents a single token.
type Token struct {
	Type    TokenType
	Literal string
	Pos     int
}

// TokenList is a list of tokens.
type TokenList []Token

// Add adds a new token to the list.
func (t *TokenList) Add(token Token) {
	*t = append(*t, token)
}

// NewToken creates a new Token.
func NewToken(typ TokenType, literal string, pos int) Token {
	return Token{
		Type:    typ,
		Literal: literal,
		Pos:     pos,
	}
}
