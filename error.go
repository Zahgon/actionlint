package actionlint

import (
	"io"
	"sync"
	"text/template"

	"github.com/fatih/color"
)

var (
	bold   = color.New(color.Bold)
	green  = color.New(color.FgGreen)
	yellow = color.New(color.FgYellow)
	gray   = color.New(color.FgHiBlack)
)

// Error represents an error detected by actionlint rules
type Error struct {
	// Message is an error message.
	Message string
	// Filepath is a file path where the error occurred.
	Filepath string
	// Line is a line number where the error occurred. This value is 1-based.
	Line int
	// Column is a column number where the error occurred. This value is 1-based.
	Column int
	// Kind is a string to represent kind of the error. Usually rule name which found the error.
	Kind string
}

// Error returns summary of the error as string.
func (e *Error) Error() string { _ = "STUB: not implemented"; return "" }

func (e *Error) String() string { _ = "STUB: not implemented"; return "" }

func errorAt(pos *Pos, kind string, msg string) *Error { _ = "STUB: not implemented"; return nil }

func errorfAt(pos *Pos, kind string, format string, args ...interface{}) *Error {
	_ = "STUB: not implemented"
	return nil
}

// GetTemplateFields fields for formatting this error with Go template.
func (e *Error) GetTemplateFields(source []byte) *ErrorTemplateFields {
	_ = "STUB: not implemented"
	return nil
}

// Byte length can be used here because this line only contains ASCII

// PrettyPrint prints the error with user-friendly way. It prints file name, source position, error
// message with colorful output and source snippet with indicator. When nil is set to source, no
// source snippet is not printed. To disable colorful output, set true to fatih/color.NoColor.
func (e *Error) PrettyPrint(w io.Writer, source []byte) { _ = "STUB: not implemented"; return }

func (e *Error) getLine(source []byte) (string, bool) { _ = "STUB: not implemented"; return "", false }

func (e *Error) getIndicator(line string) string { _ = "STUB: not implemented"; return "" }

// Column is 1-based

// Count width of non-space characters after '^' for underline

// Decrement for place for '^'

// Count width of spaces before '^'

func compareErrors(lhs, rhs *Error) int { _ = "STUB: not implemented"; return 0 }

func equalsErrors(lhs, rhs *Error) bool { _ = "STUB: not implemented"; return false }

// ErrorTemplateFields holds all fields to format one error message.
type ErrorTemplateFields struct {
	// Message is error message body.
	Message string `json:"message"`
	// Filepath is a canonical relative file path. This is empty when input was read from stdin.
	// When encoding into JSON, this field may be omitted when the file path is empty.
	Filepath string `json:"filepath,omitempty"`
	// Line is a line number of error position.
	Line int `json:"line"`
	// Column is a column number of error position.
	Column int `json:"column"`
	// Kind is a rule name the error belongs to.
	Kind string `json:"kind"`
	// Snippet is a code snippet and indicator to indicate where the error occurred.
	// When encoding into JSON, this field may be omitted when the snippet is empty.
	Snippet string `json:"snippet,omitempty"`
	// EndColumn is a column number where the error indicator (^~~~~~~) ends. When no indicator
	// can be shown, EndColumn is equal to Column.
	EndColumn int `json:"end_column"`
}

func unescapeBackslash(s string) string {
	_ = "STUB: not implemented"
	// https://golang.org/ref/spec#Rune_literals
	return ""
}

func toPascalCase(s string) string { _ = "STUB: not implemented"; return "" }

type ruleTemplateFields struct {
	Name        string
	Description string
}

func compareRuleTemplateByName(lhs, rhs *ruleTemplateFields) int {
	_ = "STUB: not implemented"
	return 0
}

// ErrorFormatter is a formatter to format a slice of ErrorTemplateFields. It is used for
// formatting error messages with -format option.
type ErrorFormatter struct {
	temp    *template.Template
	rules   map[string]*ruleTemplateFields
	rulesMu sync.Mutex
}

// NewErrorFormatter creates new ErrorFormatter instance. Given format must contain at least one
// {{ }} placeholder. Escaped characters like \n in the format string are unescaped.
func NewErrorFormatter(format string) (*ErrorFormatter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Print formats the slice of template fields and prints it with given writer.
func (f *ErrorFormatter) Print(out io.Writer, t []*ErrorTemplateFields) error {
	_ = "STUB: not implemented"
	return nil
}

// PrintErrors prints the errors after formatting them with template.
func (f *ErrorFormatter) PrintErrors(out io.Writer, errs []*Error, src []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// RegisterRule registers the rule. Registered rules are used to get description and index of error
// kinds when you use `kindDescription` or `kindIndex` functions in an error format template. This
// method can be called multiple times safely in parallel.
func (f *ErrorFormatter) RegisterRule(r Rule) {
	_ = "STUB: not implemented"
	// Synchronize access to f.rules (#370)
	return
}
