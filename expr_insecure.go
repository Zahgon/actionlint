package actionlint

import (
	"strings"
)

func isSafeFuncCall(call *FuncCallNode) bool { _ = "STUB: not implemented"; return false }

// UntrustedInputMap is a recursive map to match context object property dereferences.
// Root of this map represents each context names and their ancestors represent recursive properties.
type UntrustedInputMap struct {
	Name     string
	Parent   *UntrustedInputMap
	Children map[string]*UntrustedInputMap
}

func (m *UntrustedInputMap) String() string { _ = "STUB: not implemented"; return "" }

// Find child object property in this map
func (m *UntrustedInputMap) findObjectProp(name string) (*UntrustedInputMap, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Find child array element in this map. This is special case with object filter where its receiver is an array
func (m *UntrustedInputMap) findArrayElem() (*UntrustedInputMap, bool) {
	_ = "STUB: not implemented"
	return nil, false

	// Build path like `github.event.commits.*.body` by following parents
}

func (m *UntrustedInputMap) buildPath(b *strings.Builder) { _ = "STUB: not implemented"; return }

// NewUntrustedInputMap creates new instance of UntrustedInputMap. It is used for node of search
// tree of untrusted input checker.
func NewUntrustedInputMap(name string, children ...*UntrustedInputMap) *UntrustedInputMap {
	_ = "STUB: not implemented"
	return nil
}

// Leaf of the tree is nil

// UntrustedInputSearchRoots is a list of untrusted inputs. It forms tree structure to detect
// untrusted inputs in nested object property access, array index access, and object filters
// efficiently. Each value of this map represents a root of the search so their names are context
// names.
type UntrustedInputSearchRoots map[string]*UntrustedInputMap

// AddRoot adds a new root to search for detecting untrusted input.
func (ms UntrustedInputSearchRoots) AddRoot(m *UntrustedInputMap) {
	_ = "STUB: not implemented"

	// TODO: Automatically generate BuiltinUntrustedInputs from https://github.com/github/codeql/blob/main/javascript/ql/src/experimental/Security/CWE-094/ExpressionInjection.ql
	return
}

// BuiltinUntrustedInputs is list of untrusted inputs. These inputs are detected as untrusted in
// `run:` scripts. See the URL for more details.
// - https://securitylab.github.com/research/github-actions-untrusted-input/
// - https://docs.github.com/en/actions/reference/security/secure-use#good-practices-for-mitigating-script-injection-attacks
// - https://github.com/github/codeql/blob/main/javascript/ql/src/experimental/Security/CWE-094/ExpressionInjection.ql
var BuiltinUntrustedInputs = UntrustedInputSearchRoots{
	"github": NewUntrustedInputMap("github",
		NewUntrustedInputMap("event",
			NewUntrustedInputMap("issue",
				NewUntrustedInputMap("title"),
				NewUntrustedInputMap("body"),
			),
			NewUntrustedInputMap("pull_request",
				NewUntrustedInputMap("title"),
				NewUntrustedInputMap("body"),
				NewUntrustedInputMap("head",
					NewUntrustedInputMap("ref"),
					NewUntrustedInputMap("label"),
					NewUntrustedInputMap("repo",
						NewUntrustedInputMap("default_branch"),
					),
				),
			),
			NewUntrustedInputMap("comment",
				NewUntrustedInputMap("body"),
			),
			NewUntrustedInputMap("review",
				NewUntrustedInputMap("body"),
			),
			NewUntrustedInputMap("review_comment",
				NewUntrustedInputMap("body"),
			),
			NewUntrustedInputMap("pages",
				NewUntrustedInputMap("*",
					NewUntrustedInputMap("page_name"),
				),
			),
			NewUntrustedInputMap("commits",
				NewUntrustedInputMap("*",
					NewUntrustedInputMap("message"),
					NewUntrustedInputMap("author",
						NewUntrustedInputMap("email"),
						NewUntrustedInputMap("name"),
					),
				),
			),
			NewUntrustedInputMap("head_commit",
				NewUntrustedInputMap("message"),
				NewUntrustedInputMap("author",
					NewUntrustedInputMap("email"),
					NewUntrustedInputMap("name"),
				),
			),
			NewUntrustedInputMap("discussion",
				NewUntrustedInputMap("title"),
				NewUntrustedInputMap("body"),
			),
		),
		NewUntrustedInputMap("head_ref"),
	),
}

// UntrustedInputChecker is a checker to detect untrusted inputs in an expression syntax tree.
// This checker checks object property accesses, array index accesses, and object filters. And
// detects paths to untrusted inputs. Found errors are stored in this instance and can be get via
// Errs method.
//
// Note: To avoid breaking the state of checking property accesses on nested property accesses like
// foo[aaa.bbb].bar, IndexAccessNode.Index must be visited before IndexAccessNode.Operand.
type UntrustedInputChecker struct {
	roots           UntrustedInputSearchRoots
	filteringObject bool
	cur             []*UntrustedInputMap
	start           ExprNode
	errs            []*ExprError
	safeCalls       int
}

// NewUntrustedInputChecker creates a new UntrustedInputChecker instance. The roots argument is a
// search tree which defines untrusted input paths as trees.
func NewUntrustedInputChecker(roots UntrustedInputSearchRoots) *UntrustedInputChecker {
	_ = "STUB: not implemented"
	return nil
}

// Reset the state for next search
func (u *UntrustedInputChecker) reset() { _ = "STUB: not implemented"; return }

func (u *UntrustedInputChecker) compact() { _ = "STUB: not implemented"; return }

func (u *UntrustedInputChecker) onVar(v *VariableNode) { _ = "STUB: not implemented"; return }

// Find root context (currently only "github" exists)

func (u *UntrustedInputChecker) onPropAccess(name string) { _ = "STUB: not implemented"; return }

// depth + 1

func (u *UntrustedInputChecker) onIndexAccess() { _ = "STUB: not implemented"; return }

// For example, match `github.event.*.body[0]` as `github.event.commits[0].body`

func (u *UntrustedInputChecker) onObjectFilter() { _ = "STUB: not implemented"; return }

// Object filter for arrays

// Object filter for objects

func (u *UntrustedInputChecker) end() { _ = "STUB: not implemented"; return }

// When `Children` is nil, the node is a leaf

// When multiple untrusted inputs are detected, it means the expression extracts multiple properties with object
// filter syntax. Show all properties in error message.

func (u *UntrustedInputChecker) OnVisitNodeEnter(n ExprNode) { _ = "STUB: not implemented"; return }

// OnVisitNodeLeave is a callback which should be called on visiting node after visiting its children.
func (u *UntrustedInputChecker) OnVisitNodeLeave(n ExprNode) {
	_ = "STUB: not implemented"
	// Skip unsafe checks if we are inside of safe function call expression
	return
}

// Special case like github['event']['issue']['title']

// OnVisitEnd is a callback which should be called after visiting whole syntax tree. This callback
// is necessary to handle the case where an untrusted input access is at root of expression.
func (u *UntrustedInputChecker) OnVisitEnd() {
	_ = "STUB: not implemented"

	// Errs returns errors detected by this checker. This method should be called after visiting all
	// nodes in a syntax tree.
	return
}

func (u *UntrustedInputChecker) Errs() []*ExprError {
	_ = "STUB: not implemented"

	// Init initializes a state of checker.
	return nil
}

func (u *UntrustedInputChecker) Init() { _ = "STUB: not implemented"; return }
