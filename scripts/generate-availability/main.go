package main

import (
	"io"
	"log"
	"os"
	"regexp"

	"github.com/yuin/goldmark/ast"
	extast "github.com/yuin/goldmark/extension/ast"
)

const theURL = "https://raw.githubusercontent.com/github/docs/refs/heads/main/content/actions/reference/workflows-and-actions/contexts.md"

var dbg = log.New(io.Discard, "", log.LstdFlags)
var reReplaceholder = regexp.MustCompile("{%[^%]+%}")

// `Node.Text` method was deprecated. This is alternative to it.
// https://github.com/yuin/goldmark/issues/471
func textOf(n ast.Node, src []byte) string { _ = "STUB: not implemented"; return "" }

type switchCase struct {
	ctx  []string
	sp   []string
	cond []string
}
type switchCases map[string]*switchCase

func (sc switchCases) Add(key string, ctx []string, sp []string) { _ = "STUB: not implemented"; return }

func (sc switchCases) ForEach(pred func(c *switchCase)) { _ = "STUB: not implemented"; return }

func (sc switchCases) Contains(key string) bool { _ = "STUB: not implemented"; return false }

func parseContextAvailabilityTable(src []byte) (*extast.Table, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func cells(n *extast.TableRow, src []byte) []string { _ = "STUB: not implemented"; return nil }

func split(text string) []string { _ = "STUB: not implemented"; return nil }

func stripAndUnescape(s string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func generate(src []byte, out io.Writer) error { _ = "STUB: not implemented"; return nil }

// 'None' means no special function is available. It was added by this commit:
// https://github.com/github/docs/commit/ed18f98d128a2720d9a285b1ed48b161e4b9b7ef

// XXX: `jobs.<job_id>.snapshot.if` is missing in contexts available table.
// https://github.com/github/docs/issues/41255

// https://github.com/actions/runner/blob/c96dcd472907e514274cdb4af281116adf7aad18/src/Sdk/WorkflowParser/Conversion/WorkflowTemplateConverter.cs#L2216

// https://github.com/actions/runner/blob/c96dcd472907e514274cdb4af281116adf7aad18/src/Sdk/WorkflowParser/Conversion/WorkflowTemplateConverter.cs#L2240

// This variable is for unit tests

func source(args []string, url string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func run(args []string, stdout, stderr, dbgout io.Writer, srcURL string) int {
	_ = "STUB: not implemented"
	return 0
}

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr, os.Stderr, theURL))
}
