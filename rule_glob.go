package actionlint

// RuleGlob is a rule to check glob syntax.
// https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions#filter-pattern-cheat-sheet
type RuleGlob struct {
	RuleBase
}

// NewRuleGlob creates new RuleGlob instance.
func NewRuleGlob() *RuleGlob { _ = "STUB: not implemented"; return nil }

// VisitWorkflowPre is callback when visiting Workflow node before visiting its children.
func (rule *RuleGlob) VisitWorkflowPre(n *Workflow) error { _ = "STUB: not implemented"; return nil }

func (rule *RuleGlob) VisitJobPre(n *Job) error { _ = "STUB: not implemented"; return nil }

func (rule *RuleGlob) checkRefGlob(s *String) {
	_ = "STUB: not implemented"
	// Empty value is already checked by parser. Avoid duplicate errors
	return
}

func (rule *RuleGlob) checkGitRefGlobs(filter *WebhookEventFilter) {
	_ = "STUB: not implemented"
	return
}

func (rule *RuleGlob) checkFilePathGlobs(filter *WebhookEventFilter) {
	_ = "STUB: not implemented"
	return
}

// Empty value is already checked by parser. Avoid duplicate errors

func (rule *RuleGlob) globErrors(errs []InvalidGlobPattern, pos *Pos, quoted bool) {
	_ = "STUB: not implemented"
	return
}
