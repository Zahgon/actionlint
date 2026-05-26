package actionlint

// RuleIfCond is a rule to check if: conditions.
type RuleIfCond struct {
	RuleBase
}

// NewRuleIfCond creates new RuleIfCond instance.
func NewRuleIfCond() *RuleIfCond { _ = "STUB: not implemented"; return nil }

// VisitStep is callback when visiting Step node.
func (rule *RuleIfCond) VisitStep(n *Step) error { _ = "STUB: not implemented"; return nil }

// VisitJobPre is callback when visiting Job node before visiting its children.
func (rule *RuleIfCond) VisitJobPre(n *Job) error { _ = "STUB: not implemented"; return nil }

func (rule *RuleIfCond) checkIfCond(n *String) { _ = "STUB: not implemented"; return }

func (rule *RuleIfCond) checkPlaceholder(n *String, start, end int) {
	_ = "STUB: not implemented"
	// Check number of ${{ }} for conditions like `${{ false }} || ${{ true }}` which are always evaluated to true
	return
}

func (rule *RuleIfCond) checkExpression(pos *Pos, input string) { _ = "STUB: not implemented"; return }
