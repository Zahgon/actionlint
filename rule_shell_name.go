package actionlint

type platformKind int

const (
	platformKindAny platformKind = iota
	platformKindMacOrLinux
	platformKindWindows
)

// RuleShellName is a rule to check 'shell' field. For more details, see
// https://docs.github.com/en/actions/learn-github-actions/workflow-syntax-for-github-actions#using-a-specific-shell
type RuleShellName struct {
	RuleBase
	platform platformKind
}

// NewRuleShellName creates new RuleShellName instance.
func NewRuleShellName() *RuleShellName { _ = "STUB: not implemented"; return nil }

// VisitStep is callback when visiting Step node.
func (rule *RuleShellName) VisitStep(n *Step) error { _ = "STUB: not implemented"; return nil }

// VisitJobPre is callback when visiting Job node before visiting its children.
func (rule *RuleShellName) VisitJobPre(n *Job) error { _ = "STUB: not implemented"; return nil }

// VisitJobPost is callback when visiting Job node after visiting its children.
func (rule *RuleShellName) VisitJobPost(n *Job) error { _ = "STUB: not implemented"; return nil }

// Clear

// VisitWorkflowPre is callback when visiting Workflow node before visiting its children.
func (rule *RuleShellName) VisitWorkflowPre(n *Workflow) error {
	_ = "STUB: not implemented"
	return nil
}

func (rule *RuleShellName) checkShellName(node *String) { _ = "STUB: not implemented"; return }

// Ignore custom shell
// https://docs.github.com/en/actions/learn-github-actions/workflow-syntax-for-github-actions#custom-shell

// Ignore dynamic shell name

// ok

// only when the shell is unavailable on Windows

// only when the shell is unavailable on macOS or Linux

func getAvailableShellNames(kind platformKind) []string {
	_ = "STUB: not implemented"
	// https://docs.github.com/en/actions/learn-github-actions/workflow-syntax-for-github-actions#using-a-specific-shell
	return nil
}

func (rule *RuleShellName) getPlatformFromRunner(runner *Runner) platformKind {
	_ = "STUB: not implemented"
	return *new(platformKind)
}

// Note: Labels for self-hosted runners:
// https://docs.github.com/en/actions/hosting-your-own-runners/using-labels-with-self-hosted-runners

// Conflicts are reported by runner-label rule so simply ignore here
