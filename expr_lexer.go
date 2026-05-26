package actionlint

import (
	"text/scanner"
)

// TokenKind is kind of token.
type TokenKind int

const (
	// TokenKindUnknown is a default value of token as unknown token value.
	TokenKindUnknown TokenKind = iota
	// TokenKindEnd is a token for end of token sequence. Sequence without this
	// token means invalid.
	TokenKindEnd
	// TokenKindIdent is a token for identifier.
	TokenKindIdent
	// TokenKindString is a token for string literals.
	TokenKindString
	// TokenKindInt is a token for integers including hex integers.
	TokenKindInt
	// TokenKindFloat is a token for float numbers.
	TokenKindFloat
	// TokenKindLeftParen is a token for '('.
	TokenKindLeftParen
	// TokenKindRightParen is a token for ')'.
	TokenKindRightParen
	// TokenKindLeftBracket is a token for '['.
	TokenKindLeftBracket
	// TokenKindRightBracket is a token for ']'.
	TokenKindRightBracket
	// TokenKindDot is a token for '.'.
	TokenKindDot
	// TokenKindNot is a token for '!'.
	TokenKindNot
	// TokenKindLess is a token for '<'.
	TokenKindLess
	// TokenKindLessEq is a token for '<='.
	TokenKindLessEq
	// TokenKindGreater is a token for '>'.
	TokenKindGreater
	// TokenKindGreaterEq is a token for '>='.
	TokenKindGreaterEq
	// TokenKindEq is a token for '=='.
	TokenKindEq
	// TokenKindNotEq is a token for '!='.
	TokenKindNotEq
	// TokenKindAnd is a token for '&&'.
	TokenKindAnd
	// TokenKindOr is a token for '||'.
	TokenKindOr
	// TokenKindStar is a token for '*'.
	TokenKindStar
	// TokenKindComma is a token for ','.
	TokenKindComma
)

func (t TokenKind) String() string { _ = "STUB: not implemented"; return "" }

// Token is a token lexed from expression syntax. For more details, see
// https://docs.github.com/en/actions/learn-github-actions/expressions
type Token struct {
	// Kind is kind of the token.
	Kind TokenKind
	// Value is string representation of the token.
	Value string
	// Offset is byte offset of token string starting.
	Offset int
	// Line is line number of start position of the token. Note that this value is 1-based.
	Line int
	// Column is column number of start position of the token. Note that this value is 1-based.
	Column int
}

func (t *Token) String() string { _ = "STUB: not implemented"; return "" }

func isWhitespace(r rune) bool { _ = "STUB: not implemented"; return false }

func isAlpha(r rune) bool { _ = "STUB: not implemented"; return false }

func isNum(r rune) bool { _ = "STUB: not implemented"; return false }

func isHexNum(r rune) bool { _ = "STUB: not implemented"; return false }

func isAlnum(r rune) bool { _ = "STUB: not implemented"; return false }

const expectedPunctChars = "''', '}', '(', ')', '[', ']', '.', '!', '<', '>', '=', '&', '|', '*', ',', ' '"
const expectedDigitChars = "'0'..'9'"
const expectedAlphaChars = "'a'..'z', 'A'..'Z', '_'"
const expectedAllChars = expectedAlphaChars + ", " + expectedDigitChars + ", " + expectedPunctChars

// ExprLexer is a struct to lex expression syntax. To know the syntax, see
// https://docs.github.com/en/actions/learn-github-actions/expressions
type ExprLexer struct {
	src    string
	scan   scanner.Scanner
	lexErr *ExprError
	start  scanner.Position
}

// NewExprLexer makes new ExprLexer instance.
func NewExprLexer(src string) *ExprLexer { _ = "STUB: not implemented"; return nil }

func (lex *ExprLexer) error(msg string) { _ = "STUB: not implemented"; return }

func (lex *ExprLexer) token(kind TokenKind) *Token { _ = "STUB: not implemented"; return nil }

func (lex *ExprLexer) eof() *Token { _ = "STUB: not implemented"; return nil }

func (lex *ExprLexer) eat() rune { _ = "STUB: not implemented"; return 0 }

// unlike lex.scan.Next(), return top char *after* eating

func (lex *ExprLexer) skipWhite() { _ = "STUB: not implemented"; return }

func (lex *ExprLexer) unexpected(r rune, where string, expected string) *Token {
	_ = "STUB: not implemented"
	return nil
}

func (lex *ExprLexer) unexpectedEOF() *Token { _ = "STUB: not implemented"; return nil }

func (lex *ExprLexer) lexIdent() *Token {
	_ = "STUB: not implemented"

	// a-z, A-Z, 0-9, - or _
	// https://docs.github.com/en/actions/learn-github-actions/contexts
	return nil
}

func (lex *ExprLexer) lexNum() *Token {
	_ = "STUB: not implemented"
	// The official document says number literals are 'Any number format supported by JSON' but actually
	// hex numbers starting with 0x are supported.
	return nil
}

// precond: r is digit or '-'

// r is 1..9

// eat '.'

// eat 'e' or 'E'

// eat the '0'

// r is 1..9

func (lex *ExprLexer) lexHexInt() *Token { _ = "STUB: not implemented"; return nil }

// Note: GitHub Actions does not support exponent part like 0x1f2p-a8

func (lex *ExprLexer) lexString() *Token {
	_ = "STUB: not implemented"
	// precond: current char is '
	return nil
}

// when not escaped single quote ''

func (lex *ExprLexer) lexEnd() *Token {
	_ = "STUB: not implemented"
	// eat the first '}'
	return nil
}

// }} is an end marker of interpolation

func (lex *ExprLexer) lexLess() *Token { _ = "STUB: not implemented"; return nil }

// eat '<'

func (lex *ExprLexer) lexGreater() *Token { _ = "STUB: not implemented"; return nil }

// eat '>'

func (lex *ExprLexer) lexEq() *Token { _ = "STUB: not implemented"; return nil }

// eat '='

func (lex *ExprLexer) lexBang() *Token { _ = "STUB: not implemented"; return nil }

// eat '!'
// eat '='

func (lex *ExprLexer) lexAnd() *Token { _ = "STUB: not implemented"; return nil }

// eat the first '&'

// eat the second '&'

func (lex *ExprLexer) lexOr() *Token { _ = "STUB: not implemented"; return nil }

// eat the first '|'

// eat the second '|'

func (lex *ExprLexer) lexChar(k TokenKind) *Token { _ = "STUB: not implemented"; return nil }

// Next lexes next token to lex input incrementally. Lexer must be initialized with Init() method
// before the first call of this method. This method is stateful. Lexer advances offset by lexing
// token. To get the offset, use Offset() method.
func (lex *ExprLexer) Next() *Token { _ = "STUB: not implemented"; return nil }

// Ident starts with a-z or A-Z or _
// https://docs.github.com/en/actions/learn-github-actions/contexts

// Offset returns the current offset (scanning position).
func (lex *ExprLexer) Offset() int { _ = "STUB: not implemented"; return 0 }

// Err returns an error while lexing. When multiple errors occur, the first one is returned.
func (lex *ExprLexer) Err() *ExprError {
	_ = "STUB: not implemented"

	// LexExpression lexes the given string as expression syntax. The parameter must contain '}}' which
	// represents end of expression. Otherwise this function will report an error that it encountered
	// unexpected EOF.
	return nil
}

func LexExpression(src string) ([]*Token, int, *ExprError) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}
