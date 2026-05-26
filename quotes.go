package actionlint

import (
	"strings"
)

type quotesBuilder struct {
	inner strings.Builder
	buf   []byte
	comma bool
}

func (b *quotesBuilder) append(s string) { _ = "STUB: not implemented"; return }

func (b *quotesBuilder) appendRune(r rune) { _ = "STUB: not implemented"; return }

func (b *quotesBuilder) build() string { _ = "STUB: not implemented"; return "" }

func quotes(ss []string) string { _ = "STUB: not implemented"; return "" }

// 2 for delims

// comma

func sortedQuotes(ss []string) string { _ = "STUB: not implemented"; return "" }

func quotesAll(sss ...[]string) string { _ = "STUB: not implemented"; return "" }

// 2 for delims

// comma

// comma
