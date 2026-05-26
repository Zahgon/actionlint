package actionlint

import (
	"sync"
)

type shellIsPythonKind int

const (
	shellIsPythonKindUnspecified shellIsPythonKind = iota
	shellIsPythonKindPython
	shellIsPythonKindNotPython
)

func getShellIsPythonKind(shell *String) shellIsPythonKind {
	_ = "STUB: not implemented"
	return *new(shellIsPythonKind)
}

// RulePyflakes is a rule to check Python scripts at 'run:' using pyflakes.
// https://github.com/PyCQA/pyflakes
type RulePyflakes struct {
	RuleBase
	cmd                   *externalCommand
	workflowShellIsPython shellIsPythonKind
	jobShellIsPython      shellIsPythonKind
	mu                    sync.Mutex
}

func newRulePyflakes(cmd *externalCommand) *RulePyflakes { _ = "STUB: not implemented"; return nil }

// NewRulePyflakes creates new RulePyflakes instance. Parameter executable can be command name
// or relative/absolute file path. When the given executable is not found in system, it returns
// an error.
func NewRulePyflakes(executable string, proc *concurrentProcess) (*RulePyflakes, error) {
	_ = "STUB: not implemented"
	// Combine output because pyflakes outputs lint errors to stdout and outputs syntax errors to stderr. (#411)
	return nil, nil
}

// VisitJobPre is callback when visiting Job node before visiting its children.
func (rule *RulePyflakes) VisitJobPre(n *Job) error { _ = "STUB: not implemented"; return nil }

// VisitJobPost is callback when visiting Job node after visiting its children.
func (rule *RulePyflakes) VisitJobPost(n *Job) error { _ = "STUB: not implemented"; return nil }

// reset

// VisitWorkflowPre is callback when visiting Workflow node before visiting its children.
func (rule *RulePyflakes) VisitWorkflowPre(n *Workflow) error {
	_ = "STUB: not implemented"
	return nil
}

// VisitWorkflowPost is callback when visiting Workflow node after visiting its children.
func (rule *RulePyflakes) VisitWorkflowPost(n *Workflow) error {
	_ = "STUB: not implemented"
	return nil
}

// reset
// Wait until all processes running for this rule

// VisitStep is callback when visiting Step node.
func (rule *RulePyflakes) VisitStep(n *Step) error { _ = "STUB: not implemented"; return nil }

func (rule *RulePyflakes) isPythonShell(r *ExecRun) bool { _ = "STUB: not implemented"; return false }

func (rule *RulePyflakes) runPyflakes(src string, pos *Pos) { _ = "STUB: not implemented"; return }

// Defined at rule_shellcheck.go

func (rule *RulePyflakes) parseNextError(stdout []byte, pos *Pos) ([]byte, error) {
	_ = "STUB: not implemented"

	// Search the start of error message.
	return nil, nil
}

// Syntax errors from pyflake consist of multiple lines. Skip subsequent lines. (#411)
// ```
// <stdin>:1:7: unexpected EOF while parsing
// print(
//       ^
// ```

// This method needs to be thread-safe since concurrentProcess.run calls its callback in a different goroutine.
