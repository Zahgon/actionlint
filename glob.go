package actionlint

import (
	"fmt"
	"strings"
	"text/scanner"
)

// Note:
// - Broken pattern causes a syntax error
//   - '+' or '?' at top of pattern
//   - preceding character of '+' or '?' is special character like '+', '?', '*'
//   - Missing ] in [...] pattern like '[0-9'
//   - Missing end of range in [...] like '[0-]'
// - \ can escape special characters like '['. Otherwise \ is handled as normal character
// - invalid characters for Git ref names are not checked on GitHub Actions runtime
//   - `man git-check-ref-format` for more details
//   - \ is invalid character for ref names. it means that \ can be used only for escaping special chars

// InvalidGlobPattern is an error on invalid glob pattern.
type InvalidGlobPattern struct {
	// Message is a human readable error message.
	Message string
	// Column is a column number of the error in the glob pattern. This value is 1-based, but zero
	// is valid value. Zero means the error occurred before reading first character. This happens
	// when a given pattern is empty. When the given pattern include a newline and line number
	// increases (invalid pattern), the column number falls back into always 0.
	Column int
}

func (err *InvalidGlobPattern) Error() string { _ = "STUB: not implemented"; return "" }

func (err *InvalidGlobPattern) String() string { _ = "STUB: not implemented"; return "" }

type globValidator struct {
	isRef bool
	prec  bool
	errs  []InvalidGlobPattern
	scan  scanner.Scanner
}

func (v *globValidator) error(msg string) {
	_ = "STUB: not implemented"

	// - 1 because character at the error position is already eaten from scanner
	return
}

// fallback to 0

func (v *globValidator) unexpected(char rune, what, why string) { _ = "STUB: not implemented"; return }

func (v *globValidator) invalidRefChar(c rune, why string) { _ = "STUB: not implemented"; return }

// avoid '\\'

func (v *globValidator) init(pat string) {
	v.errs = []InvalidGlobPattern{}
	v.prec = false
	v.scan.Init(strings.NewReader(pat))
	v.scan.Error = func(s *scanner.Scanner, m string) {
		v.error(fmt.Sprintf("error while scanning glob pattern %q: %s", pat, m))
	}
}

func (v *globValidator) validateNext() bool { _ = "STUB: not implemented"; return false }

// eat escaped character

// eat escaped character

// file path can contain '\' (`mkdir 'foo\bar'` works)

// eat ]

// in case of single character

// When match is range of character like 0-9

// actually one or more. but this is ok since we only check chars > 1 later

//lint:ignore SA4006 c should always holds the current character even if it is unused
// eat -

// eat ]

// do nothing

// eat end of range

func (v *globValidator) validate(pat string) { _ = "STUB: not implemented"; return }

// Handle first character if necessary

func validateGlob(pat string, isRef bool) []InvalidGlobPattern {
	_ = "STUB: not implemented"
	return nil
}

// ValidateRefGlob checks a given input as glob pattern for Git ref names. It returns list of
// errors found by the validation. See the following URL for more details of the syntax:
// https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions#filter-pattern-cheat-sheet
func ValidateRefGlob(pat string) []InvalidGlobPattern { _ = "STUB: not implemented"; return nil }

// ValidatePathGlob checks a given input as glob pattern for file paths. It returns list of
// errors found by the validation. See the following URL for more details of the syntax:
// https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions#filter-pattern-cheat-sheet
func ValidatePathGlob(pat string) []InvalidGlobPattern { _ = "STUB: not implemented"; return nil }

// '.' is not handled by path filter (#521)
