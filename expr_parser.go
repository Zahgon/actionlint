package actionlint

func errorAtToken(t *Token, msg string) *ExprError { _ = "STUB: not implemented"; return nil }

// ExprParser is a parser for expression syntax. To know the details, see
// https://docs.github.com/en/actions/learn-github-actions/expressions
type ExprParser struct {
	cur   *Token
	lexer *ExprLexer
	err   *ExprError
}

// NewExprParser creates new ExprParser instance.
func NewExprParser() *ExprParser { _ = "STUB: not implemented"; return nil }

func (p *ExprParser) error(msg string) { _ = "STUB: not implemented"; return }

func (p *ExprParser) errorf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (p *ExprParser) unexpected(where string, expected []TokenKind) {
	_ = "STUB: not implemented"
	return
}

func (p *ExprParser) next() *Token { _ = "STUB: not implemented"; return nil }

func (p *ExprParser) peek() *Token { _ = "STUB: not implemented"; return nil }

func (p *ExprParser) parseIdent() ExprNode {
	_ = "STUB: not implemented"
	// eat ident
	return *new(ExprNode)
}

// Parse function call as primary expression though generally function call is parsed as
// postfix expression. The reason is that only built-in function call is allowed in workflow
// expression syntax, meant that callee is always built-in function name, not a general
// expression.
// eat '('

// no arguments
// eat ')'

// eat ','
// continue to next argument

// eat ')'

// Handle keywords. Note that keywords are case sensitive. TRUE, FALSE, NULL are invalid named value.

// Variable name access is case insensitive. github.event and GITHUB.event are the same.

func (p *ExprParser) parseNestedExpr() ExprNode {
	_ = "STUB: not implemented"
	// eat '('
	return *new(ExprNode)
}

// eat ')'

func (p *ExprParser) parseInt() ExprNode { _ = "STUB: not implemented"; return *new(ExprNode) }

// eat int

func (p *ExprParser) parseFloat() ExprNode { _ = "STUB: not implemented"; return *new(ExprNode) }

// eat float

func (p *ExprParser) parseString() ExprNode {
	_ = "STUB: not implemented"
	// eat string
	return *new(ExprNode)
}

// strip first and last single quotes
// unescape ''

func (p *ExprParser) parsePrimaryExpr() ExprNode { _ = "STUB: not implemented"; return *new(ExprNode) }

func (p *ExprParser) parsePostfixOp() ExprNode { _ = "STUB: not implemented"; return *new(ExprNode) }

// eat '.'

// eat '*'

// eat 'b' of 'a.b'
// Property name is case insensitive. github.event and github.EVENT are the same

// eat '['

// eat ']'

func (p *ExprParser) parsePrefixOp() ExprNode { _ = "STUB: not implemented"; return *new(ExprNode) }

// eat '!' token

func (p *ExprParser) parseCompareBinOp() ExprNode { _ = "STUB: not implemented"; return *new(ExprNode) }

// eat the operator token

func (p *ExprParser) parseLogicalAnd() ExprNode { _ = "STUB: not implemented"; return *new(ExprNode) }

// eat &&

func (p *ExprParser) parseLogicalOr() ExprNode { _ = "STUB: not implemented"; return *new(ExprNode) }

// eat ||

// Err returns an error which was caused while previous parsing.
func (p *ExprParser) Err() *ExprError { _ = "STUB: not implemented"; return nil }

// Parse parses token sequence lexed by a given lexer into syntax tree.
func (p *ExprParser) Parse(l *ExprLexer) (ExprNode, *ExprError) {
	_ = "STUB: not implemented"
	// Init
	return *new(ExprNode), nil
}

// It did not reach the end of sequence
