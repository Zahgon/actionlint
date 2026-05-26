package actionlint

import (
	"regexp"
)

var jobIDPattern = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_-]*$`)

// RuleID is a rule to check step IDs in workflow.
type RuleID struct {
	RuleBase
	seen map[string]*Pos
}

// NewRuleID creates a new RuleID instance.
func NewRuleID() *RuleID { _ = "STUB: not implemented"; return nil }

// VisitJobPre is callback when visiting Job node before visiting its children.
func (rule *RuleID) VisitJobPre(n *Job) error { _ = "STUB: not implemented"; return nil }

// VisitJobPost is callback when visiting Job node after visiting its children.
func (rule *RuleID) VisitJobPost(n *Job) error { _ = "STUB: not implemented"; return nil }

// VisitStep is callback when visiting Step node.
func (rule *RuleID) VisitStep(n *Step) error { _ = "STUB: not implemented"; return nil }

func (rule *RuleID) validateConvention(id *String, what string) { _ = "STUB: not implemented"; return }
