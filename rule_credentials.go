package actionlint

// RuleCredentials is a rule to check credentials in workflows
type RuleCredentials struct {
	RuleBase
}

// NewRuleCredentials creates new RuleCredentials instance
func NewRuleCredentials() *RuleCredentials { _ = "STUB: not implemented"; return nil }

// VisitJobPre is callback when visiting Job node before visiting its children.
func (rule *RuleCredentials) VisitJobPre(n *Job) error { _ = "STUB: not implemented"; return nil }

func (rule *RuleCredentials) checkContainer(where string, n *Container) {
	_ = "STUB: not implemented"
	return
}
