package actionlint

import (
	"sync"
)

type shellcheckError struct {
	Line    int    `json:"line"`
	Column  int    `json:"column"`
	Level   string `json:"level"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// RuleShellcheck is a rule to check shell scripts at 'run:' using shellcheck.
// https://github.com/koalaman/shellcheck
type RuleShellcheck struct {
	RuleBase
	cmd           *externalCommand
	workflowShell string
	jobShell      string
	runnerShell   string
	mu            sync.Mutex
}

func newRuleShellcheck(cmd *externalCommand) *RuleShellcheck { _ = "STUB: not implemented"; return nil }

// NewRuleShellcheck creates new RuleShellcheck instance. The executable argument can be command
// name or relative/absolute file path. When the given executable is not found in system, it returns
// an error as 2nd return value.
func NewRuleShellcheck(executable string, proc *concurrentProcess) (*RuleShellcheck, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// VisitStep is callback when visiting Step node.
func (rule *RuleShellcheck) VisitStep(n *Step) error { _ = "STUB: not implemented"; return nil }

// VisitJobPre is callback when visiting Job node before visiting its children.
func (rule *RuleShellcheck) VisitJobPre(n *Job) error { _ = "STUB: not implemented"; return nil }

// Default shell on Windows is PowerShell.
// https://docs.github.com/en/actions/learn-github-actions/workflow-syntax-for-github-actions#using-a-specific-shell

// VisitJobPost is callback when visiting Job node after visiting its children.
func (rule *RuleShellcheck) VisitJobPost(n *Job) error { _ = "STUB: not implemented"; return nil }

// VisitWorkflowPre is callback when visiting Workflow node before visiting its children.
func (rule *RuleShellcheck) VisitWorkflowPre(n *Workflow) error {
	_ = "STUB: not implemented"
	return nil
}

// VisitWorkflowPost is callback when visiting Workflow node after visiting its children.
func (rule *RuleShellcheck) VisitWorkflowPost(n *Workflow) error {
	_ = "STUB: not implemented"
	return nil
}

// Wait until all processes running for this rule

func (rule *RuleShellcheck) getShellName(exec *ExecRun) string {
	_ = "STUB: not implemented"
	return ""
}

// Note: Default shell on Windows is pwsh so this value is not always correct.
// Note: When bash is not found, GitHub-hosted runner fallbacks to sh.

// Replace ${{ ... }} with underscores like __________
// Note: replacing with spaces sometimes causes syntax error. For example,
//
//	if ${{ contains(xs, s) }}; then
//	  echo 'hello'
//	fi
func sanitizeExpressionsInScript(src string) string { _ = "STUB: not implemented"; return "" }

// 2 is offset for len("}}")

// Note: If ${{ ... }} includes newline, line and column reported by shellcheck will be
// shifted.

func (rule *RuleShellcheck) runShellcheck(src, shell string, pos *Pos) {
	_ = "STUB: not implemented"
	return
}

// Skip checking this shell script since shellcheck doesn't support it

// Reasons to exclude the rules:
//
// - SC1091: File not found. Scripts are for CI environment. Not suitable for checking this in current local
//           environment
// - SC2194: The word is constant. This sometimes happens at constants by replacing ${{ }} with underscores.
//           For example, `if ${{ matrix.foo }}; then ...` -> `if _________________; then ...`
// - SC2050: The expression is constant. This sometimes happens at `if` condition by replacing ${{ }} with
//           underscores (#45). For example, `if [ "${{ matrix.foo }}" = "x" ]` -> `if [ "_________________" = "x" ]`
// - SC2153: Same as SC2154.
// - SC2154: The var is referenced but not assigned. Script at `run:` can refer variables defined in `env:` section
//           so this rule can cause false positives (#53).
// - SC2157: Argument to -z is always false due to literal strings. When the argument of -z is replaced from ${{ }},
//           this can happen. For example, `if [ -z ${{ env.FOO }} ]` -> `if [ -z ______________ ]` (#113).
// - SC2043: Loop can be detected as only running once when the target of iteration is a placeholder. (#355)
//           e.g. `for foo in ${{ inputs.foo }}; do`

// Use same options to run shell process described at document
// https://docs.github.com/en/actions/learn-github-actions/workflow-syntax-for-github-actions#using-a-specific-shell

// Synchronize rule.Errorf calls

// It's better to show source location in the script as position of error, but it's not
// possible easily. YAML has multiple block styles with '|', '>', '|+', '>+', '|-', '>-'. Some
// of them remove indentation and/or blank lines. So restoring source position in block string
// is not possible. Sourcemap is necessary to do it.
// Instead, actionlint shows position of 'run:' as position of error. And separately show
// location in script which is reported by shellcheck in error message.

// Consider the first line is setup for running shell which was implicitly added for better check

// Trim period aligning style of error message
