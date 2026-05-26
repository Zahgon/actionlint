package actionlint

// RuleEnvVar is a rule checker to check environment variables setup.
type RuleEnvVar struct {
	RuleBase
}

// NewRuleEnvVar creates new RuleEnvVar instance.
func NewRuleEnvVar() *RuleEnvVar { _ = "STUB: not implemented"; return nil }

// VisitStep is callback when visiting Step node.
func (rule *RuleEnvVar) VisitStep(n *Step) error { _ = "STUB: not implemented"; return nil }

// VisitJobPre is callback when visiting Job node before visiting its children.
func (rule *RuleEnvVar) VisitJobPre(n *Job) error { _ = "STUB: not implemented"; return nil }

// VisitWorkflowPre is callback when visiting Workflow node before visiting its children.
func (rule *RuleEnvVar) VisitWorkflowPre(n *Workflow) error { _ = "STUB: not implemented"; return nil }

func (rule *RuleEnvVar) checkEnv(env *Env) { _ = "STUB: not implemented"; return }

// Key name can contain expressions (#312)
