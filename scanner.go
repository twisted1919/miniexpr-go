package miniexpr

import (
	"strings"
)

// Scanner holds information about the current scanner.
type Scanner struct {
	source  DataSource
	tokens  TokenList
	current int
	pos     int
}

// NewScanner creates a new scanner.
func NewScanner(src DataSource) *Scanner {
	return &Scanner{
		source:  src,
		tokens:  make(TokenList, 0),
		pos:     0,
		current: 0,
	}
}

// Scan scans given input into tokens.
// nolint: nolintlint, gocognit, cyclop, funlen
func (s *Scanner) Scan() (TokenList, error) {
	s.tokens = s.tokens[:0]

	var char byte

	var ok bool // nolint: nolintlint, varnamelen

	for {
		char, ok = s.peek()
		if !ok {
			break
		}

		switch char {
		case '+':
			_, _ = s.advance()

			char, ok = s.peek()
			if !ok || char == '+' {
				s.pos++

				return nil, NewUnexpectedTokenError('+', s.pos)
			}

			s.tokens.Add(NewToken(TokenPlus, "+", s.pos))
		case '-':
			_, _ = s.advance()

			char, ok = s.peek()
			if !ok || char == '-' {
				s.pos++

				return nil, NewUnexpectedTokenError('-', s.pos)
			}

			s.tokens.Add(NewToken(TokenMinus, "-", s.pos))
		case '*':
			_, _ = s.advance()

			char, ok = s.peek()
			if !ok {
				return nil, NewUnexpectedTokenError('*', s.pos)
			}

			if char != '*' {
				s.tokens.Add(NewToken(TokenStar, "*", s.pos))
			} else {
				s.tokens.Add(NewToken(TokenStarStar, "**", s.pos))
				_, _ = s.advance()
			}
		case '^':
			_, _ = s.advance()
			s.tokens.Add(NewToken(TokenPow, "^", s.pos))
		case '/':
			_, _ = s.advance()
			s.tokens.Add(NewToken(TokenSlash, "/", s.pos))
		case '<':
			_, _ = s.advance()

			char, ok = s.peek()
			if !ok || char != '<' {
				return nil, NewUnexpectedTokenError('<', s.pos)
			}

			s.tokens.Add(NewToken(TokenLeftShift, "<<", s.pos))
			_, _ = s.advance()
		case '>':
			_, _ = s.advance()

			char, ok = s.peek()
			if !ok || char != '>' {
				return nil, NewUnexpectedTokenError('>', s.pos)
			}

			s.tokens.Add(NewToken(TokenRightShift, ">>", s.pos))
			_, _ = s.advance()
		case '(':
			_, _ = s.advance()
			s.tokens.Add(NewToken(TokenLeftParen, "(", s.pos))
		case ')':
			_, _ = s.advance()
			s.tokens.Add(NewToken(TokenRightParen, ")", s.pos))
		case ' ', '\t', '\n', '\r':
			_, _ = s.advance()
		default:
			if s.isDigit(char) {
				err := s.scanNumber(char)
				if err != nil {
					return nil, err
				}

				continue
			}

			return nil, NewUnexpectedTokenError(char, s.pos)
		}
	}

	s.tokens.Add(NewToken(TokenEOF, "\\0", s.pos))

	return s.tokens, nil
}

func (s *Scanner) isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= s.source.Len()
}

func (s *Scanner) peek() (byte, bool) {
	if s.isAtEnd() {
		return 0, false
	}

	return s.source.At(s.current), true
}

// nolint: nolintlint, unparam
func (s *Scanner) advance() (byte, bool) {
	if s.isAtEnd() {
		return 0, false
	}

	char := s.source.At(s.current)
	s.pos++
	s.current++

	return char, true
}

// nolint: nolintlint, cyclop, funlen
func (s *Scanner) scanNumber(char byte) error {
	num := strings.Builder{}
	num.WriteByte(char)

	var hasDot = false

	_, _ = s.advance()

	// nolint: nolintlint, varnamelen
	var ok bool

	for {
		char, ok = s.peek()
		// this is the end
		if !ok {
			break
		}

		if s.isDigit(char) {
			_, _ = s.advance()

			num.WriteByte(char)

			continue
		}

		if char == '.' {
			if hasDot {
				s.pos++

				return NewUnexpectedTokenError('.', s.pos)
			}

			hasDot = true

			num.WriteByte(char)

			_, _ = s.advance()

			char, ok = s.peek()
			if !ok || !s.isDigit(char) {
				return NewUnexpectedTokenError(char, s.pos)
			}

			continue
		}

		if char == '_' {
			_, _ = s.advance()

			char, ok = s.peek()
			if !ok || !s.isDigit(char) {
				return NewUnexpectedTokenError(char, s.pos)
			}

			continue
		}

		break
	}

	s.tokens.Add(NewToken(TokenNumber, num.String(), s.pos))

	return nil
}
